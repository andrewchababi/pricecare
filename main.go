package main

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/logger"
	"github.com/andrewchababi/pricecare/backend/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Printf("Starting Server")

	e := echo.New()

	logger.RegisterLogger(e)
	router.RegisterRouter(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true, // CRITICAL for cookies
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))

	e.Logger.Fatal(e.Start(config.ServerUrl))
}
