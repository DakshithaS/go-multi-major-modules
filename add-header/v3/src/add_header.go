package addheader

import "net/http"

// HeaderAdder is an interface for adding headers.
type HeaderAdder interface {
	AddHeaders(w http.ResponseWriter)
}

// SimpleAdder implements HeaderAdder.
type SimpleAdder struct {
	headers map[string]string
}

// NewSimpleAdder creates a new SimpleAdder.
func NewSimpleAdder(headers map[string]string) *SimpleAdder {
	return &SimpleAdder{headers: headers}
}

// AddHeaders adds the headers to the response.
func (s *SimpleAdder) AddHeaders(w http.ResponseWriter) {
	for k, v := range s.headers {
		w.Header().Set(k, v)
	}
}
