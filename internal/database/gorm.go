package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/nosvagor/hgmx-builder/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func DSN(cfg utils.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)
}

func Open() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		cfg := utils.LoadConfig()
		gormCfg := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
		dsn := DSN(cfg)
		db, openErr := gorm.Open(postgres.Open(dsn), gormCfg)
		if openErr != nil {
			err = openErr
			return
		}

		sqlDB, sqlErr := db.DB()
		if sqlErr != nil {
			err = sqlErr
			return
		}

		sqlDB.SetMaxOpenConns(15)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		if pingErr := sqlDB.Ping(); pingErr != nil {
			err = pingErr
			return
		}

		dbInstance = db
	})
	return dbInstance, err
}

func MustOpen() *gorm.DB {
	db, err := Open()
	if err != nil {
		panic(err)
	}
	return db
}

func DB() *gorm.DB {
	return dbInstance
}

func Close() error {
	if dbInstance == nil {
		return nil
	}
	sqlDB, err := dbInstance.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type HealthStatus struct {
	OK                bool
	Latency           time.Duration
	Version           string
	Database          string
	User              string
	ReadOnly          bool
	OpenConnections   int
	InUse             int
	Idle              int
	WaitCount         int64
	WaitDuration      time.Duration
	MaxIdleClosed     int64
	MaxLifetimeClosed int64
	CheckedAt         time.Time
	Error             string
}

func Health(ctx context.Context) (HealthStatus, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
	}

	db, err := Open()
	if err != nil {
		return HealthStatus{OK: false, Error: err.Error(), CheckedAt: time.Now()}, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return HealthStatus{OK: false, Error: err.Error(), CheckedAt: time.Now()}, err
	}

	start := time.Now()
	if err := sqlDB.PingContext(ctx); err != nil {
		hs := HealthStatus{OK: false, Error: err.Error(), CheckedAt: time.Now()}
		return hs, err
	}
	latency := time.Since(start)

	stats := sqlDB.Stats()
	hs := HealthStatus{
		OK:                true,
		Latency:           latency,
		OpenConnections:   stats.OpenConnections,
		InUse:             stats.InUse,
		Idle:              stats.Idle,
		WaitCount:         stats.WaitCount,
		WaitDuration:      stats.WaitDuration,
		MaxIdleClosed:     stats.MaxIdleClosed,
		MaxLifetimeClosed: stats.MaxLifetimeClosed,
		CheckedAt:         time.Now(),
	}

	ctxQ, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()

	var v string
	if err := db.WithContext(ctxQ).Raw("show server_version").Scan(&v).Error; err == nil {
		hs.Version = v
	}
	var dbname string
	if err := db.WithContext(ctxQ).Raw("select current_database()").Scan(&dbname).Error; err == nil {
		hs.Database = dbname
	}
	var user string
	if err := db.WithContext(ctxQ).Raw("select current_user").Scan(&user).Error; err == nil {
		hs.User = user
	}
	var ro string
	if err := db.WithContext(ctxQ).Raw("show transaction_read_only").Scan(&ro).Error; err == nil {
		hs.ReadOnly = ro == "on"
	}

	return hs, nil
}
