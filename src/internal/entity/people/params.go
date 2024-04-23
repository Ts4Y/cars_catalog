package people

import "database/sql"

type UpdatePeopleParam struct {
	PeopleID   int            `json:"people_id" db:"people_id"`
	Name       sql.NullString `json:"name" db:"name"`
	Surname    sql.NullString `json:"surname" db:"surname"`
	Patronymic sql.NullString `json:"patronymic" db:"patronymic"`
}

func (p UpdatePeopleParam) Valid() bool {
	return p.Name.String != "" && p.Surname.String != "" && p.Patronymic.String != ""
}
