package main

import (
	"fmt"
	"log"
	"my-go-app/handler"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.AutoMigrate(&handler.User{})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Index)
	mux.HandleFunc("GET /users", handler.ListUser(db))
	mux.HandleFunc("GET /users/create", handler.InsertUser(db))

	port := os.Getenv("SERVER_PORT")
	fmt.Println(port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalln(err)
	}
}
