package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CreateDebit(c *fiber.Ctx) error {
	fmt.Println(c.Body()) // TODO: decode bytes to text
	return c.SendString("Received post debit")
}
