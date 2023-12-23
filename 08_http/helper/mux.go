package helper

import (
	"fmt"
	"net/http"
	"regexp"
)

//http method
type method struct {
	Func    http.HandlerFunc
	IdNames []string
}

type handler struct {
	Methods map[string]*method
}

type mux struct {
	Handlers  map[string]*handler
	getValues *regexp.Regexp   //Regex for get the values from Path
	getIds    *regexp.Regexp   //Regex for get the Ids names from Path
	NotFound  http.HandlerFunc //Not found method
	IdsNum    int              //Default value is 2
}

func (h *handler) newMethod(httpMethod string, hf http.HandlerFunc, ids []string) {
	if h.Methods[httpMethod] == nil {
		h.Methods[httpMethod] = &method{IdNames: make([]string, len(ids))}
	}
	h.Methods[httpMethod].Func = hf
	for i, id := range ids {
		//Remove the braces, from {id} to only id
		h.Methods[httpMethod].IdNames[i] = id[1 : len(id)-1]
	}
}

func newHandler() *handler {
	return &handler{
		Methods: make(map[string]*method),
	}
}

func newmux() *mux {
	return &mux{
		Handlers:  make(map[string]*handler),
		getValues: regexp.MustCompile("[0-9]+"),
		getIds:    regexp.MustCompile("{[A-Z-a-z]+}"),
		NotFound:  notFound,
		IdsNum:    2,
	}
}

func (a *mux) Get(pattern string, hf http.HandlerFunc) {
	a.method(pattern, hf, http.MethodGet)
}

func (a *mux) Post(pattern string, hf http.HandlerFunc) {
	a.method(pattern, hf, http.MethodPost)
}

func (a *mux) Put(pattern string, hf http.HandlerFunc) {
	a.method(pattern, hf, http.MethodPut)
}

func (a *mux) Delete(pattern string, hf http.HandlerFunc) {
	a.method(pattern, hf, http.MethodDelete)
}

func (a *mux) method(pattern string, hf http.HandlerFunc, method string) {
	ids := a.getIds.FindAllString(pattern, a.IdsNum)
	//Replace all the custom ids for a default id name
	pattern = a.getIds.ReplaceAllString(pattern, "{id}")
	if a.Handlers[pattern] == nil {
		a.Handlers[pattern] = newHandler()
	}
	a.Handlers[pattern].newMethod(method, hf, ids)
}

func (h *handler) handlerMethod(w http.ResponseWriter, r *http.Request) {
	if h.Methods[r.Method] != nil {
		h.Methods[r.Method].Func.ServeHTTP(w, r)
		return
	}
	h.notAllowed(w, r)
}

func (h *handler) handlerValues(w http.ResponseWriter, r *http.Request, v []string) {
	if h.Methods[r.Method] != nil {
		r.ParseForm()
		for i, id := range h.Methods[r.Method].IdNames {
			r.Form.Add(id, v[i])
		}
		h.Methods[r.Method].Func.ServeHTTP(w, r)
		return
	}
	h.notAllowed(w, r)
}

func (h *handler) notAllowed(w http.ResponseWriter, r *http.Request) {
	allowed := make([]string, len(h.Methods))
	i := 0
	for k := range h.Methods {
		allowed[i] = k
		i++
	}
	//List all the allowed(keys) methods and return on the header
	w.Header()["Allow"] = allowed
	m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
	http.Error(w, m, http.StatusMethodNotAllowed)
}

func (a *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//If the path matchs anything, the server runs the handler
	if a.Handlers[r.URL.Path] != nil {
		a.Handlers[r.URL.Path].handlerMethod(w, r)
		return
	}
	//Get a map key replacing all the values with a default id name
	key := a.getValues.ReplaceAllString(r.URL.Path, "{id}")
	if a.Handlers[key] != nil {
		//Get all the values from path
		v := a.getValues.FindAllString(r.URL.Path, a.IdsNum)
		a.Handlers[key].handlerValues(w, r, v)
		return
	}
	//If hits here, then a not found error is delivered
	a.NotFound(w, r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	NotFound(w)
}