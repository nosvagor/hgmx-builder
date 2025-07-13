package button

import (
	"github.com/a-h/templ"
	"github.com/nosvagor/hgmx-builder/views/components"
)

func Get(path string, content components.Customizable, extraClasses ...string) templ.Component {
	p := &Props{
		HX:           GET,
		Path:         path,
		Content:      content,
		ExtraClasses: extraClasses,
	}
	return p.Render()
}

func Post(path string, content components.Customizable, extraClasses ...string) templ.Component {
	p := &Props{
		HX:           POST,
		Path:         path,
		Content:      content,
		ExtraClasses: extraClasses,
	}
	return p.Render()
}

func Put(path string, content components.Customizable, extraClasses ...string) templ.Component {
	p := &Props{
		HX:           PUT,
		Path:         path,
		Content:      content,
		ExtraClasses: extraClasses,
	}
	return p.Render()
}

func Patch(path string, content components.Customizable, extraClasses ...string) templ.Component {
	p := &Props{
		HX:           PATCH,
		Path:         path,
		Content:      content,
		ExtraClasses: extraClasses,
	}
	return p.Render()
}

func Delete(path string, content components.Customizable, extraClasses ...string) templ.Component {
	p := &Props{
		HX:           DELETE,
		Path:         path,
		Content:      content,
		Variant:      Destructive,
		ExtraClasses: extraClasses,
	}
	return p.Render()
}
