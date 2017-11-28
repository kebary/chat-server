package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var talks []Talk
	return c.JSON(http.StatusOK, db.Find(&talks))
}

func routes() {
	Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	Echo.POST("/login", authenticate)
	Echo.GET("/talk", getTalk)
}
