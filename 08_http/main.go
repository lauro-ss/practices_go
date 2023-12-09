package main

import (
	"httpserver/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/animal", handler.GetAnimal)

	err := http.ListenAndServe(
		":4500", nil,
	)
	if err != nil {
		panic(err)
	}
}
