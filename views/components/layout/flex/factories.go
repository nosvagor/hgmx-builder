package flex

// Center creates props for a layout that centers its children both vertically and horizontally.
func Center() *Props {
	return &Props{
		Justify: "center",
		Items:   "center",
		Gap:     4,
	}
}

// Row creates props for a horizontal layout with items aligned to the start.
func Row() *Props {
	return &Props{
		Direction: "row",
		Justify:   "start",
		Items:     "center",
		Gap:       4,
	}
}

// Col creates props for a vertical layout.
func Col() *Props {
	return &Props{
		Direction: "col",
		Justify:   "start",
		Items:     "stretch",
		Gap:       4,
	}
}
