package grid

// ThreeByThree creates a 3x3 grid layout.
func ThreeByThree() *Props {
	return &Props{
		Cols:    3,
		Rows:    3,
		Gap:     4,
		Items:   "center",
		Justify: "center",
	}
}

// Columns creates a responsive grid with a specified number of columns.
func Columns(cols int) *Props {
	return &Props{
		Cols:    cols,
		Gap:     4,
		Items:   "stretch",
		Justify: "stretch",
	}
}
