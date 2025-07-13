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

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
}

func ZerologRequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			duration := time.Since(start)

			var event *zerolog.Event
			if err != nil {
				event = log.Error()
				if he, ok := err.(*echo.HTTPError); ok {
					if msg, ok := he.Message.(string); ok {
						event.Str("e", msg)
					} else {
						event.Err(err)
					}
				} else {
					event.Err(err)
				}
			} else {
				event = log.Info()
			}

			event.
				Str("hx", req.Method).
				Int("c", res.Status).
				Str("uri", req.RequestURI).
				Str("ip", c.RealIP()).
				Str("ref", req.Referer()).
				Str("host", req.Host).
				Dur("ms", duration).
				Msg("")

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
