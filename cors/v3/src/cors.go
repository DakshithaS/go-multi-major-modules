package cors

import "net/http"

// CORSHandler interface for handling CORS
type CORSHandler interface {
	HandleCORS(w http.ResponseWriter, r *http.Request)
}

// DefaultCORSHandler implements CORSHandler with default settings
type DefaultCORSHandler struct{}

// NewDefaultCORSHandler creates a new DefaultCORSHandler
func NewDefaultCORSHandler() CORSHandler {
	return &DefaultCORSHandler{}
}

// HandleCORS adds default CORS headers
func (h *DefaultCORSHandler) HandleCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}