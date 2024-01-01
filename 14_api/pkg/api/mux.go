package helper

import (
	"fmt"
	"net/http"
	"regexp"
)

//http method
type method struct {
	handlerFunc http.HandlerFunc
	idNames     []string
}

type handler struct {
	methods map[string]*method
}

type mux struct {
	handlers  map[string]*handler
	getValues *regexp.Regexp   //Regex for get the values from Path
	getIds    *regexp.Regexp   //Regex for get the Ids names from Path
	notFound  http.HandlerFunc //Not found method
	idsNum    uint             //Default value is 2
}

func (h *handler) newMethod(httpMethod string, hf http.HandlerFunc, ids []string) {
	if h.methods[httpMethod] == nil {
		h.methods[httpMethod] = &method{idNames: make([]string, len(ids))}
	}
	h.methods[httpMethod].handlerFunc = hf
	for i, id := range ids {
		//Remove the braces, from {id} to only id
		h.methods[httpMethod].idNames[i] = id[1 : len(id)-1]
	}
}

func newHandler() *handler {
	return &handler{
		methods: make(map[string]*method),
	}
}

func newmux() *mux {
	return &mux{
		handlers:  make(map[string]*handler),
		getValues: regexp.MustCompile("[0-9]+"),
		getIds:    regexp.MustCompile("{[A-Z-a-z]+}"),
		notFound:  notFound,
		idsNum:    2,
	}
}

func (m *mux) Get(pattern string, hf http.HandlerFunc) {
	m.method(pattern, hf, http.MethodGet)
}

func (m *mux) Post(pattern string, hf http.HandlerFunc) {
	m.method(pattern, hf, http.MethodPost)
}

func (m *mux) Put(pattern string, hf http.HandlerFunc) {
	m.method(pattern, hf, http.MethodPut)
}

func (m *mux) Delete(pattern string, hf http.HandlerFunc) {
	m.method(pattern, hf, http.MethodDelete)
}

func (m *mux) method(pattern string, hf http.HandlerFunc, method string) {
	ids := m.getIds.FindAllString(pattern, int(m.idsNum))
	//Replace all the custom ids for a default id name
	pattern = m.getIds.ReplaceAllString(pattern, "{id}")
	if m.handlers[pattern] == nil {
		m.handlers[pattern] = newHandler()
	}
	m.handlers[pattern].newMethod(method, hf, ids)
}

func (h *handler) handlerMethod(w http.ResponseWriter, r *http.Request) {
	if h.methods[r.Method] != nil {
		h.methods[r.Method].handlerFunc.ServeHTTP(w, r)
		return
	}
	h.notAllowed(w, r)
}

func (h *handler) handlerValues(w http.ResponseWriter, r *http.Request, v []string) {
	if h.methods[r.Method] != nil {
		r.ParseForm()
		for i, id := range h.methods[r.Method].idNames {
			r.Form.Add(id, v[i])
		}
		h.methods[r.Method].handlerFunc.ServeHTTP(w, r)
		return
	}
	h.notAllowed(w, r)
}

func (h *handler) notAllowed(w http.ResponseWriter, r *http.Request) {
	allowed := make([]string, len(h.methods))
	i := 0
	for k := range h.methods {
		allowed[i] = k
		i++
	}
	//List all the allowed(keys) methods and return on the header
	w.Header()["Allow"] = allowed
	m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
	http.Error(w, m, http.StatusMethodNotAllowed)
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//If the path matchs anything, the server runs the handler
	if m.handlers[r.URL.Path] != nil {
		m.handlers[r.URL.Path].handlerMethod(w, r)
		return
	}
	//Get a map key replacing all the values with a default id name
	key := m.getValues.ReplaceAllString(r.URL.Path, "{id}")
	if m.handlers[key] != nil {
		//Get all the values from path
		v := m.getValues.FindAllString(r.URL.Path, int(m.idsNum))
		m.handlers[key].handlerValues(w, r, v)
		return
	}
	//If hits here, then a not found error is delivered
	m.notFound(w, r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	NotFound(w)
}
