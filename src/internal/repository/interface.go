package repository

import (
	"cars_catalog/internal/entity/cars"
	"cars_catalog/internal/entity/people"

	"github.com/jmoiron/sqlx"
)

type Car interface {
	GetCarsList(tx *sqlx.Tx) ([]cars.Car, error)
	DeleteCarByID(tx *sqlx.Tx, carID int) error
	UpdateCar(tx *sqlx.Tx, c cars.UpdateCarParam) error
	AddCar(tx *sqlx.Tx, a cars.AddCarParam) error
}

type People interface {
	GetPeopleList(tx *sqlx.Tx) ([]people.People, error)
	DeletePeople(tx *sqlx.Tx, peopleID int) error
	UpdatePeople(tx *sqlx.Tx, p people.UpdatePeopleParam) error
}

type HttpCLient interface {
	GetCarInfo(regNum string) (cars.AddCarParam, error)
}
