package content

import "github.com/nosvagor/hgmx-builder/views/components"

func Text(text string) components.Customizeable {
	return &Props{Text: text}
}

func Component(c components.Customizeable) components.Customizeable {
	return &Props{Component: c}
}
