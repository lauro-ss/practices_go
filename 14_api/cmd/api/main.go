package main

import (
	"net/http"

	"github.com/lauro-ss/practices_go/14_api/internal/data"
	"github.com/lauro-ss/practices_go/14_api/internal/handler"
	"github.com/lauro-ss/practices_go/14_api/internal/service"
	"github.com/lauro-ss/practices_go/14_api/pkg/api"
	"github.com/lauro-ss/practices_go/14_api/pkg/api/middleware"
)

func main() {
	conn, err := data.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	api := api.NewApi()

	animalRepository := service.NewAnimalRepository(conn)

	api.Use(middleware.Logger)

	api.Get("/animals/{animalId}", handler.GetAnimal(animalRepository))

	http.ListenAndServe(":4500", api)
}
