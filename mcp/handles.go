package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

func (m *MCP) LoginStatus(ctx context.Context, req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	err := m.JueJin.Login()

	return nil, nil, err
}

func (m *MCP) PublishArticle(ctx context.Context, req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	err := m.JueJin.PublishArticle()

	return nil, nil, err
}
