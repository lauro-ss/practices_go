package main

import "fmt"

type Food string

type Animal interface {
	Feed(Food) string
	String() string
}

type Dog string

type Cat string

func Walk(a Animal) string {
	return fmt.Sprintf("The %v is walking", a)
}

func (d Dog) Feed(f Food) string {
	return fmt.Sprintf("The %v eats %v", d, f)
}

func (d Dog) String() string {
	return string(d)
}

func (c Cat) Feed(f Food) string {
	return fmt.Sprintf("The %v eats %v", c, f)
}

func (c Cat) String() string {
	return string(c)
}
