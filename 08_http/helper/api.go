package helper

import (
	"net/http"
	"regexp"
)

type Api struct {
	Handlers  map[string]http.HandlerFunc
	Values    map[string][]string
	getValues *regexp.Regexp //Regex for get the values from Path
	getIds    *regexp.Regexp //Regex for get the Ids names from Path
}

func NewApi() *Api {
	return &Api{
		Handlers:  make(map[string]http.HandlerFunc, 5),
		Values:    make(map[string][]string, 5),
		getValues: regexp.MustCompile("[0-9]+"),
		getIds:    regexp.MustCompile("{[A-Z-a-z]+}"),
	}
}

// Creates a unique uri
// Ex: "/clients"
// This uri only accepts GET/POST requests
func (a *Api) AddUri(pattern string, handler http.HandlerFunc) {
	a.Handlers[pattern] = handler
}

// Creates a unique uri with a custom id name
// Ex: "/clients/{clientId}"
// Ex: "/clients/{clienteId}/orders"
// Ex: "/clients/{clienteId}/orders/{orderId}"
// This uri only accepts DELETE/GET/PATCH/PUT requests
func (a *Api) AddUriId(pattern string, handler http.HandlerFunc) {
	ids := a.getIds.FindAllString(pattern, 2)
	//Replace all the custom ids for a default id name
	pattern = a.getIds.ReplaceAllString(pattern, "{id}")
	a.Handlers[pattern] = handler
	a.Values[pattern] = ids
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.Handlers[r.URL.Path] != nil {
		a.Handlers[r.URL.Path].ServeHTTP(w, r)
		return
	}
	//Get a map key replacing all the values with a default id name
	key := a.getValues.ReplaceAllString(r.URL.Path, "{id}")
	if a.Handlers[key] != nil {
		v := a.getValues.FindAllString(r.URL.Path, 2)
		r.ParseForm()
		r.Form.Add(a.Values[key][0][1:len(a.Values[key][0])-1], v[0])
		r.Form.Add(a.Values[key][1][1:len(a.Values[key][1])-1], v[1])

		a.Handlers[key].ServeHTTP(w, r)
		return
	}

	http.Error(w, "not found", http.StatusNotFound)
}
