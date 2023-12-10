package helper

import (
	"encoding/json"
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

func AsJson(o any) ([]byte, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return b, nil
}
