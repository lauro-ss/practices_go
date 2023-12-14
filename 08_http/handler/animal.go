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
	err = helper.AsJson(o, w)
	if err != nil {
		helper.InternalError(w, err)
		return
	}
}

func Animal(w http.ResponseWriter, r *http.Request) {

	id, err := helper.GetId(r.URL.Path)

	if err != nil {
		helper.InternalError(w, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Println(id)
		helper.AsJson("", w)
		return
	case http.MethodPost:
		return
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	default:
		m := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		}
		helper.NotAllowed(w, m)
	}
}
