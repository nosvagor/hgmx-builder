package components

import "github.com/a-h/templ"

type Standard interface {
	Render() templ.Component
}

type Customizable interface {
	Standard
	Classes() string
}
