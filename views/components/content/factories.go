package content

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
)

func Text(text string) components.Customizeable {
	return &Props{Text: text}
}

func Component(c templ.Component) components.Customizeable {
	return &Props{Component: c}
}
