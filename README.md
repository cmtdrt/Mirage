# Mirage

Mock API server in Golang. Reads a JSON file and exposes the defined routes on port 8080.

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

JSON file with an `endpoints` array: `method`, `path`, `response` (and optionally `description`, `status`). See `mirage.json` for an example.
