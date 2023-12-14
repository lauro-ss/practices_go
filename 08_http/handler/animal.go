package handler

import (
	"fmt"
	"httpserver/helper"
	"httpserver/service"
	"net/http"
)

func ListAnimal(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckMethod(w, r, http.MethodGet) {
		return
	}
	o, err := service.GetAllAnimalCsv("animal.csv")
	if err != nil {
		helper.InternalError(w, err)
		return
	}
	helper.AsJson(o, w)
}

func Animal(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	}

	fmt.Println(r.URL.Path)

	helper.AsJson("", w)
}
