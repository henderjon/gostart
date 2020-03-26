package favicon

import (
	"net/http"
)

var favicon = []byte("\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\xff\xff\xff\x21\xf9\x04\x01\x00\x00\x00\x00\x2c\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x01\x44\x00\x3b")

// FaviconHandler returns a 1x1 transparent gif as a favicon `mux.Handle("/favicon.ico", FaviconHandler())`
func FaviconHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "image/gif")
		w.Write(favicon)
	})
}
