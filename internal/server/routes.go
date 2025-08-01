package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/web/pages"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", pages.Home)

	e.GET("/docs", pages.Docs)
	e.GET("/palette", pages.Palette)
	e.GET("/icons", pages.Icons)
	e.GET("/blog", pages.Blog)
	e.GET("/faq", pages.FAQ)

	e.GET("/account", pages.Account)
	e.GET("/settings", pages.Settings)
}

func RegisterStaticRoutes(e *echo.Echo) {
	e.Static("/static", "views/static")
}
