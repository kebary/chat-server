package main

import (
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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

func routes() {
  Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	Echo.POST("/login", authenticate)
}
