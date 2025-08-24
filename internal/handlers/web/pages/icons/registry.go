package icons

import (
	"github.com/nosvagor/hgmx-builder/views/components/icon/icons"
)

var fullList = make([]*icons.IconDetails, 0, len(Registry))

var Registry = map[string]*icons.IconDetails{
	"github": {
		Icon:         *icons.Github(),
		Tags:         []string{"github", "brand", "code", "repo", "git"},
		Categories:   []string{"brands", "development", "social"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"orbit": {
		Icon:         *icons.Orbit(),
		Tags:         []string{"orbit", "circle", "rotation", "spin"},
		Categories:   []string{"shapes", "motion"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"palette": {
		Icon:         *icons.Palette(),
		Tags:         []string{"palette", "color", "art", "paint"},
		Categories:   []string{"design", "tools"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"question": {
		Icon:         *icons.Question(),
		Tags:         []string{"question", "help", "support", "unknown"},
		Categories:   []string{"communication"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"rss": {
		Icon:         *icons.Rss(),
		Tags:         []string{"rss", "feed", "news", "blog", "updates"},
		Categories:   []string{"communication", "social"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"scroll": {
		Icon:         *icons.Scroll(),
		Tags:         []string{"scroll", "docs", "text", "info"},
		Categories:   []string{"text", "files"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"settings": {
		Icon:         *icons.Settings(),
		Tags:         []string{"settings", "gear", "cog"},
		Categories:   []string{"tools"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userCheck": {
		Icon:         *icons.UserCheck(),
		Tags:         []string{"user", "check", "verify", "ok", "success"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userCircle": {
		Icon:         *icons.UserCircle(),
		Tags:         []string{"user", "circle", "account", "profile"},
		Categories:   []string{"people", "account", "shapes"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userCog": {
		Icon:         *icons.UserCog(),
		Tags:         []string{"user", "settings", "cog", "gear"},
		Categories:   []string{"people", "account", "tools"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userContact": {
		Icon:         *icons.UserContact(),
		Tags:         []string{"user", "id", "card", "contact"},
		Categories:   []string{"people", "files"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userMinus": {
		Icon:         *icons.UserMinus(),
		Tags:         []string{"user", "remove", "minus"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userPen": {
		Icon:         *icons.UserPen(),
		Tags:         []string{"user", "pen", "edit", "write", "update"},
		Categories:   []string{"people", "account", "tools"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userPlus": {
		Icon:         *icons.UserPlus(),
		Tags:         []string{"user", "add", "plus", "invite", "new"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userSearch": {
		Icon:         *icons.UserSearch(),
		Tags:         []string{"user", "search", "find", "filter"},
		Categories:   []string{"people", "tools"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userStar": {
		Icon:         *icons.UserStar(),
		Tags:         []string{"user", "star", "favorite", "featured"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userX": {
		Icon:         *icons.UserX(),
		Tags:         []string{"user", "remove", "x", "close"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"user": {
		Icon:         *icons.User(),
		Tags:         []string{"user", "person", "account", "profile"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"users": {
		Icon:         *icons.Users(),
		Tags:         []string{"users", "group", "team", "people"},
		Categories:   []string{"people", "account"},
		Contributors: []string{"nosvagor", "lucide"},
	},
}
