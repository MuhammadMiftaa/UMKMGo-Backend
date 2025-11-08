package dto

type Programs struct {
	ID                  int      `json:"id,omitempty"`
	Title               string   `json:"title" validate:"required"`
	Description         string   `json:"description,omitempty"`
	Banner              string   `json:"banner,omitempty"`
	Provider            string   `json:"provider,omitempty"`
	ProviderLogo        string   `json:"provider_logo,omitempty"`
	Type                string   `json:"type" validate:"required,oneof=training certification funding"`
	TrainingType        *string  `json:"training_type,omitempty" validate:"omitempty,oneof=online offline hybrid"`
	Batch               *int     `json:"batch,omitempty"`
	BatchStartDate      *string  `json:"batch_start_date,omitempty"`
	BatchEndDate        *string  `json:"batch_end_date,omitempty"`
	Location            *string  `json:"location,omitempty"`
	MinAmount           *float64 `json:"min_amount,omitempty"`
	MaxAmount           *float64 `json:"max_amount,omitempty"`
	InterestRate        *float64 `json:"interest_rate,omitempty"`
	MaxTenureMonths     *int     `json:"max_tenure_months,omitempty"`
	ApplicationDeadline string   `json:"application_deadline" validate:"required"`
	IsActive            bool     `json:"is_active"`
	CreatedBy           int      `json:"created_by,omitempty"`
	CreatedByName       string   `json:"created_by_name,omitempty"`
	CreatedAt           string   `json:"created_at,omitempty"`
	UpdatedAt           string   `json:"updated_at,omitempty"`
	Benefits            []string `json:"benefits,omitempty"`
	Requirements        []string `json:"requirements,omitempty"`
}

type ProgramBenefits struct {
	ID        int    `json:"id"`
	ProgramID int    `json:"program_id"`
	Name      string `json:"name" validate:"required"`
}

type ProgramRequirements struct {
	ID        int    `json:"id"`
	ProgramID int    `json:"program_id"`
	Name      string `json:"name" validate:"required"`
}
