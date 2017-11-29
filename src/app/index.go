package main

import (
	"app/db/migration"
	"app/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	migration.Migration()
}

func main() {
	Echo := echo.New()
	// Middleware
	Echo.Use(middleware.Logger())
	Echo.Use(middleware.Recover())
	// routes
	routes.Routes(Echo)
	// start
	Echo.Logger.Fatal(Echo.Start(":1323"))
}
