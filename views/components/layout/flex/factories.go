package flex

import "github.com/a-h/templ"

// Center creates props for a layout that centers its children both vertically and horizontally.
func Center() templ.Component {
	p := &Props{
		Justify: "center",
		Items:   "center",
		Gap:     4,
	}
	return p.Render()
}

func Between() templ.Component {
	p := &Props{
		Justify: "between",
		Items:   "center",
	}
	return p.Render()
}

// Row creates props for a horizontal layout with items aligned to the start.
func Row() templ.Component {
	p := &Props{
		Direction: "row",
		Justify:   "start",
		Items:     "center",
		Gap:       4,
	}
	return p.Render()
}

// Col creates props for a vertical layout.
func Col() templ.Component {
	p := &Props{
		Direction: "col",
		Justify:   "start",
		Items:     "stretch",
		Gap:       4,
	}
	return p.Render()
}
