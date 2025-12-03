package configs

import goMcp "github.com/modelcontextprotocol/go-sdk/mcp"

var (
	MCPName                  = "JueJin-MCP"
	MCPVersion               = "0.0.1"
	MCPStreamableHTTPOptions = &goMcp.StreamableHTTPOptions{
		JSONResponse: true,
	}
)
