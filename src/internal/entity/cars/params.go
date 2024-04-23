package cars

import "database/sql"

type AddCarParam struct {
	RegNum   *sql.NullString `json:"registration_number" db:"reg_num"`
	Mark     *sql.NullString `json:"mark" db:"mark"`
	Model    *sql.NullString `json:"model" db:"model"`
	Year     *sql.NullInt64  `json:"year" db:"year"`
	PeopleID *sql.NullInt64  `json:"people_id" db:"people_id"`
}

func (a AddCarParam) Valid() bool {
	return a.RegNum != nil
}

type UpdateCarParam struct {
	CarID    int            `json:"id" db:"id"`
	RegNum   sql.NullString `json:"registration_number" db:"reg_num"`
	Mark     sql.NullString `json:"mark" db:"mark"`
	Model    sql.NullString `json:"model" db:"model"`
	Year     sql.NullInt64  `json:"year" db:"year"`
	PeopleID sql.NullInt64  `json:"people_id" db:"people_id"`
}
