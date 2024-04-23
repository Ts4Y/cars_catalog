package postgresql

import (
	"cars_catalog/internal/entity/people"
	"cars_catalog/internal/repository"

	"github.com/jmoiron/sqlx"
)

type peopleRepo struct{}

func NewPeopleRepo() repository.People {
	return &peopleRepo{}
}

func (r *peopleRepo) GetPeopleList(tx *sqlx.Tx) ([]people.People, error) {
	var peopleList []people.People

	sqlQuery := `
	select id, name, surname, patronymic
	from people`

	err := tx.Select(&peopleList, sqlQuery)
	if err != nil {
		return nil, err
	}

	return peopleList, nil
}

func (r *peopleRepo) DeletePeople(tx *sqlx.Tx, peopleID int) error {
	sqlQuery := `
	delete from people
	where id = $1`

	_, err := tx.Exec(sqlQuery, peopleID)
	if err != nil {
		return err
	}

	return nil
}

func (r *peopleRepo) UpdatePeople(tx *sqlx.Tx, p people.UpdatePeopleParam) error {
	sqlQuery := `
	update people set
	name = coalesce(:name, name)
	surname = coalesce(:surname, surname)
	patronymic = coalesce(:patronymic, patronymic)
	where id = :people_id`

	_, err := tx.NamedExec(sqlQuery, p)
	if err != nil {
		return err
	}

	return nil
}
