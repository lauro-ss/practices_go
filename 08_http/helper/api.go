package helper

import (
	"net/http"
)

type Api struct {
	mux *mux

	f http.Handler
}

func (a *Api) Use(m func(http.Handler) http.Handler) {
	if a.f == nil {
		a.f = m(a.mux)
	} else {
		a.f = m(a.f)
	}
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
	a.f.ServeHTTP(w, r)
}
