package components

import (
	"strings"

	"github.com/a-h/templ"
)

// Standard is a component that can be rendered.
type Standard interface {
	Render() templ.Component
}

// === Customizable Classes ========================================================================

// Styler can be embedded in a component's properties to make it stylable.
// It provides the storage and base logic for adding and building class strings.
type Styler struct {
	extraClasses []string
}

// Add stores additional CSS classes to be applied.
func (s *Styler) Add(classes ...string) {
	s.extraClasses = append(s.extraClasses, classes...)
}

// Build combines the Styler's extra classes with any base classes and returns a final class string.
func (s *Styler) Build(base ...string) string {
	all := append(s.extraClasses, base...)
	return strings.Join(all, " ")
}

// Customizeable is an interface for components that can have their styles customized fluently.
type Customizeable interface {
	Standard
	Style(classes ...string) Customizeable
}
