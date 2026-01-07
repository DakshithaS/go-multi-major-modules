package cors_test

import (
	"net/http/httptest"
	"testing"

	cors "github.com/DakshithaS/go-multi-major-modules/cors/v3/src"
)

func TestDefaultCORSHandler_HandleCORS(t *testing.T) {
	handler := cors.NewDefaultCORSHandler()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	handler.HandleCORS(w, r)

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