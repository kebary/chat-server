package talk

import (
	"app/db"
	"app/db/scheme"
	"github.com/labstack/echo"
	"net/http"
)

type (
	params struct {
	  Msg string `json:"msg"`
	}
)

func Get(c echo.Context) error {
	var talks []scheme.Talk
	return c.JSON(http.StatusOK, db.Database.Find(&talks))
}

func Post(c echo.Context) (err error) {
	p := new(params)
	if err = c.Bind(p); err != nil {
    return err
  }
	talk := scheme.Talk{
		Msg: p.Msg,
	}
	db.Database.Create(&talk)
	var res interface{}
	return c.JSON(http.StatusOK, res)
}
