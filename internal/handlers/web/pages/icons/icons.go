package icons

import (
	"sort"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nosvagor/hgmx-builder/views/components/icon/icons"
	view "github.com/nosvagor/hgmx-builder/views/pages/icons"
)

func init() {
	for _, icon := range Registry {
		fullList = append(fullList, icon)
	}
	sort.Slice(fullList, func(i, j int) bool { return fullList[i].Icon.Name < fullList[j].Icon.Name })
}

func Main(c echo.Context) templ.Component {
	icons := make([]*icons.IconDetails, 0, 100)
	for i, icon := range fullList {
		icons = append(icons, icon)
		if i == 100 {
			break
		}
	}
	return view.Main(icons)
}
