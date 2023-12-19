package main

import (
	"fmt"
	"httpserver/handler"
	"httpserver/helper"
	"net/http"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/animals", handler.Home)
	// mux.HandleFunc("/animals", handler.AnimalGetPost)
	// mux.HandleFunc("/animals/", handler.Animal)
	r := helper.NewApi()
	r.Get("/animals", handler.AnimalList)
	r.Get("/animals/{animalId}", handler.AnimalGet)
	r.Get("/animals/{animalId}/foods/{foodId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("/animals/%v/foods/%v", r.Form.Get("animalId"), r.Form.Get("foodId"))))
	})

	err := http.ListenAndServe(
		":4500", r,
	)
	if err != nil {
		panic(err)
	}
}
