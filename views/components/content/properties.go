package content

import (
	"github.com/nosvagor/hgmx-builder/views/components"
)

type Props struct {
	components.Styler
	Text string

	Component components.Customizeable
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}
