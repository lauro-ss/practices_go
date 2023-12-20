package main

import (
	"httpserver/handler"
	"httpserver/helper"
	"httpserver/helper/middleware"
	"net/http"
)

func main() {
	a := helper.NewApi()
	a.Get("/animals", handler.AnimalList)
	a.Post("/animals", handler.AnimalPost)
	a.Get("/animals/{animalId}", handler.AnimalGet)

	b := middleware.NewLogger(a)
	c := middleware.NewContent(b)

	err := http.ListenAndServe(
		":4500", c,
	)
	if err != nil {
		panic(err)
	}
}
