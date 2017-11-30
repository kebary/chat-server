package talk

import (
	"app/db"
	"app/db/scheme"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
)

type (
	params struct {
	  Msg string `json:"msg"`
	}
)

var (
	upgrader = websocket.Upgrader{}
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

func Ws(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
