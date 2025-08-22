package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type FormatterParams struct {
	File string `json:"file" jsonschema:"the absolute path to the protobuf file"`
}

func formatFile(ctx context.Context, req *mcp.CallToolRequest, args FormatterParams) (*mcp.CallToolResult, any, error) {
	err := readFormatWrite(args.File)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "error: " + err.Error()}},
		}, nil, nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: "formatted " + args.File}},
	}, nil, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "Protobuf Formatter", Version: "0.1.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "protobuf_format",
		Description: "Formats a protobuf (.proto) file",
	}, formatFile)
	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
