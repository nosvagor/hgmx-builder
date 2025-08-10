package web

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views"
	"github.com/nosvagor/hgmx-builder/views/components/app/logo"
	"github.com/nosvagor/hgmx-builder/views/components/app/source"
	"github.com/nosvagor/hgmx-builder/views/components/button"
	"github.com/nosvagor/hgmx-builder/views/components/icon"
	"github.com/nosvagor/hgmx-builder/views/components/nav"
	"github.com/nosvagor/hgmx-builder/views/icons"
	"github.com/nosvagor/hgmx-builder/views/utilities/htmx"
)

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	res := ctx.Response()
	req := ctx.Request()

	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	cacheControl := res.Header().Get("Cache-Control")
	if cacheControl != "" && !strings.Contains(cacheControl, "no-cache") && !strings.Contains(cacheControl, "no-store") {
		hash := sha1.New()
		hash.Write(buf.Bytes())
		etag := `"` + hex.EncodeToString(hash.Sum(nil)) + `"`

		if match := req.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				return ctx.NoContent(http.StatusNotModified)
			}
		}
		res.Header().Set("ETag", etag)
	}

	return ctx.HTML(statusCode, buf.String())
}

func Page(c echo.Context, content templ.Component, title string, maxAge ...int) error {
	req := c.Request()
	res := c.Response()

	hxRequest := req.Header.Get("HX-Request") == "true"

	if len(maxAge) > 0 {
		maxAge := maxAge[0]
		switch maxAge {
		case -1:
			res.Header().Set("Cache-Control", "no-store")
		case 0:
			res.Header().Set("Cache-Control", "no-cache")
		default:
			res.Header().Set("Cache-Control", fmt.Sprintf("private, max-age=%d", maxAge))
		}
	}

	if t, ok := c.Get("title").(string); ok && t != "" {
		title = t
	}

	res.Header().Add("Vary", "HX-Request")

	if hxRequest {
		return OK(c, views.Main(content, title))
	}

	return OK(c, views.Full(navbar(), views.Main(content, title)))
}

func navbar() templ.Component {
	opts := []htmx.Options{htmx.Target("#main"), htmx.Swap("outerHTML"), htmx.PushURL("true")}

	p := &nav.Props{
		Logo:     button.GetCustom("/", logo.Icon("text-surface-50", "text-[2rem]"), button.Primary, "pl-1 pr-1", opts...),
		Settings: button.Get("/settings", icon.Text(icons.Settings(), "Settings"), opts...),
		Account:  button.Get("/account", icon.Text(icons.User(), "Account"), opts...),
		Source:   source.Source("github.com/nosvagor/hgmx-builder", icons.BrandGithub(), "15.3k"),
	}

	bookmarks := []nav.Link{
		{Path: "docs", Icon: icons.Scroll()},
		{Path: "palette", Icon: icons.Palette()},
		{Path: "icons", Icon: icons.Orbit()},
		{Path: "blog", Icon: icons.Rss()},
		{Path: "faq", Icon: icons.Question(), Label: "FAQ"},
		{Path: "users", Icon: icons.Users(), Label: "Users"},
	}

	for _, link := range bookmarks {
		p.Bookmarks = append(p.Bookmarks,
			button.GetCustom(
				"/"+link.Path,
				icon.Text(link.Icon, link.Text()),
				button.Primary,
				"justify-start",
			),
		)
	}

	return p.Render()
}

func OK(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusOK, templ)
}

func BadRequest(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusBadRequest, templ)
}

func Forbidden(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusForbidden, templ)
}

func Unauthorized(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusUnauthorized, templ)
}

func Error(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusInternalServerError, templ)
}

func NotImplemented(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusNotImplemented, templ)
}
