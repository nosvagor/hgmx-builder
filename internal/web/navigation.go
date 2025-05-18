package web

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views/components/app/logo"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
	"github.com/nosvagor/hgmx-builder/views/components/navigation/navbar"
)

func NewNavbar(c echo.Context) *navbar.Props {
	// TODO: do stuff with context and application config to fullfill navbar props
	return &navbar.Props{
		Links: []navbar.PageLink{
			{Label: "Docs", Path: "/docs", Icon: *icons.Scroll()},
			{Label: "Palette", Path: "/palette", Icon: *icons.Palette()},
			{Label: "Icons", Path: "/icons", Icon: *icons.Orbit()},
			{Label: "Blog", Path: "/blog", Icon: *icons.Rss()},
		},
		Logo:     logo.Text("HGMX"),
		Settings: navbar.PageLink{Label: "Settings Tooltip", Path: "/settings", Icon: *icons.Settings()},
	}
}
