package routes

import (
	"app/controllers/api/v1/login"
	"app/controllers/api/v1/root"
	"app/controllers/api/v1/talk"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routes(Echo *echo.Echo) {
	Echo.GET("/", root.Get)
	Echo.POST("/login", login.Post)
	apis := Echo.Group("/api/v1")
	apis.Use(middleware.JWT([]byte("UNnx5PbFtH7b8HKazP")))
	apis.GET("/talk", talk.Get)
	apis.POST("/talk", talk.Post)
}
