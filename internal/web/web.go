package web

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views"
	"github.com/nosvagor/hgmx-builder/views/components/navigation/navbar"
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

	headers := req.Header
	for k, v := range headers {
		log.Println(k, v)
	}

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

	return OK(c, views.Full(views.Main(content, title), navbar.Init()))
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
