package mcp

import (
	"context"
	"fmt"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/browser"
	"github.com/unomcp/JueJin-MCP/juejin"
)

func loginStatusTool(ctx context.Context, _req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	fmt.Println("登录中...")
	juejin.Login(p, ctx)

	return nil, nil, nil
}
