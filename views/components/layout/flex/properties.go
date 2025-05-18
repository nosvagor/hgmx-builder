package flex

import (
	"fmt"
	"strings"
)

// Props defines the properties for a flexbox container.
type Props struct {
	Class *string // For custom overrides

	Direction string // Maps to flex-direction
	Justify   string // Maps to justify-content
	Items     string // Maps to align-items
	Wrap      string // Maps to flex-wrap
	Gap       int    // Maps to gap
}

// toClasses converts the Props struct into a slice of Tailwind CSS classes.
func (p *Props) toClasses() []string {
	classes := []string{"flex"}

	if p.Direction != "" {
		classes = append(classes, "flex-"+p.Direction)
	}
	if p.Justify != "" {
		classes = append(classes, "justify-"+p.Justify)
	}
	if p.Items != "" {
		classes = append(classes, "items-"+p.Items)
	}
	if p.Wrap != "" {
		classes = append(classes, "flex-"+p.Wrap)
	}
	if p.Gap > 0 {
		classes = append(classes, fmt.Sprintf("gap-%d", p.Gap))
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
