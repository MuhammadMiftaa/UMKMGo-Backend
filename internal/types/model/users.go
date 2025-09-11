package model

import "time"

type Users struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Email       string    `json:"email" gorm:"uniqueIndex"`
	Password    string    `json:"password"`
	RoleID      int       `json:"role_id"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	LastLoginAt time.Time `json:"last_login_at"`

	Base
	Roles Roles `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}
