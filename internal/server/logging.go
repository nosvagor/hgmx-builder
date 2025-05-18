package server

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	cfg := utils.LoadConfig()

	var lvl zerolog.Level
	switch cfg.LogLevel {
	case "DEBUG":
		lvl = zerolog.DebugLevel
	case "INFO":
		lvl = zerolog.InfoLevel
	case "WARN":
		lvl = zerolog.WarnLevel
	case "ERROR":
		lvl = zerolog.ErrorLevel
	default:
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)

	if cfg.ENV == "PROD" {
		log.Logger = log.Output(zerolog.New(os.Stderr).With().Timestamp().Logger())
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
	}
}

func ZerologRequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			err := next(c)

			duration := time.Since(start)
			logger := log.Info()
			if err != nil {
				logger = log.Error().Err(err)
				c.Error(err)
			}

			logger.Str("hx", req.Method).
				Int("c", res.Status).
				Str("uri", req.RequestURI).
				Str("ip", c.RealIP()).
				Dur("ms", duration).
				Msg(" ")

			return nil
		}
	}
}

func ZerologRecoverer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := debug.Stack()

					log.Error().
						Err(err).
						Str("uri", c.Request().RequestURI).
						Msg("panic recovered")

					fmt.Fprintf(os.Stderr, "\n--- Stack Trace ---\n%s\n-------------------\n", string(stack))

					c.Error(echo.NewHTTPError(500, "Internal Server Error"))
				}
			}()
			return next(c)
		}
	}
}
