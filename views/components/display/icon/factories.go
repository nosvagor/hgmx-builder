package icon

import (
	"github.com/nosvagor/hgmx-builder/views/components"
	"github.com/nosvagor/hgmx-builder/views/components/display/content"
	"github.com/nosvagor/hgmx-builder/views/components/icons"
)

func Text(icon *icons.Icon, text string) components.Customizeable {
	return &Props{Icon: *icon, Type: ICON_TEXT, Content: content.Text(text)}
}

func Icon(icon *icons.Icon) components.Customizeable {
	return &Props{Icon: *icon}
}

func Tip(icon *icons.Icon, content components.Customizeable) components.Customizeable {
	p := &Props{Icon: *icon, Tooltip: content}
	p.Styler.Add("py-1")
	return p
}
