package dto

type Permissions struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
