package addheader

import "net/http"

// AddHeaders adds multiple headers to the response.
func AddHeaders(w http.ResponseWriter, headers map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}
