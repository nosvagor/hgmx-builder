package footer

import "github.com/a-h/templ"

type Props struct {
}

func Init() templ.Component {
	p := &Props{}
	return p.Render()
}
