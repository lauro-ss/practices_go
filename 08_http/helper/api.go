package helper

import (
	"net/http"
)

type Api struct {
	mux *mux

	stackServer http.Handler
}

//Pushs the middleware on top of stack server
func (a *Api) Use(m func(http.Handler) http.Handler) {
	if a.stackServer == nil {
		a.stackServer = m(a.mux)
	} else {
		a.stackServer = m(a.stackServer)
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
	a.stackServer.ServeHTTP(w, r)
}
