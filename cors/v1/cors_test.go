package cors_test

import (
	"net/http/httptest"
	"testing"

	cors "github.com/DakshithaS/go-multi-major-modules/cors/v1/src"
)

func TestAddCORSHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	cors.AddCORSHeaders(w)

	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin to be '*', got %s", w.Header().Get("Access-Control-Allow-Origin"))
	}
	if w.Header().Get("Access-Control-Allow-Methods") != "GET, POST, PUT, DELETE, OPTIONS" {
		t.Errorf("Expected Access-Control-Allow-Methods to be 'GET, POST, PUT, DELETE, OPTIONS', got %s", w.Header().Get("Access-Control-Allow-Methods"))
	}
	if w.Header().Get("Access-Control-Allow-Headers") != "Content-Type, Authorization" {
		t.Errorf("Expected Access-Control-Allow-Headers to be 'Content-Type, Authorization', got %s", w.Header().Get("Access-Control-Allow-Headers"))
	}
}