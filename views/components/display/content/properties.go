package content

import "github.com/nosvagor/hgmx-builder/views/components"

type Props struct {
	Text      string
	Component components.Customizable
}

func (c *Props) Classes() string {
	return ""
}
