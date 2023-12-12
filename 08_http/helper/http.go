package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CheckMethod(w http.ResponseWriter, r *http.Request, m string) bool {
	if r.Method != m {
		w.Header()["Allow"] = []string{m}
		m = fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
		http.Error(w, m, http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func InternalError(w http.ResponseWriter, err error) {
	log.Fatalln(err)
	m := fmt.Sprintf("%v internal server error", http.StatusInternalServerError)
	http.Error(w, m, http.StatusInternalServerError)
}

func AsJson(o any, w http.ResponseWriter) {
	b, err := json.Marshal(o)
	if err != nil {
		InternalError(w, err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(b))
	//w.Write(b)
}
