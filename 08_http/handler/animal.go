package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Animal")
	fmt.Println(r.Header)
	fmt.Println(r.Form.Get("id"))
	io.WriteString(w, "Geted Animal :3")
}
