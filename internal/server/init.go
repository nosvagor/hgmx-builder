package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/services/stats"
	database "github.com/nosvagor/hgmx-builder/internal/store"
	"github.com/nosvagor/hgmx-builder/internal/utils"
)

func New() *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = CustomHTTPErrorHandler
	e.Use(ZerologRecoverer())
	e.Use(ZerologRequestLogger())

	RegisterStaticRoutes(e)
	RegisterWebRoutes(e)
	RegisterHealthRoutes(e)

	// Optional: Register API routes if you want to expose APIs
	// RegisterAPIRoutes(e)

	// Optional: Register advanced routes with middleware
	// RegisterAdvancedRoutes(e)

	return e
}

const (
	blue   = "\033[34m"
	yellow = "\033[33m"
	faint  = "\033[2;37m"
	red    = "\033[31m"
	purple = "\033[35m"
	green  = "\033[32m"
	reset  = "\033[0m"
)

func printBanner(cfg utils.Config) {
	now := time.Now()
	zone, _ := now.Zone()
	tzOffset := now.Format("-07:00")

	ascii := `
    ██╗  ██╗ ██████╗ ███╗   ███╗██╗  ██╗
    ██║  ██║██╔════╝ ████╗ ████║╚██╗██╔╝
    ███████║██║  ███╗██╔████╔██║ ╚███╔╝ 
    ██╔══██║██║   ██║██║╚██╔╝██║ ██╔██╗ 
    ██║  ██║╚██████╔╝██║ ╚═╝ ██║██╔╝ ██╗
    ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝`
	builder := ".builder"

	asciiWidth := 47
	left := fmt.Sprintf("    http://%s:%s", cfg.Host, cfg.Port)

	logLevel := strings.ToUpper(cfg.LogLevel)
	var logColor string
	switch logLevel {
	case "DEBUG":
		logColor = purple
	case "INFO":
		logColor = blue
	case "WARN":
		logColor = yellow
	case "ERROR":
		logColor = red
	default:
		logColor = faint
	}

	right := fmt.Sprintf("%s[%s]%s", logColor, logLevel, reset)
	padding := max(asciiWidth-len(left)-len(stripAnsi(right)), 1)

	left2 := fmt.Sprintf("    %s (%s)", zone, tzOffset)
	right2 := fmt.Sprintf("%s (%s-%s)", runtime.Version(), runtime.GOARCH, runtime.GOOS)
	padding2 := max(asciiWidth-len(left2)-len(right2), 1)

	// Print colored banner
	fmt.Print(blue, ascii, reset)
	fmt.Print(yellow, builder, reset, "\n")
	fmt.Printf("%s%s%s%s%s\n", green, left, reset, strings.Repeat(" ", padding), right)
	fmt.Printf("%s%s%s%s%s\n", faint, left2, strings.Repeat(" ", padding2), right2, reset)
}

func printDBHealth(cfg utils.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	hs, err := database.Health(ctx)
	if err != nil || !hs.OK {
		msg := ""
		if err != nil {
			msg = err.Error()
		}
		if msg == "" {
			msg = "unhealthy"
		}
		fmt.Printf("%s    db: %s%s%s\n", red, msg, reset, faint)
		return
	}

	prod := utils.If(cfg.ENV == "DEV", "DEV ", "PROD")
	color := utils.If(cfg.ENV == "DEV", blue, red)
	open := fmt.Sprintf("%d", hs.OpenConnections)
	alive := fmt.Sprintf("%d", hs.InUse)
	idle := fmt.Sprintf("%d", hs.Idle)
	op := fmt.Sprintf("%s%s@%s%s", yellow, hs.User, hs.Database, blue)
	write := utils.If(hs.ReadOnly, fmt.Sprintf("%sread%s", red, color), fmt.Sprintf("%swrite%s", yellow, color))

	fmt.Printf(`%s
    |------------------------------------------------
    | %s%s%s_________ open:%s alive:%s idle:%s | %s
    |________  /_____________________________________
    |__  ___/  __/  __ \_  ___/  _ \  %s
    |_(__  )/ /_ / /_/ /  /   /  __/  %s 
    |/____/ \__/ \____//_/    \___/  󱐌 %s
    |------------------------------------------------ %s

`, color, yellow, prod, color, open, alive, idle, write, op, hs.Version, hs.Latency.String(), reset)
}

func stripAnsi(str string) string {
	res := ""
	skip := false
	for i := range len(str) {
		if str[i] == '\033' {
			skip = true
		}
		if !skip {
			res += string(str[i])
		}
		if skip && str[i] == 'm' {
			skip = false
		}
	}
	return res
}

func Start(e *echo.Echo) error {
	cfg := utils.LoadConfig()

	_ = database.MustOpen()
	_ = stats.Migrate()
	printBanner(cfg)
	printDBHealth(cfg)

	srv := &http.Server{Addr: cfg.Host + ":" + cfg.Port, Handler: e, ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 120 * time.Second}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = database.Close()
		_ = srv.Shutdown(ctx)
	}()

	return srv.ListenAndServe()
}
