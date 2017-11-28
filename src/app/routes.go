package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func routes() {
  Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
