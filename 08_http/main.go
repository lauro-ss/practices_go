package main

import (
	"fmt"
	"net/http"
	"os"
)

type Test interface{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	if re, ok := resp.Body.(Test); ok {
		fmt.Print(re)
	}
	// fmt.Println(string(bs))
}
