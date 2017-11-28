package main

import (
	"github.com/labstack/echo"
	"net/http"
	. "app/routes"
)

var Echo = echo.New()

func main() {
	Test()
	Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	Echo.Logger.Fatal(Echo.Start(":1323"))
}
