package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Protobuf Formatter",
		"0.1.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)
	formatterTool := mcp.NewTool("protobuf_format",
		mcp.WithDescription("Formats a protobuf (.proto) file"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("The absolute path to the protobuf file"),
		),
	)
	s.AddTool(formatterTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Using helper functions for type-safe argument access
		fileName, err := request.RequireString("file")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		err = readFormatWrite(fileName)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("formatted " + fileName), nil
	})
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
