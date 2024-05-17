package main

import (
	"log"
	"my-go-app/handler"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Index)

	port := ":8080"

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalln(err)
	}
}
