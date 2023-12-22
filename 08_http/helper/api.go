package helper

import (
	"net/http"
)

type Api struct {
	mux *mux
}

func NewApi() *Api {
	return &Api{mux: newmux()}
}

func (a *Api) Get(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodGet)
}

func (a *Api) Post(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodPost)
}

func (a *Api) Put(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodPut)
}

func (a *Api) Delete(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodDelete)
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
