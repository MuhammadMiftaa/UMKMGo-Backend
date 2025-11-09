package dto

type RolePermissions struct {
	RoleID      int      `json:"role_id" validate:"required"`
	Permissions []string `json:"permissions" validate:"required,dive,gt=0"`
}
