package main

import "fmt"

type enum map[string]string

func main() {
	var t enum
	t = map[string]string{}
	fmt.Println(t)
}
