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

## Config

JSON file with an `endpoints` array: `method`, `path`, `response` (and optionally `description`, `status`, `delay`).

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