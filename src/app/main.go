package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Echo = echo.New()
)

func main() {
	// Middleware
	Echo.Use(middleware.Logger())
	Echo.Use(middleware.Recover())
  // routes
	routes()
  // start
	Echo.Logger.Fatal(Echo.Start(":1323"))
}
