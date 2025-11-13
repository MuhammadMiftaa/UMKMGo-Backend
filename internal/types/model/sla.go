package model

type SLA struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Status      string `json:"status" gorm:"type:varchar(50);not null;unique"`
	MaxDays     int    `json:"max_days" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	Base
}
