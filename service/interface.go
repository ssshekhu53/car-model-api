package service

import "github.com/ssshekhu53/car-model-api/models"

type Car interface { //interface has methods only.
	Get() ([]models.Car, error)
	Post(car models.Car) (models.Car, error)
	Put(id string, car models.Car) (models.Car, error)
	Delete(id string) error
}
