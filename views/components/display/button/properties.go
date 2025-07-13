package button

import (
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/htmx"
)

type Variant int

const (
	Base Variant = iota
	Primary

	Constructive
	Transformative
	Destructive
)

type Props struct {
	components.Styler

	HX      *htmx.Props
	Content components.Customizeable
	Variant Variant
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}
