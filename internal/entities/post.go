package entities

type Post struct {
	ID int `json:"id"`
	CompanyID int `json:"company_id"`
	EmployeeID int `json:"employee_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}