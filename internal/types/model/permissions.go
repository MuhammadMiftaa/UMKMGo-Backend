package model

type Permission struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ParentID    *int   `json:"parent_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`

	Base
}
