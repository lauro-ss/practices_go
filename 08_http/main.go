package main

import (
	"fmt"
	"httpserver/handler"
	"net/http"
	"regexp"
)

type Reader struct {
	Handlers map[string]http.HandlerFunc
	Patterns map[string]*regexp.Regexp
}

func NewReader() *Reader {
	return &Reader{
		Handlers: make(map[string]http.HandlerFunc, 5),
		Patterns: make(map[string]*regexp.Regexp, 5)}
}

func (re *Reader) Add(pattern string, handler http.HandlerFunc) {
	//pattern = strings.ReplaceAll(pattern, "{id}", "[A-Za-z0-9]*")
	re.Patterns[pattern] = regexp.MustCompile(pattern)
	re.Handlers[pattern] = handler
}

func (re *Reader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if re.Handlers[r.URL.Path] != nil {
		re.Handlers[r.URL.Path].ServeHTTP(w, r)
		return
	}
	rex := regexp.MustCompile("[0-9]+")
	mapString := rex.ReplaceAllString(r.URL.Path, "{id}")
	if re.Handlers[mapString] != nil {
		re.Handlers[mapString].ServeHTTP(w, r)
		return
	}
	fmt.Println(mapString)
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
	r.Add("/animals/{id}/foods/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/animals/{id}/foods/{id}"))
	})

	err := http.ListenAndServe(
		":4500", r,
	)
	if err != nil {
		panic(err)
	}
}
