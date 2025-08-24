package icons

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/handlers/web"
	"github.com/nosvagor/hgmx-builder/views/components/icon/icons"
	view "github.com/nosvagor/hgmx-builder/views/pages/icons"
)

func Search(c echo.Context) error {
	q := strings.ToLower(c.QueryParam("q"))
	limit := 100
	results := make([]*icons.IconDetails, 0, limit)
	for _, icon := range fullList {
		if q == "" ||
			strings.Contains(string(icon.Icon.Name), q) ||
			containsIn(icon.Tags, q) {
			results = append(results, icon)
			if len(results) == limit {
				return web.OK(c, view.List(results))
			}
		}
	}
	return web.OK(c, view.List(results))
}

func containsIn(list []string, fragment string) bool {
	if fragment == "" {
		return false
	}
	for _, s := range list {
		if strings.Contains(s, fragment) {
			return true
		}
	}
	return false
}
