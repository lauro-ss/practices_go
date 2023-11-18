package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func ReadFromFile(fileName string) (Deck, error) {
	fileBytes, er := os.ReadFile(fileName)
	deck := Deck{}
	if er != nil {
		return nil, er
	}
	stringDeck := string(fileBytes)
	stringCards := strings.Split(stringDeck, "\n")

	//the last indice it's a empty string
	for _, card := range stringCards[:len(stringCards)-1] {
		if string(card[1]) != " " {
			stringValue := string(card[0]) + string(card[1])
			num, _ := strconv.Atoi(stringValue)
			deck = append(deck, Card{Num: num, Name: string(card[6:])})
		} else {
			num, _ := strconv.Atoi(string(card[0]))
			deck = append(deck, Card{Num: num, Name: string(card[5:])})
		}

	}

	return deck, er
}

func (d Deck) SaveToBinFile(fileName string) error {
	file, er := os.Create(fileName)
	defer file.Close()
	if er != nil {
		log.Fatal(er)
	}

	enc := gob.NewEncoder(file)
	er = enc.Encode(d)
	if er != nil {
		log.Fatal(er)
	}

	return nil
}

// Reads from a binary fileName and returns a Deck
func ReadFromBinFile(fileName string) (Deck, error) {
	var d *Deck = &Deck{}
	file, er := os.OpenFile(fileName, os.O_RDONLY, 0666)
	defer file.Close()
	if er != nil {
		return nil, er
	}

	dec := gob.NewDecoder(file)
	dec.Decode(d)

	return *d, nil
}
