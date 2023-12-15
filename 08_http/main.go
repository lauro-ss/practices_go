package main

import (
	"fmt"
	"httpserver/handler"
	"net/http"
	"regexp"
)

type Router struct {
	Handlers map[string]http.HandlerFunc
	Values   map[string][]string
}

func NewReader() *Router {
	return &Router{
		Handlers: make(map[string]http.HandlerFunc, 5),
		Values:   make(map[string][]string, 5)}
}

func (re *Router) Add(pattern string, handler http.HandlerFunc) {
	rex := regexp.MustCompile("{[A-Z-a-z]+}")
	v := rex.FindAllString(pattern, 2)
	pattern = rex.ReplaceAllString(pattern, "{id}")
	re.Handlers[pattern] = handler
	re.Values[pattern] = v //Remover { } chaves daqui
}

func (re *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if re.Handlers[r.URL.Path] != nil {
		re.Handlers[r.URL.Path].ServeHTTP(w, r)
		return
	}
	rex := regexp.MustCompile("[0-9]+")
	mapString := rex.ReplaceAllString(r.URL.Path, "{id}")
	if re.Handlers[mapString] != nil {
		v := rex.FindAllString(r.URL.Path, 2)
		r.ParseForm()
		r.Form.Add(re.Values[mapString][0][1:len(re.Values[mapString][0])-1], v[0])
		r.Form.Add(re.Values[mapString][1][1:len(re.Values[mapString][1])-1], v[1])

		re.Handlers[mapString].ServeHTTP(w, r)
		return
	}

	http.Error(w, "not found", http.StatusNotFound)
	//r.URL.Query().
	//re.Patterns[r.URL.Path].

}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/animals", handler.Home)
	// mux.HandleFunc("/animals", handler.AnimalGetPost)
	// mux.HandleFunc("/animals/", handler.Animal)
	r := NewReader()
	r.Add("/animals", handler.Home)
	r.Add("/animals/{id}", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("/animals/{id}")) })
	r.Add("/animals/{animalId}/foods/{foodId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("/animals/%v/foods/%v", r.Form.Get("animalId"), r.Form.Get("foodId"))))
	})

	err := http.ListenAndServe(
		":4500", r,
	)
	if err != nil {
		panic(err)
	}
}
