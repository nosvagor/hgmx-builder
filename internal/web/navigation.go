package web

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views/components/navigation"
	"github.com/nosvagor/hgmx-builder/views/shared/icons"
)

func NewNavbar(c echo.Context) *navigation.NavbarProps {
	// TODO: do stuff with context and application config to fullfill navbar props
	return &navigation.NavbarProps{
		App: "HGMX",
		Links: []navigation.PageLink{
			{Label: "Docs", URL: "/docs", Icon: icons.Scroll()},
			{Label: "Palette", URL: "/palette", Icon: icons.Palette()},
			{Label: "Icons", URL: "/icons", Icon: icons.Orbit()},
			{Label: "Blog", URL: "/blog", Icon: icons.Rss()},
		},
		// CTA: &navigation.PageLink{
		// 	Label: "Download",
		// 	URL:   "/download",
		// },
		Settings: &navigation.PageLink{Label: "Settings", URL: "/settings", Icon: icons.Settings()},
	}
}
