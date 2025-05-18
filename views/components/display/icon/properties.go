package icon

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/display/tooltip"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
)

type Display int

const (
	ICON Display = iota
	ICON_TEXT
)

type Props struct {
	Icon icons.Icon

	Type    Display
	Content components.Customizable
	Tooltip components.Customizable

	NoAnimate bool
}

func (p *Props) RenderTooltip() templ.Component {
	if p.Tooltip != nil {
		return tooltip.Render(p.Tooltip)
	}
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
