package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/services/colors"
	"github.com/nosvagor/hgmx-builder/internal/web"
	"github.com/nosvagor/hgmx-builder/views/pages/account"
	"github.com/nosvagor/hgmx-builder/views/pages/blog"
	"github.com/nosvagor/hgmx-builder/views/pages/docs"
	"github.com/nosvagor/hgmx-builder/views/pages/faq"
	"github.com/nosvagor/hgmx-builder/views/pages/home"
	"github.com/nosvagor/hgmx-builder/views/pages/notfound"
	"github.com/nosvagor/hgmx-builder/views/pages/palette"
	"github.com/nosvagor/hgmx-builder/views/pages/settings"
)

type CacheDuration int

const (
	NoStore = -1
	None    = 0
	Preload = 2
	Minute  = 60
	Hour    = 3600
	Day     = 86400
	Week    = 604800
)

func Home(c echo.Context) error {
	return web.Page(c, home.Main(), "Home", Preload)
}

func Docs(c echo.Context) error {
	return web.Page(c, docs.Main(), "Docs", Preload)
}

func Palette(c echo.Context) error {
	hex := c.QueryParam("hex")
	if hex == "" {
		hex = "#222536"
	}
	view := colors.Generate(hex).ToView()
	return web.Page(c, palette.Main(view, hex), "Palette", Preload)
}

func Blog(c echo.Context) error {
	return web.Page(c, blog.Main(), "Blog", Preload)
}

func FAQ(c echo.Context) error {
	return web.Page(c, faq.Main(), "FAQ", Preload)
}

func Settings(c echo.Context) error {
	return web.Page(c, settings.Main(), "Settings", Preload)
}

func Account(c echo.Context) error {
	return web.Page(c, account.Main(), "Account", Preload)
}

func NotFound(c echo.Context) error {
	return web.Page(c, notfound.Main(), "Not Found", Preload)
}
