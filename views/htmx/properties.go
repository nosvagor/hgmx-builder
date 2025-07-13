package htmx

import (
	"maps"

	"github.com/a-h/templ"
)

type MethodType string

const (
	GET    MethodType = "get"
	POST   MethodType = "post"
	PUT    MethodType = "put"
	PATCH  MethodType = "patch"
	DELETE MethodType = "delete"
)

type Props struct {
	Method  MethodType
	URL     string
	Options Options
}

type Options templ.Attributes

func (o Options) Merge(other Options) {
	maps.Copy(o, other)
}

func (o Options) Enable(key string) {
	o[key] = true
}

func Init(method MethodType, url string, opts ...Options) *Props {
	p := &Props{
		Method:  method,
		URL:     url,
		Options: Options{},
	}
	for _, opt := range opts {
		p.Options.Merge(opt)
	}
	return p
}

func (p *Props) Opts() templ.Attributes {
	attrs := templ.Attributes{
		"hx-" + string(p.Method): p.URL,
	}
	if p.Options != nil {
		maps.Copy(attrs, p.Options)
	}

	return attrs
}
