package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nosvagor/hgmx-builder/internal/handlers/web/pages"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == http.StatusNotFound {
		if notFoundErr := pages.NotFound(c); notFoundErr != nil {
			c.Logger().Error(notFoundErr)
		}
		return
	}

	// Extend custom error handling here, or fall back to Echo's default error handler.
	if !c.Response().Committed {
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
