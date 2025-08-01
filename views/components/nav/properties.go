package nav

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/display/button"
	"github.com/nosvagor/hgmx-builder/views/components/display/content"
	"github.com/nosvagor/hgmx-builder/views/components/display/icon"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
	"github.com/nosvagor/hgmx-builder/views/htmx"
)

type Props struct {
	components.Standard

	Logo      templ.Component
	Bookmarks []templ.Component

	Account  templ.Component
	Settings templ.Component
}

func Init() templ.Component {
	opts := []htmx.Options{htmx.Target("#main"), htmx.Swap("outerHTML"), htmx.PushURL("true")}

	p := &Props{
		Logo:     button.Get("/", content.Text("hgmx"), opts...),
		Settings: button.Get("/settings", icon.Text(icons.Settings(), "Settings"), opts...),
		Account:  button.Get("/account", icon.Text(icons.User(), "Account"), opts...),
	}

	links := []struct {
		path string
		icon *icons.Icon
	}{
		{path: "docs", icon: icons.Scroll()},
		{path: "palette", icon: icons.Palette()},
		{path: "icons", icon: icons.Orbit()},
		{path: "showcase", icon: icons.Rss()},
		{path: "FAQ", icon: icons.Question()},
		{path: "users", icon: icons.Users()},
		{path: "user", icon: icons.User()},
		{path: "user-check", icon: icons.UserCheck()},
		{path: "user-circle", icon: icons.UserCircle()},
		{path: "user-cog", icon: icons.UserCog()},
		{path: "user-contact", icon: icons.UserContact()},
		{path: "user-minus", icon: icons.UserMinus()},
		{path: "user-pen", icon: icons.UserPen()},
		{path: "user-search", icon: icons.UserSearch()},
		{path: "user-star", icon: icons.UserStar()},
		{path: "user-plus", icon: icons.UserPlus()},
		{path: "user-x", icon: icons.UserX()},
	}

	for _, link := range links {
		p.Bookmarks = append(p.Bookmarks,
			button.GetCustom(
				"/"+strings.ToLower(link.path),
				icon.Text(link.icon, strings.ToUpper(link.path[:1])+link.path[1:]),
				button.Base,
				"justify-start",
			),
		)
	}

	return p.Render()
}
