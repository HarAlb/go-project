package entities

type Company struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Employers []User `json:"employers"`
	OwenrID int `json:"owner_id"`
}