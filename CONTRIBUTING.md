# Contributing

Thanks for your interest in contributing to **lea**!

## Getting Started

```bash
go mod download
go test ./...
```

## Development Workflow

- Run tests before pushing:

```bash
go test -race -cover ./...
```

- Run linting (CI uses golangci-lint):

```bash
golangci-lint run ./...
```

- Keep code formatted:

```bash
go fmt ./...
```

## Pull Requests

- Keep PRs focused and small where possible.
- Include context and any relevant benchmarks.
- Update docs when behavior changes.

## Reporting Issues

Use GitHub Issues with clear reproduction steps and expected behavior.
