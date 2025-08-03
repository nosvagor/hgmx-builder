package logo

import "github.com/nosvagor/hgmx-builder/views/components"

func Icon(classes ...string) components.Customizeable {
	p := &Props{Logo: icon()}
	p.Styler.Add(classes...)
	return p
}

func Full(classes ...string) components.Customizeable {
	p := &Props{Logo: full()}
	p.Styler.Add(classes...)
	return p
}

func Duo(classes ...string) components.Customizeable {
	p := &Props{Logo: duo()}
	p.Styler.Add(classes...)
	return p
}
