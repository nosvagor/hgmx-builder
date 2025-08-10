package button

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/utilities/htmx"
)

func Get(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Get(path, opts...), Content: content, Variant: Primary}
	return p.Render()
}

func GetCustom(path string, content components.Customizeable, variant Variant, classes string, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Get(path, opts...), Content: content, Variant: variant}
	p.Styler.Add(classes)
	return p.Render()
}

func Link(href string, content components.Customizeable) templ.Component {
	p := &Props{Href: href, Content: content, Variant: External}
	return p.Render()
}

func Bust(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Bust(path, opts...), Content: content, Variant: Primary}
	return p.Render()
}

func Post(path string, content components.Customizeable, opts htmx.Options) templ.Component {
	p := &Props{HX: htmx.Post(path, opts), Content: content, Variant: Constructive}
	return p.Render()
}

func Put(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Put(path, opts...), Content: content, Variant: Transformative}
	return p.Render()
}

func Patch(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Patch(path, opts...), Content: content, Variant: Transformative}
	return p.Render()
}

func Delete(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Delete(path, opts...), Content: content, Variant: Destructive}
	return p.Render()
}
