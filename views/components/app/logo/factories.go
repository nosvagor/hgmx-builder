package logo

import "github.com/a-h/templ"

func Logo(svg templ.Component) *Props {
	return &Props{SVG: svg}
}

func Text(name string) *Props {
	return &Props{Name: name}
}
