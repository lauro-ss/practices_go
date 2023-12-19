package helper

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

func NewError(t string, s int) Error {
	return Error{
		Title:  t,
		Status: s,
	}
}

func AsJsonError(w http.ResponseWriter, e Error) {
	b, err := json.Marshal(e)
	if err != nil {
		InternalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(e.Status)
	w.Write(b)
}

func AsJson(w http.ResponseWriter, o any) {
	b, err := json.Marshal(o)
	if err != nil {
		InternalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
