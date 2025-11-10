package model

type Role struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Base
}
