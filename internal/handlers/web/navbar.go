package web

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/internal/services/stats"
	"github.com/nosvagor/hgmx-builder/views/components/app/logo"
	"github.com/nosvagor/hgmx-builder/views/components/app/source"
	"github.com/nosvagor/hgmx-builder/views/components/button"
	"github.com/nosvagor/hgmx-builder/views/components/icon"
	"github.com/nosvagor/hgmx-builder/views/components/icon/icons"
	"github.com/nosvagor/hgmx-builder/views/components/nav"
	"github.com/nosvagor/hgmx-builder/views/utils/htmx"
)

func fmtStars(stars int) string {
	if stars < 1000 {
		return fmt.Sprintf("%d", stars)
	}
	k := float64(stars) / 1000.0
	var s string
	if k < 100 {
		s = fmt.Sprintf("%.1f", k)
	} else {
		s = fmt.Sprintf("%.0f", k)
	}
	s = strings.TrimSuffix(s, ".0")
	return s + "k"
}

func navbar() templ.Component {
	opts := []htmx.Options{htmx.Target("#main"), htmx.Swap("outerHTML"), htmx.PushURL("true")}

	github := stats.Github{}
	gitStats, err := github.GetRepoStats(context.Background(), "nosvagor", "hgmx-builder", 1*time.Hour)
	if err != nil {
		fmt.Println(err)
	}

	p := &nav.Props{
		Logo:     button.GetCustom("/", logo.Full("text-surface-50", "text-lg"), button.Primary, "p-2", opts...),
		Settings: button.Get("/settings", icon.Text(icons.Settings(), "Settings"), opts...),
		Account:  button.Get("/account", icon.Text(icons.User(), "Account"), opts...),
		Source:   source.Source("github.com/nosvagor/hgmx-builder", icons.Github(), fmtStars(gitStats.Stars)),
		BookmarksMobile: [3]nav.Link{
			{Path: "docs", Icon: icons.Scroll()},
			{Path: "palette", Icon: icons.Palette()},
			{Path: "blog", Icon: icons.Rss()},
		},
		Bookmarks: []nav.Link{
			{Path: "docs", Icon: icons.Scroll()},
			{Path: "palette", Icon: icons.Palette()},
			{Path: "icons", Icon: icons.Orbit()},
			{Path: "blog", Icon: icons.Rss()},
			{Path: "faq", Icon: icons.Question(), Label: "FAQ"},
			{Path: "users", Icon: icons.Users()},
		},
	}

	return p.Render()
}
