package htmx

import (
	"strings"
)

func Get(url string, opts ...Options) *Props {
	opts = append(opts, Preload("mousedown"))
	return Init(GET, url, opts...)
}

func GetExternal(url string, opts ...Options) *Props {
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	return Init(GET, url, opts...)
}

func Post(url string, opts ...Options) *Props {
	return Init(POST, url, opts...)
}

func Put(url string, opts ...Options) *Props {
	return Init(PUT, url, opts...)
}

func Patch(url string, opts ...Options) *Props {
	return Init(PATCH, url, opts...)
}

func Delete(url string, opts ...Options) *Props {
	return Init(DELETE, url, opts...)
}
