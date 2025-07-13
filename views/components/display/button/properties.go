package button

import "github.com/nosvagor/hgmx-builder/views/components"

type HX int

const (
	GET HX = iota
	POST
	PUT
	PATCH
	DELETE
)

type Variant int

const (
	Base Variant = iota
	Primary
	Destructive
)

type Props struct {
	HX      HX
	Path    string
	Trigger string
	Target  string
	Swap    string

	Variant      Variant
	ExtraClasses []string

	Content components.Customizable
}
