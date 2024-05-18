package handler

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Docker é bacana \U0001F433 \n"))
}

func ListUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := User{}
		db.First(&u)
		w.Write([]byte(fmt.Sprint(u)))
	}
}

func InsertUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{Name: "Lauro"}
		fmt.Println("Criando User")
		db.Create(&user)
	}
}
