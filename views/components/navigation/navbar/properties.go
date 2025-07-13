package navbar

import (
	_ "embed"

	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
)

type Props struct {
	Links             []PageLink
	Account, Settings PageLink
	CTA               templ.Component
	Logo              components.Standard
}

type PageLink struct {
	Icon  icons.Icon
	Label string
	Path  string
}
