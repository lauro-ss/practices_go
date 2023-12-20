package handler

import (
	"fmt"
	"httpserver/helper"
	"httpserver/service"
	"io"
	"log"
	"net/http"
	"strconv"
)

func AnimalList(w http.ResponseWriter, r *http.Request) {
	o, err := service.GetAllAnimalCsv()
	if err != nil {
		helper.InternalError(w)
		return
	}
	helper.AsJson(w, o)
}

func AnimalGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("animalId"))
	if err != nil {
		log.Println("animal:", err)
		helper.InternalError(w)
		return
	}
	a, err := service.GetAnimalCsv(id)
	if err != nil {
		log.Println("animal:", err)
		helper.InternalError(w)
		return
	}
	if a == nil {
		helper.NotFound(w)
		return
	}
	helper.AsJson(w, a)
}

func AnimalPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Content-Type"))
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("animal:", err)
		helper.InternalError(w)
		return
	}
	defer r.Body.Close()
	fmt.Println(string(b))
}
