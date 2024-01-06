package handler

import (
	"log"
	"net/http"

	"github.com/lauro-ss/practices_go/14_api/internal/service"
	"github.com/lauro-ss/practices_go/14_api/pkg/api"
)

func ListAnimal(ar service.AnimalRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := ar.List()

		if err != nil {
			log.Println(err)
			api.InternalError(w)
			return
		}
		for i := range a {
			a[i].Hateoas.SelfGet(r)
		}
		api.AsJson(w, a)
	}
}

func GetAnimal(ar service.AnimalRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := ar.Get(r.Form.Get("animalId"))

		if err != nil {
			api.InternalError(w)
			return
		}

		if a == nil {
			api.NotFound(w)
		}
		a.Hateoas.SelfGet(r)
		api.AsJson(w, a)
	}
}
