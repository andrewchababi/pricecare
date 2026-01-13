package httpHandlers

import (
	"log"
	"net/http"

	"github.com/andrewchababi/pricecare/backend/auth"
	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/labstack/echo/v4"
)

type loginRequestWrapper struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req loginRequestWrapper
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, loginToken := auth.Login(req.Username, req.Password)
	if loginToken != "" {
		c.SetCookie(&http.Cookie{
			Name:   "loginToken",
			Value:  loginToken,
			Path:   "/",
			MaxAge: config.LoginTokenCookieDuration,
		})
	}

	return c.JSON(http.StatusOK, user)
}

func Logout(c echo.Context) error {
	log.Println("Logout http handler called [+]")
	loginToken, err := c.Cookie("loginToken")

	if err == nil {
		auth.Logout(loginToken.Value)
	}

	c.SetCookie(&http.Cookie{
		Name:   "loginToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	return c.NoContent(http.StatusOK)
}
