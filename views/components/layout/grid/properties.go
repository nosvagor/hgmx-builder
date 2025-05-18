package grid

import (
	"fmt"
	"strings"
)

type Props struct {
	Class *string // For custom overrides

	Cols    int    // Maps to grid-template-columns
	Rows    int    // Maps to grid-template-rows
	Gap     int    // Maps to gap
	Flow    string // Maps to grid-auto-flow
	Justify string // Maps to justify-items
	Items   string // Maps to align-items
}

// toClasses converts the Props struct into a slice of Tailwind CSS classes.
func (p *Props) toClasses() []string {
	classes := []string{"grid"}

	if p.Cols > 0 {
		classes = append(classes, fmt.Sprintf("grid-cols-%d", p.Cols))
	}
	if p.Rows > 0 {
		classes = append(classes, fmt.Sprintf("grid-rows-%d", p.Rows))
	}
	if p.Gap > 0 {
		classes = append(classes, fmt.Sprintf("gap-%d", p.Gap))
	}
	if p.Flow != "" {
		classes = append(classes, "grid-flow-"+p.Flow)
	}
	if p.Justify != "" {
		classes = append(classes, "justify-"+p.Justify)
	}
	if p.Items != "" {
		classes = append(classes, "items-"+p.Items)
	}

	return classes
}

// String returns the full class string, including any custom overrides.
func (p *Props) Classes() string {
	allClasses := p.toClasses()
	if p.Class != nil {
		allClasses = append(allClasses, *p.Class)
	}
	return strings.Join(allClasses, " ")
}
