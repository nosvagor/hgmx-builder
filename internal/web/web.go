package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/views"
)

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}
	return ctx.HTML(statusCode, buf.String())
}

func Page(c echo.Context, content templ.Component, title string) error {
	req := c.Request()
	hxRequest := req.Header.Get("HX-Request") == "true"
	hxRefresh := req.Header.Get("HX-Refresh") == "true"
	hxBoosted := req.Header.Get("HX-Boosted") == "true"
	hxHistoryRestoreRequest := req.Header.Get("HX-History-Restore-Request") == "true"

	if t, ok := c.Get("title").(string); ok && t != "" {
		title = t
	}

	if !hxRequest || hxRefresh || hxHistoryRestoreRequest || !hxBoosted {
		return OK(c, views.Full(views.Main(content, title), NewNavbar(c)))
	}
	return OK(c, views.Main(content, title))
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

func NotFound(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusNotFound, templ)
}

func Error(c echo.Context, templ templ.Component) error {
	return render(c, http.StatusInternalServerError, templ)
}

func Redirect(c echo.Context, url string) error {
	return c.Redirect(http.StatusSeeOther, url)
}
