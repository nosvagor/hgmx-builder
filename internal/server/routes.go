package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/web/home"
)

func RegisterRoutes(e *echo.Echo) {
	// Home
	e.GET("/", home.Main)

	// // Settings
	// e.GET("/settings", settings.Main)

	// // Palette
	// e.GET("/palette", palette.Main)
}

func RegisterStaticRoutes(e *echo.Echo) {
	e.Static("/static", "views/static")
}
