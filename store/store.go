package store

import (
	"context"
	"database/sql"

	"github.com/ssshekhu53/car-model-api/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) Get() ([]models.Car, error) {
	_, err := s.db.QueryContext(context.Background(), "SELECT * FROM `cars`")
	return []models.Car{}, err
}

func (s store) Post(car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s store) Put(id string, car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s store) Delete(id string) error {
	return nil
}
