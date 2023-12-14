package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	log.Println(err)
	m := fmt.Sprintf("%v internal server error", http.StatusInternalServerError)
	http.Error(w, m, http.StatusInternalServerError)
}

func AsJson(w http.ResponseWriter, o any) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(b))
	//w.Write(b)
	return nil
}

func GetIntId(url string) (int, error) {
	rx := "[0-9]+"
	r, err := regexp.Compile(rx)
	if err != nil {
		return 0, err
	}
	str := r.FindString(url)
	if str == "" {
		return 0, nil
	}
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NotAllowed(w http.ResponseWriter, allowed []string) {
	w.Header()["Allow"] = allowed
	m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
	http.Error(w, m, http.StatusMethodNotAllowed)
}

func NotFound(w http.ResponseWriter) {
	m := fmt.Sprintf("%v not found", http.StatusNotFound)
	http.Error(w, m, http.StatusNotFound)
}
