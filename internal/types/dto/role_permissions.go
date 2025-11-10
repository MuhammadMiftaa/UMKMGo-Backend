package dto

import "encoding/json"

type RolePermissions struct {
	RoleID      int      `json:"role_id" validate:"required"`
	Permissions []string `json:"permissions" validate:"required,dive,gt=0"`
}

type RolePermissionsResponse struct {
	RoleID      int             `json:"role_id"`
	RoleName    string          `json:"role_name"`
	Permissions json.RawMessage `json:"permissions"`
}
