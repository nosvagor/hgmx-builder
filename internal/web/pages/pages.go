package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/web"
	"github.com/nosvagor/hgmx-builder/views/pages/account"
	"github.com/nosvagor/hgmx-builder/views/pages/blog"
	"github.com/nosvagor/hgmx-builder/views/pages/docs"
	"github.com/nosvagor/hgmx-builder/views/pages/home"
	"github.com/nosvagor/hgmx-builder/views/pages/icons"
	"github.com/nosvagor/hgmx-builder/views/pages/notfound"
	"github.com/nosvagor/hgmx-builder/views/pages/palette"
	"github.com/nosvagor/hgmx-builder/views/pages/settings"
)

type CacheDuration int

const (
	NoStore = -1
	None    = 0
	Minute  = 60
	Hour    = 3600
	Day     = 86400
	Week    = 604800
)

func Home(c echo.Context) error {
	return web.Page(c, home.Main(), "Home", Minute)
}

func Docs(c echo.Context) error {
	return web.Page(c, docs.Main(), "Docs", 10*Minute)
}

func Palette(c echo.Context) error {
	return web.Page(c, palette.Main(), "Palette", 10*Minute)
}

func Icons(c echo.Context) error {
	return web.Page(c, icons.Main(), "Icons", 10*Minute)
}

func Blog(c echo.Context) error {
	return web.Page(c, blog.Main(), "Blog", Minute)
}

func Settings(c echo.Context) error {
	return web.Page(c, settings.Main(), "Settings", 10)
}

func Account(c echo.Context) error {
	return web.Page(c, account.Main(), "Account", 10)
}

func NotFound(c echo.Context) error {
	return web.Page(c, notfound.Main(), "Not Found", Week)
}
