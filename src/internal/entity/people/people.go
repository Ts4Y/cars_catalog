package people

type People struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}
