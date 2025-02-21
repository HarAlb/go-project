package entities

type Employee struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	CompanyId int `json:"company_id"`
	Position string `json:"position"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}