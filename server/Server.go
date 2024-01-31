package server

import (
	"github.com/gofiber/fiber/v2"
	"go-blogsite/config"
)

func StartServer(config config.AppConfig) {
	fiberApp := fiber.New()

	// Routes
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("~/server/index.html")
	})

	//Listen and Serve
	fiberApp.Listen(config.ServerPort)
}
