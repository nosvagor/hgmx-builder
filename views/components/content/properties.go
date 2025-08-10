package content

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
)

type Props struct {
	components.Styler
	Text string

	Component templ.Component
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}
