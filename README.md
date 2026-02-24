# Mirage

**Generate mock APIs instantly from JSON. No backend code needed.**

Mirage is a lightweight ![Go](https://img.shields.io/badge/-%2300ADD8.svg?logo=go&logoColor=white) CLI tool that turns a simple JSON file into a fully functional HTTP API server. Perfect for frontend development, testing, prototyping, and demos.

## Why Mirage?

- **Instant setup** - Define endpoints in JSON, start the server, done
- **Zero dependencies** - Single binary, no database, no framework
- **Flexible** - Custom status codes, delays, path variables, and more
- **Fast** - Built with Go, starts in milliseconds

## Quick Example

Create `mirage.json`:

```json
{
  "endpoints": [
    {
      "method": "GET",
      "path": "/api/users",
      "response": {"users": [{"id": 1, "name": "Alice"}]}
    },
    {
      "method": "GET",
      "path": "/api/users/{id}",
      "response": {"id": "{id}", "name": "Alice"}
    }
  ]
}
```

Run `mirage serve`. Call `GET /api/users/32` → the response will have `"id": 32`. Any `{varName}` in the response is replaced by the URL value.

Your API is live on `http://localhost:8080` ✨

## Installation & Usage

```bash
# Build
go build -o mirage main.go

# Start with your config
mirage serve whatever-you-want.json

# Or use the built-in example
mirage serve --example

# Custom port
mirage serve --port=3000
```

## What's Inside?

- **Path variables** - `/users/{id}` matches any ID; put `"{id}"` in the response and it’s replaced by the URL value (numbers/decimals typed correctly)
- **Custom status codes** - Return 201, 404, 500, etc.
- **Response delays** - Simulate slow networks
- **Built-in health check** - `/health` endpoint always available
- **Request logging** - a `mirage-logs-<timestamp>.txt` file is created on startup (1 line per request) + `GET /logs` returns logs as JSON
- **Auto-config detection** - Just run `mirage serve` if `mirage.json` exists

## Learn More

For complete documentation, examples, and advanced features:

```bash
mirage guide-en   # Generate English guide
mirage guide-fr   # Generate French guide
```

This creates `mirage-guide-en.md` or `mirage-guide-fr.md` with everything you need to know.

---

Made with ❤️ by [Myself](https://github.com/cmtdrt)