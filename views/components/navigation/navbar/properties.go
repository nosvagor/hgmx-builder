package navbar

import (
	_ "embed"

	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
)

type Props struct {
	components.Styler

	Links    []templ.Component
	Account  templ.Component
	Settings templ.Component
	CTA      templ.Component
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}
