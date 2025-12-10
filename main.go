package main

import (
	"log/slog"
	"os"

	"github.com/unomcp/JueJin-MCP/bootstrap"
	"github.com/unomcp/JueJin-MCP/configs"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	
	app := bootstrap.NewApplication(configs.FiberConfig)
	bootstrap.Start(app)
}
