package main

import (
	"github.com/bshome19/my-cards-go-gin-fiber/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to My Cards")
	})
	configs.ConnectDB()
	app.Listen(":5000")
}