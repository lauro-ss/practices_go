package api

import (
	"net/http"
)

type api struct {
	mux *mux

	stackServer http.Handler
}

// Pushs the middleware on top of stack server
func (a *api) Use(m func(http.Handler) http.Handler) {
	if a.stackServer == nil {
		a.stackServer = m(a.mux)
	} else {
		a.stackServer = m(a.stackServer)
	}
}

func NewApi() *api {
	return &api{mux: newmux()}
}

// Overrides the default not found function
func (a *api) SetNotFound(hf http.HandlerFunc) {
	a.mux.notFound = hf
}

// The ids number default is 2
// if you need more then two ids examle:
// "animals/{animalId}/foods/{foodId}/type/{typeId}"
// you need to call this function with value as 3
func (a *api) SetIdNumber(v uint) {
	a.mux.idsNum = v
}

func (a *api) Get(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodGet)
}

func (a *api) Post(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodPost)
}

func (a *api) Put(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodPut)
}

func (a *api) Delete(pattern string, hf http.HandlerFunc) {
	a.mux.method(pattern, hf, http.MethodDelete)
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.stackServer.ServeHTTP(w, r)
}
