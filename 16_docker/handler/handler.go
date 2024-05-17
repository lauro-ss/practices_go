package handler

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Docker é bacana \U0001F433 \n"))
}
