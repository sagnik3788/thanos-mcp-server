package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"thanos-mcp-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// tool that allows LLMs to execute PromQL queries against Thanos
	queryTool := mcp.NewTool("query",
		mcp.WithDescription("Query all prometheus metrics"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("The query to execute"),
		),
	)

	s.AddTool(queryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		query := args["query"].(string)
		resp, err := http.Get(fmt.Sprintf("http://localhost:9090/api/v1/query?query=%s", query))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var result map[string]any
		if err := json.Unmarshal(body, &result); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.NewTextContent(string(body)),
			},
		}, nil
	})
	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
