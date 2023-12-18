package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AsJson(w http.ResponseWriter, o any) {
	b, err := json.Marshal(o)
	if err != nil {
		InternalError(w, "util: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func InternalError(w http.ResponseWriter, err string) {
	log.Println(err)
	m := fmt.Sprintf("%v internal server error", http.StatusInternalServerError)
	http.Error(w, m, http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	m := fmt.Sprintf("%v not found", http.StatusNotFound)
	http.Error(w, m, http.StatusNotFound)
}

func AsJsonError(w http.ResponseWriter, o any, code int) {
	b, err := json.Marshal(o)
	if err != nil {
		InternalError(w, "util: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write(b)
	w.WriteHeader(code)
}
