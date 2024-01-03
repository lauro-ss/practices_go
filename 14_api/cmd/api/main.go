package main

import (
	"github.com/lauro-ss/practices_go/14_api/internal/data"
)

func main() {
	conn, err := data.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

}
