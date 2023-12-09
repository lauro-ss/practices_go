package handler

import (
	"fmt"
	"httpserver/helper"
	"io"
	"log"
	"net/http"
)

func GetAnimal(w http.ResponseWriter, r *http.Request) {

	if !helper.CheckMethod(w, r, http.MethodGet) {
		return
	}

	log.Println("Get Animal")
	fmt.Println(r.URL.Query().Get("id"))
	io.WriteString(w, "Geted Animal :3")
}
