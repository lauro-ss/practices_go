package middleware

import (
	"httpserver/helper"
	"net/http"
)

func ContentIsJson(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" && r.Header.Get("Content-Type") != "application/json" {
			helper.BadRequest(w)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
