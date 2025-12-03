package mcp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/configs"
)

func initMCP() *goMcp.Server {
	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    configs.MCPName,
		Version: configs.MCPVersion,
	}, nil)

	// 添加登录状态工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "login status",
		Description: "获取登录状态",
	}, loginStatusTool)

	return server
}

func MCP() fiber.Handler {
	mcp := initMCP()
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return mcp
	}, configs.MCPStreamableHTTPOptions)
	return adaptor.HTTPHandler(mcpHandler)
}
