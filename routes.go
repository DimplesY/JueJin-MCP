package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/mcp"
)

func SetupRoutes(app *fiber.App) {
	app.All("/mcp", mcp.MCP())
}
