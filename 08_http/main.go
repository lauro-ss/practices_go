package main

import (
	"httpserver/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/animal", handler.ListAnimal)

	err := http.ListenAndServe(
		":4500", mux,
	)
	if err != nil {
		panic(err)
	}
}
