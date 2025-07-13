package navbar

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components/display/button"
	"github.com/nosvagor/hgmx-builder/views/components/display/content"
	"github.com/nosvagor/hgmx-builder/views/components/display/icon"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
	"github.com/nosvagor/hgmx-builder/views/htmx"
)

func New() *Props {
	return &Props{
		Links: []templ.Component{
			button.Get("/docs", icon.Text(icons.Scroll(), "Docs")),
			button.Get("/palette", icon.Text(icons.Palette(), "Palette")),
			button.Get("/icons", icon.Text(icons.Orbit(), "Icons")),
			button.Get("/blog", icon.Text(icons.Rss(), "Blog")),
		},
		Settings: button.Get("/settings",
			icon.Tip(icons.Settings(), content.Text("Settings Tooltip")),
			htmx.Target("#main"),
			htmx.Swap("outerHTML"),
		),
	}
}
