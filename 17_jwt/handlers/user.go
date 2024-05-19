package handlers

import (
	"encoding/json"
	"jwt-golang/services"
	"net/http"
)

func NewUser(ur services.UserRepository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var user services.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
		}

		id, err := ur.New(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
		}

		w.Write([]byte(id))
	}
	return http.HandlerFunc(fn)
}

func UserLogin(ur services.UserRepository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
		}

		tokens, err := ur.Authentication(request.Login, request.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error 400"))
		}

		jsonBytes, err := json.Marshal(tokens)
		if err != nil {
			w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
		}
		w.Write(jsonBytes)
	}
	return http.HandlerFunc(fn)
}
