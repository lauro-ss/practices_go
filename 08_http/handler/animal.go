package handler

import (
	"httpserver/helper"
	"httpserver/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const PATH = "./data/"

func ListAnimal(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckMethod(w, r, http.MethodGet) {
		return
	}
	b, err := os.ReadFile(PATH + "animal.csv")
	if err != nil {
		helper.InternalError(w, err)
		return
	}

	//To add database
	s := strings.Split(string(b), "\n")[1:]

	o := make([]model.Animal, len(s))
	for i, v := range s {
		c := strings.Split(v, ";")
		o[i].Id, err = strconv.Atoi(c[0])
		if err != nil {
			helper.InternalError(w, err)
			return
		}
		o[i].Name = c[1]
		o[i].Icon = c[2]
	}

	helper.AsJson(o, w)
}

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

	helper.AsJson(a, w)
}
