package content

import "github.com/nosvagor/hgmx-builder/views/components"

func Text(text string) *Props {
	return &Props{Text: text}
}

func Component(c components.Customizable) *Props {
	return &Props{Component: c}
}
