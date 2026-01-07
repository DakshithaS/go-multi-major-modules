package addheader_test

import (
	"net/http/httptest"
	"testing"

	addheader "github.com/DakshithaS/go-multi-major-modules/add-header/v1/src"
)

func TestAddHeader(t *testing.T) {
	w := httptest.NewRecorder()
	addheader.AddHeader(w, "X-Test", "value")
	if w.Header().Get("X-Test") != "value" {
		t.Errorf("Expected X-Test header to be value, got %s", w.Header().Get("X-Test"))
	}
}
