package main

import "fmt"

type dict map[any]any

func main() {
	var t dict
	t = map[any]any{}
	t[0] = "Zero"
	t[1] = "One"

	delete(t, 1)

	fmt.Println(t)
}
