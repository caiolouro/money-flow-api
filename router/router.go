package router

import (
	"github.com/caiolouro/money-flow-api/handler"
	"github.com/caiolouro/money-flow-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.Auth())

	api.Get("/", handler.GetHelloWorld)

	api.Post("/debit", handler.CreateDebit)
}
