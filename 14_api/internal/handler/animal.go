package handler

import (
	"net/http"

	"github.com/lauro-ss/practices_go/14_api/internal/service"
	"github.com/lauro-ss/practices_go/14_api/pkg/api"
)

func GetAnimal(ar service.AnimalRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := ar.Get(r.Form.Get("animalId"))

		if err != nil {
			api.InternalError(w)
			return
		}

		api.AsJson(w, a)
	}
}
