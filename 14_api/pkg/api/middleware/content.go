package middleware

import (
	"net/http"

	"github.com/lauro-ss/practices_go/14_api/pkg/api"
)

func ContentIsJson(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" && r.Header.Get("Content-Type") != "application/json" {
			api.BadRequest(w)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
