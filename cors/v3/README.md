# CORS v3

This module provides an interface-based CORS handler for HTTP responses.

## Usage

```go
import "github.com/DakshithaS/go-multi-major-modules/cors/v3"

handler := cors.NewDefaultCORSHandler()
handler.HandleCORS(w, r)
```