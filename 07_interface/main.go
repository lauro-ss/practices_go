package main

import "fmt"

type Animal interface {
	Walk() string
}

func main() {
	value := "\U0001f600"
	fmt.Printf("%v \n", value)

	value = "\U0001F92A"
	fmt.Printf("%v \n", value)

	value = "\U0001F914"
	fmt.Printf("%v", value)
}
