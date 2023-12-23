package main

import (
	"httpserver/handler"
	"httpserver/helper"
	"httpserver/helper/middleware"
	"net/http"
)

func main() {
	a := helper.NewApi()

	a.Use(middleware.ContentIsJson)
	a.Use(middleware.Logger)

	a.Get("/animals", handler.AnimalList)
	a.Post("/animals", handler.AnimalPost)
	a.Get("/animals/{animalId}", handler.AnimalGet)

	err := http.ListenAndServe(
		":4500", a,
	)
	if err != nil {
		panic(err)
	}
}
