package main

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/logger"
	"github.com/andrewchababi/pricecare/backend/router"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Printf("Starting Server")

	e := echo.New()

	logger.RegisterLogger(e)
	router.RegisterRouter(e)

	e.Logger.Fatal(e.Start(config.ServerUrl))
}
