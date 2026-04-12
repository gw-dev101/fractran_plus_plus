# Fractran++

This repository is scaffolded as a Go project for a FRACTRAN++ lexer, parser, interpreter, and compiler.

## Layout

- `internal/ast` - shared syntax tree types
- `internal/compiler` - compilation stubs
- `internal/interpreter` - execution stubs
- `internal/lexer` - tokenization
- `internal/parser` - parsing

## Testing

Run all tests with:

```bash
go test ./...
```

The repository also includes a `make test` shortcut.

## Status

The lexer and parser are implemented and tested. The compiler and interpreter are stubs that need to be implemented and tested.
