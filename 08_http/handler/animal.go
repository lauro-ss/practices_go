package handler

import (
	"httpserver/helper"
	"httpserver/model"
	"log"
	"net/http"
	"strconv"
)

func GetAnimal(w http.ResponseWriter, r *http.Request) {

	if !helper.CheckMethod(w, r, http.MethodGet) {
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	a := model.Animal{Id: id, Name: "Cat", Icon: "\U0001F431"}

	j, err := helper.AsJson(a)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	log.Println("Get Animal")
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
