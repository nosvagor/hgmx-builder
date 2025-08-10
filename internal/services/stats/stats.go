package stats

import (
	"context"
	"time"

	"github.com/nosvagor/hgmx-builder/internal/database"
	"gorm.io/gorm/clause"
)

type RepoStat struct {
	ID          uint   `gorm:"primaryKey"`
	Owner       string `gorm:"size:100;not null;uniqueIndex:idx_owner_repo"`
	Repo        string `gorm:"size:200;not null;uniqueIndex:idx_owner_repo"`
	Stars       int
	Forks       int
	OpenIssues  int
	Subscribers int
	Size        int
	FetchedAt   time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (RepoStat) TableName() string { return "repo_stats" }

func Migrate() error {
	db := database.DB()
	if db == nil {
		db = database.MustOpen()
	}
	return db.AutoMigrate(&RepoStat{})
}

func GetCachedRepoStats(ctx context.Context, owner, repo string) (RepoStat, error) {
	var rs RepoStat
	DB := database.DB()
	if DB == nil {
		DB = database.MustOpen()
	}
	if err := DB.WithContext(ctx).Where("owner = ? AND repo = ?", owner, repo).Take(&rs).Error; err != nil {
		return RepoStat{}, err
	}
	return rs, nil
}

func UpsertRepoStats(ctx context.Context, rs RepoStat) (RepoStat, error) {
	DB := database.DB()
	if DB == nil {
		DB = database.MustOpen()
	}
	err := DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "owner"}, {Name: "repo"}},
		DoUpdates: clause.AssignmentColumns([]string{"stars", "forks", "open_issues", "subscribers", "size", "fetched_at", "updated_at"}),
	}).Create(&rs).Error
	if err != nil {
		return RepoStat{}, err
	}
	return rs, nil
}

func Stale(rs RepoStat, maxAge time.Duration) bool {
	if rs.ID == 0 {
		return true
	}
	if rs.FetchedAt.IsZero() {
		return true
	}
	return time.Since(rs.FetchedAt) > maxAge
}
