package middlewares

import (
	"log"
	"net/http"
)

func Logger(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL)
		h.ServeHTTP(w, r)
	}
}
