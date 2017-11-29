package root

import (
	"github.com/labstack/echo"
	"net/http"
)

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
