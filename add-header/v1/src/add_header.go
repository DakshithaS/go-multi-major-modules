package addheader

import "net/http"

// AddHeader adds a single header to the response.
func AddHeader(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}
