package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(os.Stdout, "\r%v", i)
	}
}
