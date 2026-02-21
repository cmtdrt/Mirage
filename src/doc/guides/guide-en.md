# Mirage — User Guide

Mirage is a lightweight CLI tool that generates mock APIs from a simple JSON configuration. It lets you simulate real API calls, prototype frontends, and test integrations without writing backend code.

---

## What is Mirage?

Mirage reads a JSON file that describes HTTP endpoints (method, path, response body, optional status and delay). It starts a local HTTP server that serves those endpoints. Any client (browser, Postman, curl, or your app) can call them like a real API.

**Use cases:**
- Frontend development before the real API exists
- Integration tests with predictable responses
- Demos and prototypes
- Learning and experimenting with HTTP APIs

---

## Installation & Run

### Build

```bash
go build -o mirage main.go
```

Or use the Makefile:

```bash
make build
```

### Start the server

```bash
mirage serve <config.json>
```

Example:

```bash
mirage serve mirage.json
```

The server runs on **port 8080** by default.

---

## Command-line options

| Option | Description |
|--------|-------------|
| `mirage serve <config.json>` | Use the given JSON file as configuration |
| `mirage serve` | Look for `mirage.json` in the current directory; error if not found |
| `mirage serve --example` | Create `mirage.example.json` from the built-in example and use it |
| `mirage serve --port=8081` | Run on port 8081 (default: 8080) |
| `mirage guide-en` | Generate this guide in English as `mirage-guide-en.md` (then exit) |
| `mirage guide-fr` | Generate this guide in French as `mirage-guide-fr.md` (then exit) |

Flags for `serve` can be combined, in any order:

```bash
mirage serve --example --port=3000
```

---

## Configuration file

The config file is JSON with a single top-level key: **`endpoints`**, an array of endpoint objects.

### Required fields (per endpoint)

These fields must be present in every endpoint:

| Field | Description |
|-------|-------------|
| `method` | HTTP method: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, etc. |
| `path` | URL path (e.g. `/api/users`). Supports path variables: `/users/{id}` |
| `response` | Response body: string, object, or array (JSON) |

### Optional fields (per endpoint)

These fields can be omitted; defaults or no effect apply when absent:

| Field | Description |
|-------|-------------|
| `description` | Short description printed when the server starts |
| `status` | HTTP status code (default: 200) |
| `delay` | Delay in **milliseconds** before sending the response |

### Minimal example

```json
{
  "endpoints": [
    {
      "method": "GET",
      "path": "/hello",
      "response": "Hello world"
    }
  ]
}
```

This exposes `GET /hello` and returns the string `"Hello world"` with status 200.

---

## Features in detail

### 1. Static responses

Define any JSON (or string) as `response`. It is sent as-is with `Content-Type: application/json` (or as a string).

```json
{
  "method": "GET",
  "path": "/api/config",
  "response": {
    "theme": "dark",
    "version": "1.0"
  }
}
```

### 2. Custom status code

Use `status` for error or success codes (e.g. 201, 404, 500):

```json
{
  "method": "POST",
  "path": "/api/users",
  "status": 201,
  "response": { "id": 1, "username": "newuser" }
}
```

### 3. Response delay

Use `delay` (in milliseconds) to simulate slow networks or backend latency:

```json
{
  "method": "GET",
  "path": "/api/slow",
  "delay": 2000,
  "response": { "message": "This took 2 seconds" }
}
```

### 4. Path variables

Paths can include placeholders with `{name}`. They match any value in that segment.

Examples:
- `/users/{id}` → matches `/users/1`, `/users/42`, `/users/abc`
- `/posts/{postId}/comments/{commentId}` → matches `/posts/10/comments/5`

**Using path values in the response:** any string value in the response that is exactly `"{varName}"` (e.g. `"{id}"`) is replaced by the value from the URL. So `GET /api/v1/users/32` can return a user with `"id": 32`. Numbers and decimals are typed correctly (REST convention: integer → number, decimal → number, else → string).

```json
{
  "method": "GET",
  "path": "/api/v1/users/{id}",
  "response": {
    "id": "{id}",
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

Calling `GET /api/v1/users/32` returns `{"id": 32, "username": "johndoe", ...}`; calling `GET /api/v1/users/alice` returns `{"id": "alice", ...}`. Path variables are matched by the router (Go 1.22+).

### 5. Descriptions

Use `description` to document endpoints. Descriptions are printed in the console when the server starts:

```json
{
  "method": "GET",
  "path": "/health",
  "description": "Health check endpoint",
  "response": { "status": "ok" }
}
```

---

## Quick start

### 1. Run with the built-in example

```bash
mirage serve --example
```

This creates `mirage.example.json` and starts the server with it. Handy to see a full sample.

### 2. Run with a custom port

```bash
mirage serve mirage.json --port=3000
```

### 3. Rely on default config file

If `mirage.json` exists in the current directory:

```bash
mirage serve
```

---

## Generating this guide

You can generate the user guide in the current directory (no server is started):

- **English:** `mirage guide-en` → creates `mirage-guide-en.md` and exits
- **French:** `mirage guide-fr` → creates `mirage-guide-fr.md` and exits

---

## Summary

- **Config:** one JSON file with an `endpoints` array.
- **Endpoints:** `method`, `path`, `response`; optionally `description`, `status`, `delay`.
- **Paths:** use `{variableName}` for dynamic segments; put `"{varName}"` in the response to inject the URL value (typed as number or string).
- **CLI:** `serve`, `--example`, `--port=…`, and `guide-en` / `guide-fr` for generating this guide.

For more examples, run `mirage serve --example` and open `mirage.example.json`.
