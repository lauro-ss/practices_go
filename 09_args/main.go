package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0755)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
