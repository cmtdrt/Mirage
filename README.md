# Mirage

**Generate mock APIs instantly from JSON. No backend code needed.**

Mirage is a lightweight CLI tool that turns a simple JSON file into a fully functional HTTP API server. Perfect for frontend development, testing, prototyping, and demos.

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
      "response": {"id": 1, "name": "Alice"}
    }
  ]
}
```

Run:

```bash
mirage serve
```

Your API is live on `http://localhost:8080` ✨

## Installation & Usage

```bash
# Build
go build -o mirage main.go

# Start with your config
mirage serve mirage.json

# Or use the built-in example
mirage serve --example

# Custom port
mirage serve --port=3000
```

## What's Inside?

- **Path variables** - `/users/{id}` matches any ID
- **Custom status codes** - Return 201, 404, 500, etc.
- **Response delays** - Simulate slow networks
- **Built-in health check** - `/health` endpoint always available
- **Auto-config detection** - Just run `mirage serve` if `mirage.json` exists

## Learn More

For complete documentation, examples, and advanced features:

```bash
mirage guide-en   # Generate English guide
mirage guide-fr   # Generate French guide
```

This creates `mirage-guide-en.md` or `mirage-guide-fr.md` with everything you need to know.

---

**Ready to mock?** Start with `mirage serve --example` and explore the generated `mirage.example.json` file!
