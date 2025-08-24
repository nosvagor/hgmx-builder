package pages

import (
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/handlers/web"
	"github.com/nosvagor/hgmx-builder/views/icons"
	view "github.com/nosvagor/hgmx-builder/views/pages/icons"
)

var Registry = map[string]*icons.Icon{
	"github":       icons.Github(),
	"orbit":        icons.Orbit(),
	"palette":      icons.Palette(),
	"question":     icons.Question(),
	"rss":          icons.Rss(),
	"scroll":       icons.Scroll(),
	"settings":     icons.Settings(),
	"user check":   icons.UserCheck(),
	"user circle":  icons.UserCircle(),
	"user cog":     icons.UserCog(),
	"user contact": icons.UserContact(),
	"user minus":   icons.UserMinus(),
	"user pen":     icons.UserPen(),
	"user plus":    icons.UserPlus(),
	"user search":  icons.UserSearch(),
	"user star":    icons.UserStar(),
	"user x":       icons.UserX(),
	"user":         icons.User(),
	"users":        icons.Users(),
}

var fullList = make([]*icons.Icon, 0, len(Registry))

func init() {
	for _, icon := range Registry {
		fullList = append(fullList, icon)
	}
	sort.Slice(fullList, func(i, j int) bool { return fullList[i].Name < fullList[j].Name })
}

func Icons(c echo.Context) error {
	return web.Page(c, view.Main(fullList), "Icons", Preload)
}

func Search(c echo.Context) error {
	q := strings.ToLower(c.QueryParam("q"))
	limit := 100
	results := make([]*icons.Icon, 0, limit)
	for _, icon := range fullList {
		if q == "" ||
			strings.Contains(strings.ToLower(icon.Name), q) ||
			contains(icon.Aliases, q) ||
			containsIn(icon.Tags, q) {
			results = append(results, icon)
			if len(results) == limit {
				return web.OK(c, view.List(results))
			}
		}
	}
	return web.OK(c, view.List(results))
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if strings.ToLower(s) == target {
			return true
		}
	}
	return false
}

func containsIn(list []string, fragment string) bool {
	if fragment == "" {
		return false
	}
	f := strings.ToLower(fragment)
	for _, s := range list {
		if strings.Contains(strings.ToLower(s), f) {
			return true
		}
	}
	return false
}
