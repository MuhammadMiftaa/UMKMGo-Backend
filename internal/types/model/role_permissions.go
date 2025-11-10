package model

type RolePermission struct {
	RoleID       int `json:"role_id" gorm:"primary_key"`
	PermissionID int `json:"permission_id" gorm:"primary_key"`

	Base
}
