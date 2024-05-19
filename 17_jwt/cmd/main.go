package main

import (
	"fmt"
	"jwt-golang/config"
	"jwt-golang/handlers"
	"jwt-golang/services"
	"log"
	"net/http"
)

func main() {
	db := config.ConnectDatabase()
	err := db.AutoMigrate(services.User{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	userRepository := services.NewUserRepository(db)

	mux := http.NewServeMux()
	mux.Handle("POST /users", handlers.NewUser(userRepository))
	mux.Handle("POST /login", handlers.UserLogin(userRepository))

	fmt.Println("Server runing on localhost 8080")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf(err.Error())
	}
}
