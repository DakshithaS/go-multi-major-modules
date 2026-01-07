package cors

import "net/http"

// CORSOptions holds configuration for CORS headers
type CORSOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// AddCORSHeadersWithOptions adds CORS headers based on provided options
func AddCORSHeadersWithOptions(w http.ResponseWriter, options CORSOptions) {
	if len(options.AllowedOrigins) > 0 {
		w.Header().Set("Access-Control-Allow-Origin", options.AllowedOrigins[0]) // Simplified: just first one
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	if len(options.AllowedMethods) > 0 {
		methods := ""
		for i, method := range options.AllowedMethods {
			if i > 0 {
				methods += ", "
			}
			methods += method
		}
		w.Header().Set("Access-Control-Allow-Methods", methods)
	} else {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	}

	if len(options.AllowedHeaders) > 0 {
		headers := ""
		for i, header := range options.AllowedHeaders {
			if i > 0 {
				headers += ", "
			}
			headers += header
		}
		w.Header().Set("Access-Control-Allow-Headers", headers)
	} else {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	}

	if options.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}