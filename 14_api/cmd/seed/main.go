package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lauro-ss/practices_go/14_api/internal/data"
)

func main() {
	conn, err := data.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// SeedAnimal(conn)
	// SeedFood(conn)
	SeedAnimalFood(conn)
}

func SeedAnimal(c *pgxpool.Pool) {
	a := []data.Animal{
		{
			Name:  "Cat",
			Emoji: "🐱",
		},
		{
			Name:  "Dog",
			Emoji: "🐶",
		},
		{
			Name:  "Cow",
			Emoji: "🐮",
		},
		{
			Name:  "Tiger",
			Emoji: "🐯",
		},
	}

	for _, v := range a {
		_, err := c.Exec(context.Background(), "insert into animal (name, emoji) values ($1, $2);", v.Name, v.Emoji)
		if err != nil {
			panic(err)
		}
	}
}

func SeedFood(c *pgxpool.Pool) {
	a := []data.Food{
		{
			Name:  "Meat",
			Emoji: "🥩",
		},
		{
			Name:  "Bacon",
			Emoji: "🥓",
		},
		{
			Name:  "Egg",
			Emoji: "🥚",
		},
		{
			Name:  "Grass",
			Emoji: "🌿",
		},
	}

	for _, v := range a {
		_, err := c.Exec(context.Background(), "insert into food (name, emoji) values ($1, $2);", v.Name, v.Emoji)
		if err != nil {
			panic(err)
		}
	}
}

func SeedAnimalFood(c *pgxpool.Pool) {
	rowsAnimal, err := c.Query(context.Background(), "select id from animal;")
	if err != nil {
		panic(err)
	}
	defer rowsAnimal.Close()

	rowsFood, err := c.Query(context.Background(), "select id from food;")
	if err != nil {
		panic(err)
	}
	defer rowsFood.Close()

	var count int
	err = c.QueryRow(context.Background(), "select count(1) from food;").Scan(&count)
	if err != nil {
		panic(err)
	}
	foods := make([]string, count)
	count--
	for rowsFood.Next() {
		rowsFood.Scan(&foods[count])
		count--
	}

	var (
		idAnimal string
	)
	for rowsAnimal.Next() {
		err = rowsAnimal.Scan(&idAnimal)
		if err != nil {
			panic(err)
		}

		for _, v := range foods {
			_, err = c.Exec(context.Background(), "insert into animalfood (idanimal, idfood) values ($1, $2);", idAnimal, v)
			if err != nil {
				panic(err)
			}
		}

	}

	if rowsAnimal.Err() != nil {
		panic(rowsAnimal.Err())
	}

	if rowsFood.Err() != nil {
		panic(rowsFood.Err())
	}
}
