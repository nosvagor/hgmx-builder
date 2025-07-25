package icon

import (
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/display/content"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
)

func Text(icon icons.Icon, text string) *Props {
	return &Props{Icon: icon, Type: ICON_TEXT, Content: content.Text(text)}
}

func Icon(icon icons.Icon) *Props {
	return &Props{Icon: icon}
}

func IconTip(icon icons.Icon, content components.Customizable) *Props {
	return &Props{Icon: icon, Tooltip: content}
}
