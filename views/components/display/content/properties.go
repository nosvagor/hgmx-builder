package content

import (
	"strings"

	"github.com/nosvagor/hgmx-builder/views/components"
)

type Props struct {
	Text      string
	Component components.Customizable
}

func (c *Props) Classes(extra ...string) string {
	return strings.Join(extra, " ")
}
