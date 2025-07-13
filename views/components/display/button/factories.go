package button

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/htmx"
)

func Get(path string, content components.Customizeable, opts ...htmx.Options) templ.Component {
	p := &Props{HX: htmx.Get(path, opts...), Content: content, Variant: Base}
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
