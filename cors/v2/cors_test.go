package cors_test

import (
	"net/http/httptest"
	"testing"

	cors "github.com/DakshithaS/go-multi-major-modules/cors/v2/src"
)

func TestAddCORSHeadersWithOptions(t *testing.T) {
	w := httptest.NewRecorder()
	options := cors.CORSOptions{
		AllowedOrigins:   []string{"http://example.com"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}
	cors.AddCORSHeadersWithOptions(w, options)

	if w.Header().Get("Access-Control-Allow-Origin") != "http://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin to be 'http://example.com', got %s", w.Header().Get("Access-Control-Allow-Origin"))
	}
	if w.Header().Get("Access-Control-Allow-Methods") != "GET, POST" {
		t.Errorf("Expected Access-Control-Allow-Methods to be 'GET, POST', got %s", w.Header().Get("Access-Control-Allow-Methods"))
	}
	if w.Header().Get("Access-Control-Allow-Headers") != "Content-Type" {
		t.Errorf("Expected Access-Control-Allow-Headers to be 'Content-Type', got %s", w.Header().Get("Access-Control-Allow-Headers"))
	}
	if w.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials to be 'true', got %s", w.Header().Get("Access-Control-Allow-Credentials"))
	}
}