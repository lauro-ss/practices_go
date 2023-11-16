package main

import (
	"fmt"
	"log"
)

func main() {
	//myDeck := initDeck()
	var myDeck deck = nil
	fmt.Print(myDeck)
	fmt.Println("Dealing hand...")
	hand, er := myDeck.Deal(0)

	if er != nil {
		log.Fatal(er)
	}
	fmt.Println("Your hand...")
	fmt.Print(hand)
	fmt.Println("Your deck...")
	fmt.Print(myDeck)
}
