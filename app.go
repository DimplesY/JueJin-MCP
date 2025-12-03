package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/configs"
	"github.com/unomcp/JueJin-MCP/middleware"
)

func start() {
	app := fiber.New(configs.FiberConfig)

	app.Use(middleware.Cros())

	SetupRoutes(app)

	log.Fatal(app.Listen(configs.Port))
}
