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
	UMKM        *UMKMWeb               `json:"umkm,omitempty"`
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
	ActionedBy     *int   `json:"actioned_by,omitempty"`
	ActionedByName string `json:"actioned_by_name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

type ApplicationDecision struct {
	ApplicationID int    `json:"application_id" validate:"required"`
	Action        string `json:"action" validate:"required,oneof=approve reject revise"`
	Notes         string `json:"notes,omitempty"`
}
