package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lauro-ss/practices_go/14_api/internal/data"
)

type AnimalRepository interface {
	// List() ([]data.Animal, error)
	Get(id string) (data.Animal, error)
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

func (a *AnimalSQL) Get(id string) (*data.Animal, error) {
	animal := data.Animal{}
	err := a.db.QueryRow(context.Background(), "select id, name, emoji from animal where id = $1", id).
		Scan(&animal.Id, &animal.Name, &animal.Emoji)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}
