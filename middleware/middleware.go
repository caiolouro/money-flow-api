package middleware

import (
	"github.com/caiolouro/money-flow-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Auth() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("HTTP_USERNAME"): config.Config("HTTP_PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}
