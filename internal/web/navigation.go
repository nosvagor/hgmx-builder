package web

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views/components/app/actions"
	"github.com/nosvagor/hgmx-builder/views/components/navigation/navbar"
	"github.com/nosvagor/hgmx-builder/views/shared/icons"
)

func NewNavbar(c echo.Context) *navbar.Props {
	// TODO: do stuff with context and application config to fullfill navbar props
	return &navbar.Props{
		Name: "HGMX",
		Links: &[]navbar.PageLink{
			{Label: "Docs", Path: "/docs", Icon: icons.Scroll()},
			{Label: "Palette", Path: "/palette", Icon: icons.Palette()},
			{Label: "Icons", Path: "/icons", Icon: icons.Orbit()},
			{Label: "Blog", Path: "/blog", Icon: icons.Rss()},
		},
		CTA:      &actions.Props{Label: "Download", Path: "/download"},
		Settings: &navbar.PageLink{Label: "Settings", Path: "/settings", Icon: icons.Settings()},
	}
}
