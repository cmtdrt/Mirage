# Mirage

Mirage is a lightweight CLI tool that instantly generates mock APIs from a simple JSON configuration, allowing developers to simulate real API calls, prototype, and test endpoints without writing any backend code.

## Run

```bash
make start
```

Or manually:

```bash
go build -o mirage.exe main.go
./mirage.exe serve mirage.json
```

### Command Options

- `mirage serve <config.json>` - Run with a specific config file
- `mirage serve` - Automatically looks for `mirage.json` in the current directory
- `mirage serve --example` - Creates and runs with an example configuration file
- `mirage serve --port=8081` - Run on a custom port (default: 8080)
- Flags can be combined: `mirage serve --example --port=8081`

## Config

JSON file with an `endpoints` array. Each endpoint supports:

- `method` (required) - HTTP method: `GET`, `POST`, `PUT`, `DELETE`, etc.
- `path` (required) - URL path, supports path variables: `/users/{id}`, `/posts/{id}/comments/{commentId}`
- `response` (required) - Response body (string, object, or array)
- `description` (optional) - Description shown when server starts
- `status` (optional) - HTTP status code (default: 200)
- `delay` (optional) - Delay in milliseconds before sending response

## Example

Given this `mirage.json`:

```json
{
  "endpoints": [
    {
      "method": "GET",
      "description": "Just saying hello",
      "path": "/hello",
      "response": "Hi there 👋"
    },
    {
      "method": "GET",
      "path": "/api/v1/users",
      "response": {
        "users": [
          {
            "id": 1,
            "username": "cmtdrt",
            "email": "cmtdrt@example.com",
            "firstName": "Clément",
            "lastName": "Drt",
            "role": "ADMIN",
            "isActive": true
          }
        ]
      }
    },
    {
      "method": "POST",
      "description": "Create a new user",
      "path": "/api/v1/users",
      "status": 201,
      "delay": 1500,
      "response": {
        "id": 42,
        "username": "newuser",
        "email": "newuser@example.com",
        "message": "User created successfully"
      }
    },
    {
      "method": "GET",
      "description": "Get user by ID",
      "path": "/api/v1/users/{id}",
      "response": {
        "id": 1,
        "username": "cmtdrt",
        "email": "cmtdrt@example.com"
      }
    }
  ]
}
```

This creates the following endpoints:

| Method | Path | Description |
|--------|------|-------------|
| GET | `/hello` | Just saying hello |
| GET | `/api/v1/users` | Returns a list of users |
| POST | `/api/v1/users` | Create a new user (status: 201, delay: 1.5s) |
| GET | `/api/v1/users/{id}` | Get user by ID (matches any ID value) |

### Path Variables

Mirage supports path variables using `{variableName}` syntax. For example:

- `/users/{id}` matches `/users/1`, `/users/42`, `/users/abc`, etc.
- `/posts/{postId}/comments/{commentId}` matches `/posts/123/comments/456`

The path variables are automatically handled by Go's built-in router.

**What each endpoint returns:**

- **GET /hello**  
  `"Hi there 👋"`

- **GET /api/v1/users**  
  ```json
  {
    "users": [
      {
        "id": 1,
        "username": "cmtdrt",
        "email": "cmtdrt@example.com",
        "firstName": "Clément",
        "lastName": "Drt",
        "role": "ADMIN",
        "isActive": true
      }
    ]
  }
  ```

- **POST /api/v1/users**  
  Status: `201 Created`  
  Delay: `1500ms` (1.5 seconds)  
  Response:
  ```json
  {
    "id": 42,
    "username": "newuser",
    "email": "newuser@example.com",
    "message": "User created successfully"
  }
  ```

- **GET /api/v1/users/{id}**  
  Matches: `/api/v1/users/1`, `/api/v1/users/42`, `/api/v1/users/abc`, etc.  
  Response:
  ```json
  {
    "id": 1,
    "username": "cmtdrt",
    "email": "cmtdrt@example.com"
  }
  ```

## Quick Start

### Using the Example

```bash
mirage serve --example
```

This creates `mirage.example.json` with sample endpoints and starts the server.

### Custom Port

```bash
mirage serve mirage.json --port=3000
```

### Auto-detect Config

If you have a `mirage.json` file in your current directory, you can simply run:

```bash
mirage serve
```