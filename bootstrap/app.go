package bootstrap

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/configs"
	"github.com/unomcp/JueJin-MCP/mcp"
)

type application struct {
	Configs fiber.Config
}

func (app *application) mount() *fiber.App  {
	fiberApp := fiber.New(app.Configs)
	fiberApp.All("/mcp", mcp.MCP())
	return fiberApp
}

func (app *application) run(fiber *fiber.App) {
	log.Fatal(fiber.Listen(configs.Port))
}


func NewApplication(configs fiber.Config) *application {
	return &application{
		Configs: configs,
	}
}

func Start(app *application) {
	fiber := app.mount()
	app.run(fiber)
}
