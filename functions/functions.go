package main

import "fmt"

func Show(value any) {
	fmt.Println(value)
}

func Hex(v int) string {
	return fmt.Sprintf("%x", v)
}

func Oct(v int) string {
	return fmt.Sprintf("%o", v)
}

func Bin(v int) string {
	return fmt.Sprintf("%b", v)
}

func Factorial(v int) int {
	if v > 1 {
		return v * Factorial(v-1)
	} else {
		return 1
	}
}

func Fibonacci(v int) int {
	if v > 2 {
		return Fibonacci(v-1) + Fibonacci(v-2)
	} else {
		return 1
	}
}

func main() {
	v := 10
	Show(Bin(v))
	Show(Oct(v))
	Show(Hex(v))
	Show(Factorial(5))
	Show(Fibonacci(10))
}
