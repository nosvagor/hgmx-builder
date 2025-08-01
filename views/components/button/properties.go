package button

import (
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/utilities/htmx"
)

type Variant int

const (
	Base Variant = iota

	Primary
	Secondary
	External

	Constructive
	Transformative
	Destructive
)

type Props struct {
	components.Styler

	Href    string
	HX      *htmx.Props
	Content components.Customizeable
	Variant Variant
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}
