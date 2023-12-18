package util

import (
	"encoding/json"
	"net/http"
)

func AsJson(w http.ResponseWriter, o any) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
	return nil
}
