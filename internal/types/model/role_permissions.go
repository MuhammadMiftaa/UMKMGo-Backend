package model

import "encoding/json"

type RolePermissions struct {
	RoleID       int `json:"role_id" gorm:"primary_key"`
	PermissionID int `json:"permission_id" gorm:"primary_key"`

	Base
}

type RolePermissionsResponse struct {
	RoleID      int             `json:"role_id"`
	RoleName    string          `json:"role_name"`
	Permissions json.RawMessage `json:"permissions"`
}
