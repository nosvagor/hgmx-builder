package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string, lang ...language.Tag) string {
	if len(lang) == 0 {
		lang = []language.Tag{language.English}
	}

	return cases.Title(lang[0]).String(s)
}
