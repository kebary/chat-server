package authenticate

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type (
	User struct {
	  Username string `json:"username"`
	  Password string `json:"password"`
	}
)

func Post(c echo.Context) (err error)  {
	u := new(User)
  if err = c.Bind(u); err != nil {
    return err
  }
	username := u.Username
	password := u.Password
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
