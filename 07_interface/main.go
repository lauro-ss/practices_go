package main

func main() {
	// dog := Dog("\U0001F436")
	// cat := Cat("\U0001F431")

	// fmt.Println(Walk(dog))
	// fmt.Println(Walk(cat))
	// fmt.Println(dog.Feed("\U0001F356"))
	// fmt.Println(cat.Feed("\U0001F41F"))

	t := Triangle{Height: 10, Base: 10}
	s := Square{Side: 10}

	PrintArea(t)
	PrintArea(s)
}
