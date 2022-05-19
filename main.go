package main

import (
	"log"

	"github.com/caiolouro/money-flow-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
