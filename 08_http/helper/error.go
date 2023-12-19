package helper

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter) {
	m := fmt.Sprintf(`{"title":"not found","status":%v}`, http.StatusNotFound)
	writeError(w, []byte(m), http.StatusNotFound)
}

func NotAllowed(w http.ResponseWriter, allowed []string) {
	m := fmt.Sprintf(`{"title":"method not allowed","status":%v, "allowed":%v}`,
		http.StatusMethodNotAllowed,
		allowed)
	writeError(w, []byte(m), http.StatusMethodNotAllowed)
}

func InternalError(w http.ResponseWriter) {
	m := fmt.Sprintf(`{"title":"internal server error","status":%v}`, http.StatusInternalServerError)
	writeError(w, []byte(m), http.StatusMethodNotAllowed)
}

func writeError(w http.ResponseWriter, b []byte, status int) {
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	w.Write(b)
}
