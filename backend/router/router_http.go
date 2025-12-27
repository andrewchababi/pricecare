package router

import (
	httpHandlers "github.com/andrewchababi/pricecare/backend/http_handlers"
	"github.com/labstack/echo/v4"
)

func registerHTTPHandlers(e *echo.Echo) {

	e.Static("/", "web/static")

	e.POST("/api/calculate", httpHandlers.CalculatePanelPrice)

	e.POST("/api/login", httpHandlers.Login)
}
