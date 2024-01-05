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
	rows, err := a.db.Query(context.Background(), "select id, name, emoji from animal;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int
	err = a.db.QueryRow(context.Background(), "select count(1) from animal;").Scan(&count)
	if err != nil {
		return nil, err
	}
	animals := make([]data.Animal, count)
	count = 0
	for rows.Next() {
		err := rows.Scan(&animals[count].Id, &animals[count].Name, &animals[count].Emoji)
		if err != nil {
			animals = nil
			return nil, err
		}
		count++
	}

	if rows.Err() != nil {
		return nil, err
	}

	return animals, nil
}
