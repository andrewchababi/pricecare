package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
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
