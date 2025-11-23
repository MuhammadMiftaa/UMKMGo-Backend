package model

import "time"

type ApplicationHistory struct {
	ID            int       `json:"id" gorm:"primary_key"`
	ApplicationID int       `json:"application_id" gorm:"not null"`
	Status        string    `json:"status" gorm:"type:application_history_action;not null"`
	Notes         string    `json:"notes" gorm:"type:text"`
	ActionedAt    time.Time `json:"actioned_at" gorm:"default:NOW()"`
	ActionedBy    *int      `json:"actioned_by" gorm:"not null"`
	Base

	User User `json:"user" gorm:"foreignKey:ActionedBy"`
}
