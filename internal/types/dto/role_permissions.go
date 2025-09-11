package dto

type RolePermissions struct {
	RoleID      int   `json:"role_id" validate:"required"`
	Permissions []int `json:"permissions" validate:"required,dive,gt=0"`
}