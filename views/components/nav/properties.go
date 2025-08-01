package nav

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/internal/utils"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/button"
	"github.com/nosvagor/hgmx-builder/views/components/icon"
	"github.com/nosvagor/hgmx-builder/views/components/logo"
	"github.com/nosvagor/hgmx-builder/views/icons"
	"github.com/nosvagor/hgmx-builder/views/utilities/htmx"
)

type Props struct {
	components.Standard

	Logo      templ.Component
	Bookmarks []templ.Component
	Account   templ.Component
	Settings  templ.Component
}

type Link struct {
	Path  string
	Icon  *icons.Icon
	Label string
}

func (b *Link) label() string {
	if b.Label != "" {
		return b.Label
	}

	return utils.Title(b.Path)
}

func Init() templ.Component {
	opts := []htmx.Options{htmx.Target("#main"), htmx.Swap("outerHTML"), htmx.PushURL("true")}

	p := &Props{
		Logo:     button.GetCustom("/", logo.Icon("text-surface-50", "text-[2rem]"), button.Base, "pl-1 pr-1", opts...),
		Settings: button.Get("/settings", icon.Text(icons.Settings(), "Settings"), opts...),
		Account:  button.Get("/account", icon.Text(icons.User(), "Account"), opts...),
	}

	bookmarks := []Link{
		{Path: "docs", Icon: icons.Scroll()},
		{Path: "palette", Icon: icons.Palette()},
		{Path: "icons", Icon: icons.Orbit()},
		{Path: "blog", Icon: icons.Rss()},
		{Path: "faq", Icon: icons.Question(), Label: "FAQ"},
		{Path: "users", Icon: icons.Users(), Label: "Users"},
		{Path: "github", Icon: icons.BrandGithub(), Label: "GitHub"},
		// {Path: "user-check", Icon: icons.UserCheck(), Label: "User Check"},
		// {Path: "user-plus", Icon: icons.UserPlus(), Label: "User Plus"},
		// {Path: "user-minus", Icon: icons.UserMinus(), Label: "User Minus"},
		// {Path: "user-x", Icon: icons.UserX(), Label: "User X"},
		// {Path: "user-circle", Icon: icons.UserCircle(), Label: "User Circle"},
		// {Path: "user-contact", Icon: icons.UserContact(), Label: "User Contact"},
		// {Path: "user-pen", Icon: icons.UserPen(), Label: "User Pen"},
		// {Path: "user-search", Icon: icons.UserSearch(), Label: "User Search"},
		// {Path: "user-star", Icon: icons.UserStar(), Label: "User Star"},
	}

	for _, link := range bookmarks {
		p.Bookmarks = append(p.Bookmarks,
			button.GetCustom(
				"/"+link.Path,
				icon.Text(link.Icon, link.label()),
				button.Base,
				"justify-start",
			),
		)
	}

	return p.Render()
}
