package main

import "fmt"

func main() {
	myDeck := initDeck()
	// er := myDeck.SaveToFile("my_deck.txt")
	// if er != nil {
	// 	log.Fatalln(er)
	// }
	// i, er := myDeck.SaveToBinFile("my_deck.bin")
	// fmt.Println(i)
	// if er != nil {
	// 	log.Fatalln(er)
	// }
	// myDeck.SaveToBinFile("my_deck.bin")
	// LoadFromBinFile("my_deck.bin")
	fmt.Println(Shuffle(myDeck))
}
