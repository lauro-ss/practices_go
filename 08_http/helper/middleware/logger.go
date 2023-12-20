package middleware

import (
	"log"
	"net/http"
	"time"
)

type logger struct {
	handler http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v \n", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(h http.Handler) *logger {
	return &logger{h}
}
