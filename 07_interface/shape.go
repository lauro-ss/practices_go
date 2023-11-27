package main

import "fmt"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Square struct {
	Side float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}
