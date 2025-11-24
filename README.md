# Profiled Number Adder

This project is a simple HTTP service that provides an API to add two numbers. It is implemented in Go and focuses on performance profiling and optimization using Go's built-in profiling tools and benchmarks.

## Features

- Addition of two integers via an HTTP POST endpoint `/numbers/add-up`.
- Performance profiling with `pprof` and Go trace.
- Benchmarks for core business logic (addition) and HTTP handler.
- Structured logging using `zerolog` with log rotation and adjustable log levels.
- Graceful shutdown handling on system signals.
- Basic configuration management using YAML.

## Getting Started

### Prerequisites

- Docker and Docker Compose (optional, for containerized environment)

### Running

1. Clone the repository:

```bash
git cloneg github.com/avraam311/profiled-numbers-adder
cd profiled-number-adder
```

2. Start service:

```bash
make up
```

### Configuration

Configuration is managed via `config/local.yaml`. You can set:

- `server.port`: Port for the HTTP server to listen on.
- `debug`: Enable or disable debug logging.

Example:

```yaml
server:
  port: ":8080"
debug: true
```

### API

- **POST** `/numbers/add-up`

Request JSON body:

```json
{
  "num1": 5,
  "num2": 7
}
```

Response:

- `201 Created` with the sum of the two numbers (integer)

### Profiling and Benchmarks

- Profiling endpoints are available at `/debug/pprof`.

- CPU, heap, goroutine, threadcreate profiles and execution trace are recorded in `pprof/` and `trace/` folders.

- Benchmarks are in the `benchmarks` folder and can be run with:

```bash
go test -bench=. ./benchmarks -run=^$
```

### Logging

- Uses `zerolog` for performant structured logging.

- Logs are output both to console and file (`/app/logs/app.log`) with rotation and compression enabled to reduce disk usage.

- Logging level controlled via config (`debug` flag).

### Graceful Shutdown

The service listens for termination signals (`SIGINT, SIGTERM`) and performs a graceful shutdown with a timeout of 5 seconds.

## Optimization Summary

- Added logging rotation and compression using `lumberjack`.
- Adjusted HTTP server timeouts for better resource usage.
- Minimized logging overhead by logging only errors in hot paths.
- Improved logger initialization for dynamic log levels.
- Implemented graceful shutdown and signal handling.
