package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lauro-ss/practices_go/14_api/internal/data"
)

type FoodRepository interface {
	List() ([]data.Food, error)
	Get(id string) (*data.Food, error)
	GetFoodByAnimalId(idAnimal string) ([]data.Food, error)
	// Update(data.Food) (bool, error)
	// Delete(data.Food) (bool, error)
	// Create(data.Food) (string, error)
}

type FoodSQL struct {
	db *pgxpool.Pool
}

func NewFoodRepository(db *pgxpool.Pool) *FoodSQL {
	return &FoodSQL{db}
}

func (f *FoodSQL) Get(id string) (Food *data.Food, err error) {
	Food = &data.Food{}
	err = f.db.QueryRow(context.Background(), "select id, name, emoji from Food where id = $1", id).
		Scan(&Food.Id, &Food.Name, &Food.Emoji)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return Food, nil
}
