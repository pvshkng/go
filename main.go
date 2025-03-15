package main

import (
	//"net/http"
	//"os"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

	s := server.NewMCPServer(
		"MCP Server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	tool := mcp.NewTool(
		"stock",
		mcp.WithDescription("Retrieve stock data"),
		mcp.WithString("symbol", mcp.Description("Stock symbol"), mcp.Required()),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
        symbol := request.Params.Arguments["symbol"].(string)
        return mcp.NewToolResultText(fmt.Sprintf("%.2f", symbol)), nil
    })

	godotenv.Load()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok bb",
		})
	})

	r.Run()
}
