package main

import "fmt"

//import sys "fmt"

const PI = 3.14
const (
	B     = 1
	C int = 2
)

func main() {
	var numberInt int = 10
	numberFloat := 0.1

	var name, lastName string = "Lauro", "Santana"

	height, age := 1.69, 25

	/*
		var (
			a int8  = 1
			b int16 = 2
			c int32 = 3
			d int64 = 4
			e uint8 = 5
		)
	*/
	fmt.Println(numberInt)
	fmt.Printf("The number value is: %v and his type: %T \n", numberFloat, numberFloat)
	fmt.Printf("Name: %v Last Name: %v, Age: %v, Height: %v \n", name, lastName, age, height)
	fmt.Printf("%1.f", PI)
}
