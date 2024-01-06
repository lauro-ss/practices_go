package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lauro-ss/practices_go/14_api/internal/data"
)

type AnimalRepository interface {
	List() ([]data.Animal, error)
	Get(id string) (*data.Animal, error)
	// Update(data.Animal) (bool, error)
	// Delete(data.Animal) (bool, error)
	// Create(data.Animal) (string, error)
}

type AnimalSQL struct {
	db *pgxpool.Pool
}

func NewAnimalRepository(db *pgxpool.Pool) *AnimalSQL {
	return &AnimalSQL{db}
}

func (a *AnimalSQL) Get(id string) (animal *data.Animal, err error) {
	animal = &data.Animal{}
	err = a.db.QueryRow(context.Background(), "select id, name, emoji from animal where id = $1", id).
		Scan(&animal.Id, &animal.Name, &animal.Emoji)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return animal, nil
}

func (a *AnimalSQL) List() ([]data.Animal, error) {
	rows, err := a.db.Query(context.Background(),
		`select id, name, emoji from animal;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 10

	animals := make([]data.Animal, 0, count)
	count = 0
	animal := data.Animal{}
	for rows.Next() {
		err = rows.Scan(&animal.Id, &animal.Name, &animal.Emoji)
		if err != nil {
			animals = nil
			return nil, err
		}
		// animal.Foods, err = a.GetFoodByAnimalId(animal.Id)
		// if err != nil {
		// 	animals = nil
		// 	return nil, err
		// }
		if count == cap(animals) {
			animals = append(animals, make([]data.Animal, count)...)
		}
		animals = append(animals, animal)
		count++
	}

	if rows.Err() != nil {
		return nil, err
	}

	return animals, nil
}

func (a *AnimalSQL) GetFoodByAnimalId(idAnimal string) ([]data.Food, error) {
	rows, err := a.db.Query(context.Background(),
		`select f.id, f.name, f.emoji from animal a
		inner join animalfood af
		on (a.id = af.idanimal)
		inner join food f
		on (f.id  = af.idfood)
		where a.id = $1;`, idAnimal)

	if err != nil {
		return nil, err
	}

	count := 10

	foods := make([]data.Food, 0, count)
	count = 0
	food := data.Food{}

	for rows.Next() {
		err = rows.Scan(&food.Id, &food.Name, &food.Emoji)
		if err != nil {
			foods = nil
			return nil, err
		}
		if count == cap(foods) {
			foods = append(foods, make([]data.Food, count)...)
		}
		foods = append(foods, food)
		count++
	}

	return foods, nil
}
