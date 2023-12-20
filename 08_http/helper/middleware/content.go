package middleware

import (
	"httpserver/helper"
	"net/http"
)

type content struct {
	handler http.Handler
}

func (c *content) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" && r.Header.Get("Content-Type") != "application/json" {
		helper.BadRequest(w)
		return
	}
	c.handler.ServeHTTP(w, r)
}

func NewContent(h http.Handler) *content {
	return &content{h}
}
