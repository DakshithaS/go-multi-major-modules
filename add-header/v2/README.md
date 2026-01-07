# add-header/v2

This is version 2 of the add-header module.

## API

- `AddHeaders(w http.ResponseWriter, headers map[string]string)`: Adds multiple headers to the response.

## Breaking Changes

- Renamed `AddHeader` to `AddHeaders` and changed to accept a map from v1.