package main

import (
	"github.com/labstack/echo"
)

var (
	Echo = echo.New()
)

func main() {
	routes()
	Echo.Logger.Fatal(Echo.Start(":1323"))
}
