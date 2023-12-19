package helper

import "net/http"

func NotFound(w http.ResponseWriter) {
	b := []byte(`{"title":"not found","status":404}`)
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(404)
	w.Write(b)
}

func InternalError(w http.ResponseWriter) {
	b := []byte(`{"title":"internal server error","status":500}`)
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(500)
	w.Write(b)
}
