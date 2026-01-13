package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/andrewchababi/pricecare/backend/auth"
	"github.com/andrewchababi/pricecare/backend/models"
	"github.com/labstack/echo/v4"
)

const GUEST = models.UserTypeNone
const STAFF = models.UserTypeStaff
const ADMIN = models.UserTypeAdminLab

func RegisterRouter(e *echo.Echo) {
	e.Use(auth.AuthMiddleware)
	registerErrorHandler(e)

	registerHTTPHandlers(e)
}

func registerErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		isNotFound := ok && he.Code == http.StatusNotFound
		if isNotFound && isHtmlRequest(c) {
			c.Redirect(http.StatusPermanentRedirect, "/")
		} else {
			e.DefaultHTTPErrorHandler(err, c)
		}
	}
}

func isHtmlRequest(c echo.Context) bool {
	request := c.Request()

	if request.Method != http.MethodGet {
		return false
	}

	for _, prefix := range []string{"/api", "/data", "/fonts", "/images", "/scripts", "/styles"} {
		if strings.HasPrefix(request.URL.Path, prefix) {
			return false
		}
	}

	if ext := path.Ext(request.URL.Path); ext != "" {
		return false
	}

	if request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		return false
	}

	accept := request.Header.Get("Accept")
	return strings.Contains(accept, "text/html") || accept == "" || strings.Contains(accept, "*/*")
}
