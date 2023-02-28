package service

import (
	"github.com/ssshekhu53/car-model-api/models"
	"github.com/ssshekhu53/car-model-api/store"
)

type service struct {
	store store.Car
}

func New(store store.Car) service {
	return service{store: store}
}

func (s service) Get() ([]models.Car, error) {
	return []models.Car{}, nil
}

func (s service) Post(car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s service) Put(id string, car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s service) Delete(id string) error {
	return nil
}
