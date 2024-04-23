package cars

type Car struct {
	ID      int    `json:"id" db:"id"`
	RegNum  string `json:"registration_number" db:"reg_num"`
	Mark    string `json:"mark" db:"mark"`
	Model   string `json:"model" db:"model"`
	Year    int    `json:"year" db:"year"`
	PeopleID int    `json:"people_id" db:"people_id"`
}
