package icon

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/icons"
)

type Display int

const (
	ICON Display = iota
	ICON_TEXT
)

type Props struct {
	components.Styler

	Icon      icons.Icon
	Type      Display
	Content   components.Customizeable
	Tooltip   components.Customizeable
	NoAnimate bool
}

func (p *Props) Style(classes ...string) components.Customizeable {
	p.Styler.Add(classes...)
	return p
}

func (p *Props) RenderTooltip() templ.Component {
	// if p.Tooltip != nil {
	// 	return tooltip.Render(p.Tooltip) // TODO: implement tooltip
	// }
	return templ.NopComponent
}

func (p *Props) RenderContent() templ.Component {
	if p.Type == ICON_TEXT {
		return p.Content.Render()
	}
	return templ.NopComponent
}

func (p *Props) RenderIcon() templ.Component {
	if p.Icon.Name == "" {
		icon := icons.Question()
		return icon.Render()
	}
	return p.Icon.Render()
}
