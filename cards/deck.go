package main

import (
	"errors"
	"fmt"
)

type card struct {
	num  int
	name string
}

type deck []card

func initDeck() deck {
	d := deck{}
	shapes := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for i := 1; i <= 10; i++ {
		for _, shape := range shapes {
			d = append(d, card{num: i, name: shape})
		}
	}
	return d
}

func (d *deck) Deal(size int) (deck, error) {
	if *d == nil {
		return nil, errors.New("You need to init your deck with initDeck function")
	}
	if size > len(*d) {
		return deck{},
			errors.New(fmt.Sprintf("The deck has a size of %v and don't support a hand size of %v", len(*d), size))
	}
	dx := *d
	hand := dx[:size]
	*d = dx[size:]
	return hand, nil
}

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
