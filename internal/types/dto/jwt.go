package dto

type UserData struct {
	ID           float64 `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Role         float64 `json:"role"`
	RoleName     string  `json:"role_name"`
	BusinessName string  `json:"business_name"`
	KartuType    string  `json:"kartu_type"`
	Phone        string  `json:"phone"`
}
