package server

import (
	"github.com/gofiber/fiber/v2"
	"go-blogsite/config"
)

func StartServer(config config.AppConfig) {
	fiberApp := fiber.New()

	fiberApp.Listen(config.ServerPort)
}
