# Mirage

Mirage is a lightweight CLI tool that instantly generates mock APIs from a simple JSON configuration, enabling developers to prototype and test endpoints without writing backend code.

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
