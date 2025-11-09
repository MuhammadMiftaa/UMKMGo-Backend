package dto

type Users struct {
	ID              int      `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Email           string   `json:"email,omitempty" validate:"required,email"`
	Password        string   `json:"password,omitempty" validate:"required,min=8"`
	ConfirmPassword string   `json:"confirm_password,omitempty" validate:"required,min=8"`
	RoleID          *int     `json:"role_id,omitempty"`
	RoleName        string   `json:"role_name,omitempty"`
	LastLoginAt     string   `json:"last_login_at,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
	CreatedAt       string   `json:"created_at,omitempty"`
	UpdatedAt       string   `json:"updated_at,omitempty"`
	IsActive        bool     `json:"is_active,omitempty"`
}

type OTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
