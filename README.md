# Fractran++

This repository is scaffolded as a Go project for a FRACTRAN++ lexer, parser, interpreter, and compiler.

## Layout

- `internal/ast` - shared syntax tree types
- `internal/compiler` - compilation stubs
- `internal/interpreter` - execution stubs
- `internal/lexer` - tokenization stubs
- `internal/parser` - parsing stubs

## Testing

Run all tests with:

```bash
go test ./...
```

The repository also includes a `make test` shortcut.

## Status

No FRACTRAN++ behavior is implemented yet. The packages compile as placeholders so tests can be added incrementally.
