# CORS v2

This module provides configurable CORS header functionality for HTTP responses.

## Usage

```go
import "github.com/DakshithaS/go-multi-major-modules/cors/v2"

options := cors.CORSOptions{
	AllowedOrigins: []string{"http://example.com"},
	AllowedMethods: []string{"GET", "POST"},
	AllowedHeaders: []string{"Content-Type"},
}

cors.AddCORSHeadersWithOptions(w, options)
```