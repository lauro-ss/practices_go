package helper

import (
	"fmt"
	"net/http"
)

func CheckMethod(w http.ResponseWriter, r *http.Request, m string) bool {
	if r.Method != http.MethodGet {
		w.Header()["Allow"] = []string{http.MethodGet}
		m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
		http.Error(w, m, http.StatusMethodNotAllowed)
		return false
	}

	return true
}
