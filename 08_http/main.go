package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Test interface{}

type LogWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// if re, ok := resp.Body.(Test); ok {
	// 	fmt.Print(re)
	// }
	// fmt.Println(string(bs))
	lw := LogWriter{}
	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
