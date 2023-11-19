package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
		return nil, errors.New("you need to init your Deck with initDeck function")
	}
	if size > len(*d) {
		return Deck{},
			fmt.Errorf("the Deck has a size of %v and don't support a hand size of %v", len(*d), size)
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

func shuffle(d Deck, runs int, fillDeck *Deck, r *rand.Rand) {
	if runs > 0 && len(d) > 0 {
		flag := r.Int31n(int32(len(d)))
		shuffle(d[flag:], runs-1, fillDeck, r)
		shuffle(d[:flag], runs-1, fillDeck, r)
	} else {
		for _, card := range d {
			*fillDeck = append(*fillDeck, Card{Num: card.Num, Name: card.Name})
		}
	}
}

func Shuffle(d Deck) Deck {
	deck := Deck{}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffle(d, 3, &deck, rand)
	return deck
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

func (d Deck) SaveToBinFile(fileName string) {
	file, er := os.Create(fileName)
	if er != nil {
		log.Fatalln(er)
		os.Exit(1)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	er = enc.Encode(d)
	if er != nil {
		log.Fatalln(er)
		os.Exit(1)
	}
}

// Reads from a binary fileName and returns a Deck
func ReadFromBinFile(fileName string) (Deck, error) {
	var d *Deck = &Deck{}
	file, er := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if er != nil {
		log.Fatalln(er)
		os.Exit(1)
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	dec.Decode(d)

	return *d, nil
}
