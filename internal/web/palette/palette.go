package palette

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Main(c echo.Context) error {
	return c.String(http.StatusOK, "Palette")
}
