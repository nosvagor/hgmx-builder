package home

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/web"
	"github.com/nosvagor/hgmx-builder/views/pages/home"
)

func Main(c echo.Context) error {
	return web.Page(c, home.Main(), "Home")
}
