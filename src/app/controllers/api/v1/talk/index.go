package talk

import (
	"app/db"
	"app/db/scheme"
	"github.com/labstack/echo"
	"net/http"
)

func Get(c echo.Context) error {
	var talks []scheme.Talk
	return c.JSON(http.StatusOK, db.Database.Find(&talks))
}

func Post(c echo.Context) error {
	talk := scheme.Talk{
		Msg: c.FormValue("msg"),
	}
	db.Database.Create(&talk)
	var res interface{}
	return c.JSON(http.StatusOK, res)
}
