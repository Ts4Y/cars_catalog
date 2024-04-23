package postgresql

import (
	"cars_catalog/internal/entity/cars"
	"cars_catalog/internal/repository"

	"github.com/jmoiron/sqlx"
)

type CarsRepository struct {
}

func NewCarsRepo() repository.Car {
	return &CarsRepository{}
}

func (c *CarsRepository) GetCarsList(tx *sqlx.Tx) ([]cars.Car, error) {
	var cars []cars.Car

	sqlQuery := `
	select reg_num,mark,model,year,people_id from actors;`

	err := tx.Select(&cars, sqlQuery)
	return cars, err
}

func (c *CarsRepository) DeleteCarByID(tx *sqlx.Tx, carid int) error {

	sqlQuery := `
	delete from cars
	where id = $1`

	_, err := tx.Exec(sqlQuery, carid)
	return err
}

func (c *CarsRepository) UpdateCar(tx *sqlx.Tx, car cars.UpdateCarParam) error {

	sqlQuery := `
	update cars set
	reg_num = coalesce(:reg_num,reg_num),
	mark = coalesce(:mark,mark),
	model = coalesce(:model,model)
	year = coalesce(:year,year)
	people_id = coalesce(:people_id,people_id)
	where id = :id`

	_, err := tx.NamedExec(sqlQuery, car)
	return err
}

func (c *CarsRepository) AddCar(tx *sqlx.Tx, a cars.AddCarParam) error {
	sqlQuery := `
	insert into cars (reg_num, mark, model, year, people_id),
	values(:reg_num,:mark,:model,:year,:people_id)`

	_, err := tx.NamedExec(sqlQuery, a)
	return err
}
