package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type FormatterParams struct {
	File string `json:"file" jsonschema:"the absolute path to the protobuf file"`
}

func formatFile(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[FormatterParams]) (*mcp.CallToolResultFor[any], error) {
	err := readFormatWrite(params.Arguments.File)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: "error: " + err.Error()}},
		}, nil
	}
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: "formatted " + params.Arguments.File}},
	}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "Protobuf Formatter", Version: "0.1.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "protobuf_format",
		Description: "Formats a protobuf (.proto) file",
	}, formatFile)
	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
