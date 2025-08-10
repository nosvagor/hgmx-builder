package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/database"
	"github.com/nosvagor/hgmx-builder/internal/web/pages"
)

// =============================================================================
// === Web Routes ==============================================================
// =============================================================================

func RegisterWebRoutes(e *echo.Echo) {
	e.GET("/", pages.Home)

	e.GET("/docs", pages.Docs)
	e.GET("/palette", pages.Palette)
	e.GET("/icons", pages.Icons)
	e.GET("/blog", pages.Blog)
	e.GET("/faq", pages.FAQ)

	e.GET("/account", pages.Account)
	e.GET("/settings", pages.Settings)
}

// =============================================================================
// === Health Routes ===========================================================
// =============================================================================

func RegisterHealthRoutes(e *echo.Echo) {
	e.GET("/health/db", func(c echo.Context) error {
		hs, err := database.Health(c.Request().Context())
		if err != nil {
			return c.JSON(503, hs)
		}
		return c.JSON(200, hs)
	})
}

// =============================================================================
// === Static Routes ===========================================================
// =============================================================================

func RegisterStaticRoutes(e *echo.Echo) {
	e.Static("/static", "views/static")
}
