package routes

import (
	"app/db"
	"app/db/scheme"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func authenticate(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "test" && password == "test" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "test"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("UNnx5PbFtH7b8HKazP"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func getTalk(c echo.Context) error {
	var talks []scheme.Talk
	return c.JSON(http.StatusOK, db.Database.Find(&talks))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Routes(Echo *echo.Echo) {
	Echo.GET("/", index)
	Echo.POST("/login", authenticate)
	Echo.GET("/talk", getTalk)
}
