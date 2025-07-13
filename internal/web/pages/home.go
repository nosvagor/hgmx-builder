package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/web"
	"github.com/nosvagor/hgmx-builder/views/pages/blog"
	"github.com/nosvagor/hgmx-builder/views/pages/docs"
	"github.com/nosvagor/hgmx-builder/views/pages/home"
	"github.com/nosvagor/hgmx-builder/views/pages/icons"
	"github.com/nosvagor/hgmx-builder/views/pages/palette"
	"github.com/nosvagor/hgmx-builder/views/pages/settings"
)

func Home(c echo.Context) error {
	return web.Page(c, home.Main(), "Home")
}

func Docs(c echo.Context) error {
	return web.Page(c, docs.Main(), "Docs")
}

func Palette(c echo.Context) error {
	return web.Page(c, palette.Main(), "Palette")
}

func Icons(c echo.Context) error {
	return web.Page(c, icons.Main(), "Icons")
}

func Blog(c echo.Context) error {
	return web.Page(c, blog.Main(), "Blog")
}

func Settings(c echo.Context) error {
	return web.Page(c, settings.Main(), "Settings")
}
