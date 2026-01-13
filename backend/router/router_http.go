package router

import (
	"github.com/andrewchababi/pricecare/backend/auth"
	httpHandlers "github.com/andrewchababi/pricecare/backend/http_handlers"
	"github.com/labstack/echo/v4"
)

func registerHTTPHandlers(e *echo.Echo) {

	e.Static("/", "web/static")

	e.GET("/", httpHandlers.CalculatorPage, auth.RedirectMiddleware(STAFF, ADMIN))
	e.GET("/login", httpHandlers.LoginPage, auth.RedirectMiddleware(GUEST))

	e.POST("/api/calculate", httpHandlers.CalculatePanelPrice)
	e.POST("/api/login", httpHandlers.Login)
	e.POST("/api/logout", httpHandlers.Logout)

}
