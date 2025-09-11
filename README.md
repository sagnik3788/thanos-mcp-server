# thanos-mcp-server

> ⚠️ **NOT READY FOR PRODUCTION** - This is a development/experimental project and should not be used in production environments.

A Thanos MCP server that provides AI agents with a single endpoint to query and analyze globally aggregated, long-term Prometheus metrics across your entire infrastructure.

## Features

- **PromQL Query Tool**: Execute PromQL queries against Thanos/Prometheus endpoints
- **MCP Integration**: Works seamlessly with MCP-compatible clients like Cursor
- **Global Metrics Access**: Query aggregated metrics across your entire monitoring infrastructure

## Prerequisites

- Thanos Query or Prometheus running on `localhost:9090`
- MCP-compatible client (e.g., Cursor)

## Installation

1. Clone this repository
2. Build the server: `go build -o thanos-mcp.exe src/main.go`
3. Configure in your MCP client (see example below)

## Configuration

Add to your MCP client configuration:

```json
{
  "mcpServers": {
    "thanos-mcp-server": {
      "command": "/path/to/thanos-mcp.exe",
      "args": []
    }
  }
}
```

## Usage

The server provides a `query` tool that accepts PromQL queries:

**Example Query:**
```promql
topk(10, sum by (job) (process_resident_memory_bytes{job!=""}))
```

This returns the top 10 jobs with highest memory usage.

## Example

![Example Usage](assets/Screenshot%202025-09-12%20000443.png)

## Development Status

This project is currently in active development. Features may change and breaking changes are expected. Use at your own risk.

## License

See [LICENSE](LICENSE) file for details.