package addheader_test

import (
	"net/http/httptest"
	"testing"

	addheader "github.com/dakshithas/go-multi-major-modules/add-header/v2/src"
)

func TestAddHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	headers := map[string]string{
		"X-Test1": "value1",
		"X-Test2": "value2",
	}
	addheader.AddHeaders(w, headers)
	if w.Header().Get("X-Test1") != "value1" {
		t.Errorf("Expected X-Test1 header to be value1, got %s", w.Header().Get("X-Test1"))
	}
	if w.Header().Get("X-Test2") != "value2" {
		t.Errorf("Expected X-Test2 header to be value2, got %s", w.Header().Get("X-Test2"))
	}
}
