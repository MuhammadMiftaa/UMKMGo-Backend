package dto

type Users struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
	RoleID   int    `json:"role_id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}

type OTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
