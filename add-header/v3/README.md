# add-header/v3

This is version 3 of the add-header module.

## API

- `HeaderAdder` interface with `AddHeaders(w http.ResponseWriter)`.
- `SimpleAdder` struct implementing `HeaderAdder`.
- `NewSimpleAdder(headers map[string]string) *SimpleAdder`: Creates a new SimpleAdder.

## Breaking Changes

- Replaced function with interface and struct from v2.