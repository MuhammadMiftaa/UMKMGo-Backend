package dto

type SLA struct {
	ID          int    `json:"id,omitempty"`
	Status      string `json:"status" validate:"required,oneof=screening final"`
	MaxDays     int    `json:"max_days" validate:"required,min=1"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type ExportRequest struct {
	FileType        string `json:"file_type" validate:"required,oneof=pdf excel"`
	ApplicationType string `json:"application_type" validate:"required,oneof=all funding training certification"`
}
