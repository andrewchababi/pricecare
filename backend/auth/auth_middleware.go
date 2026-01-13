package auth

import (
	"net/http"

	"github.com/andrewchababi/pricecare/backend/models"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginToken := ""
		cookie, err := c.Cookie("loginToken")
		if err == nil {
			loginToken = cookie.Value
		}
		c.Set("loginToken", loginToken)

		user := getUserFromLoginToken(loginToken)
		c.Set("user", user)

		return next(c)
	}
}

func RedirectMiddleware(allowed ...models.UserType) echo.MiddlewareFunc {
	allowedSet := map[models.UserType]struct{}{}
	for _, t := range allowed {
		allowedSet[t] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
			c.Response().Header().Set("Expires", "0")
			c.Response().Header().Set("Pragma", "no-cache")

			user, _ := c.Get("user").(models.User)

			_, ok := allowedSet[user.UserType]
			if ok {
				return next(c)
			}

			if user.UserType == models.UserTypeNone {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
}

func BlockMiddleware(allowed ...models.UserType) echo.MiddlewareFunc {
	allowedSet := map[models.UserType]struct{}{}
	for _, t := range allowed {
		allowedSet[t] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, _ := c.Get("user").(models.User)

			_, ok := allowedSet[user.UserType]
			if ok {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	}
}
