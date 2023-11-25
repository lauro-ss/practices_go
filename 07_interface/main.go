package main

import "fmt"

type Dog string

type Cat string

func main() {
	dog := Dog("\U0001F436")
	cat := Cat("\U0001F431")

	fmt.Println(Walk(dog))
	fmt.Println(Walk(cat))
	fmt.Println(dog.Feed("\U0001F356"))
	fmt.Println(cat.Feed("\U0001F41F"))
}

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
