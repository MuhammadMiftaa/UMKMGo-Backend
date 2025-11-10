package dto

type Applications struct {
	ID          int                    `json:"id,omitempty"`
	UMKMID      int                    `json:"umkm_id" validate:"required"`
	ProgramID   int                    `json:"program_id" validate:"required"`
	Type        string                 `json:"type" validate:"required,oneof=training certification funding"`
	Status      string                 `json:"status,omitempty"`
	SubmittedAt string                 `json:"submitted_at,omitempty"`
	ExpiredAt   string                 `json:"expired_at,omitempty"`
	CreatedAt   string                 `json:"created_at,omitempty"`
	UpdatedAt   string                 `json:"updated_at,omitempty"`
	Documents   []ApplicationDocuments `json:"documents,omitempty"`
	Histories   []ApplicationHistories `json:"histories,omitempty"`
	Program     *Programs              `json:"program,omitempty"`
	UMKM        *UMKM                  `json:"umkm,omitempty"`
}

type ApplicationDocuments struct {
	ID            int    `json:"id,omitempty"`
	ApplicationID int    `json:"application_id,omitempty"`
	Type          string `json:"type" validate:"required,oneof=ktp nib npwp proposal portfolio rekening other"`
	File          string `json:"file" validate:"required"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

type ApplicationHistories struct {
	ID             int    `json:"id,omitempty"`
	ApplicationID  int    `json:"application_id,omitempty"`
	Status         string `json:"status" validate:"required"`
	Notes          string `json:"notes,omitempty"`
	ActionedAt     string `json:"actioned_at,omitempty"`
	ActionedBy     int    `json:"actioned_by,omitempty"`
	ActionedByName string `json:"actioned_by_name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

type ApplicationDecision struct {
	ApplicationID int    `json:"application_id" validate:"required"`
	Action        string `json:"action" validate:"required,oneof=approve reject revise"`
	Notes         string `json:"notes,omitempty"`
}

type UMKM struct {
	ID           int      `json:"id,omitempty"`
	UserID       int      `json:"user_id,omitempty"`
	BusinessName string   `json:"business_name,omitempty"`
	NIK          string   `json:"nik,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	BirthDate    string   `json:"birth_date,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	Address      string   `json:"address,omitempty"`
	ProvinceID   int      `json:"province_id,omitempty"`
	CityID       int      `json:"city_id,omitempty"`
	District     string   `json:"district,omitempty"`
	Subdistrict  string   `json:"subdistrict,omitempty"`
	PostalCode   string   `json:"postal_code,omitempty"`
	NIB          string   `json:"nib,omitempty"`
	NPWP         string   `json:"npwp,omitempty"`
	KartuType    string   `json:"kartu_type,omitempty"`
	KartuNumber  string   `json:"kartu_number,omitempty"`
	User         User     `json:"user"`
	Province     Province `json:"province"`
	City         City     `json:"city"`
}

type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

type City struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ProvinceID int    `json:"province_id,omitempty"`
}

type Province struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
