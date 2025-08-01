package include

import "github.com/a-h/templ"

func Script(path string) templ.Component {
	p := &Props{
		Html: SCRIPT,
		Path: path,
	}
	return p.Render()
}

func DeferedScript(path string) templ.Component {
	p := &Props{
		Html:  SCRIPT,
		Path:  path,
		Defer: true,
	}
	return p.Render()
}

func ModuleScript(path string) templ.Component {
	p := &Props{
		Html:   SCRIPT,
		Path:   path,
		Module: true,
	}
	return p.Render()
}

func Style(path string) templ.Component {
	p := &Props{
		Html: STYLE,
		Path: path,
	}
	return p.Render()
}
