package web

import (
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/shared"
	"github.com/nosvagor/hgmx-builder/views/components/navigation"
)

func NewNavbar(c echo.Context) *navigation.NavbarProps {
	// TODO: do stuff with context and application config to fullfill navbar props
	return &navigation.NavbarProps{
		App: "HGMX",
		Links: []navigation.PageLink{
			{Label: "Docs", URL: "/docs", Icon: shared.Ref("scroll")},
			{Label: "Palette", URL: "/palette"},
			{Label: "Icons", URL: "/icons", Icon: shared.Ref("orbit")},
			{Label: "Blog", URL: "/blog", Icon: shared.Ref("rss")},
		},
		// CTA: &navigation.PageLink{
		// 	Label: "Download",
		// 	URL:   "/download",
		// },
		Account: []navigation.PageLink{
			{Label: "Account", URL: "/account"},
			{Label: "Settings", URL: "/settings"},
		},
	}
}
