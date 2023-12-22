package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %v \n", r.Method, r.URL.Path, time.Since(start))
	}
	return http.HandlerFunc(fn)
}
