package main

import "fmt"

type card struct {
	num  int
	name string
}

func initDeck() deck {
	d := deck{}
	for i := 1; i <= 10; i++ {
		d = append(d, card{num: i, name: "Spades"})
	}
	return d
}

type deck []card

func printCards(i int, d deck) string {
	if i == len(d) {
		return ""
	} else {
		return fmt.Sprintf("%v of %v\n", d[i].num, d[i].name) + printCards(i+1, d)
	}
}

func (d deck) String() string {
	return printCards(0, d)
}
