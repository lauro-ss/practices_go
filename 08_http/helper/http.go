package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
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

func AsJson(o any, w http.ResponseWriter) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(b))
	//w.Write(b)
	return nil
}

func GetId(url string) (string, error) {
	rx := "[0-9]+"
	r, err := regexp.Compile(rx)
	if err != nil {
		return "", err
	}

	return r.FindString(url), nil
}

func NotAllowed(w http.ResponseWriter, allowed []string) {
	w.Header()["Allow"] = allowed
	m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
	http.Error(w, m, http.StatusMethodNotAllowed)
}
