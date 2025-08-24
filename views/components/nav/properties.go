package nav

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/internal/utils"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/icon/icons"
)

type Props struct {
	components.Standard

	Bookmarks []templ.Component
	Account   templ.Component
	Settings  templ.Component

	Logo   templ.Component
	Source templ.Component
}

type Link struct {
	Path  string
	Icon  *icons.Icon
	Label string
}

func (b *Link) Text() string {
	if b.Label != "" {
		return b.Label
	}

	return utils.Title(b.Path)
}
