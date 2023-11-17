package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
)

type Card struct {
	Num  int
	Name string
}

type Deck []Card

func initDeck() Deck {
	d := Deck{}
	shapes := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for i := 1; i <= 10; i++ {
		for _, shape := range shapes {
			d = append(d, Card{Num: i, Name: shape})
		}
	}
	return d
}

func (d *Deck) Deal(size int) (Deck, error) {
	if *d == nil {
		return nil, errors.New("You need to init your Deck with initDeck function")
	}
	if size > len(*d) {
		return Deck{},
			errors.New(fmt.Sprintf("The Deck has a size of %v and don't support a hand size of %v", len(*d), size))
	}
	dx := *d
	hand := dx[:size]
	*d = dx[size:]
	return hand, nil
}

func printCards(i int, d Deck) string {
	if i == len(d) {
		return ""
	} else {
		return fmt.Sprintf("%v of %v\n", d[i].Num, d[i].Name) + printCards(i+1, d)
	}
}

func (d Deck) String() string {
	return printCards(0, d)
}

func (d Deck) SaveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.String()), 0666)
}

func (d Deck) SaveToBinFile(fileName string) {
	file, _ := os.Create(fileName)
	defer file.Close()
	// if er != nil {
	// 	log.Fatalln(er)
	// }
	//var pi int8 = 2
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(file)
	er := enc.Encode(d)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(buf.Bytes())
	//return file.Write(bin.Bytes())
}
