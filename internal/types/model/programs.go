package model

type Program struct {
	ID                  int      `json:"id" gorm:"primary_key"`
	Title               string   `json:"title" gorm:"type:varchar(100);not null"`
	Description         string   `json:"description" gorm:"type:text"`
	Banner              string   `json:"banner" gorm:"type:text"`
	Provider            string   `json:"provider" gorm:"type:varchar(100)"`
	ProviderLogo        string   `json:"provider_logo" gorm:"type:text"`
	Type                string   `json:"type" gorm:"type:program_type;not null"`
	TrainingType        *string  `json:"training_type" gorm:"type:training_type"`
	Batch               *int     `json:"batch"`
	BatchStartDate      *string  `json:"batch_start_date" gorm:"type:date"`
	BatchEndDate        *string  `json:"batch_end_date" gorm:"type:date"`
	Location            *string  `json:"location" gorm:"type:varchar(100)"`
	MinAmount           *float64 `json:"min_amount" gorm:"type:numeric(15,2)"`
	MaxAmount           *float64 `json:"max_amount" gorm:"type:numeric(15,2)"`
	InterestRate        *float64 `json:"interest_rate" gorm:"type:numeric(5,2)"`
	MaxTenureMonths     *int     `json:"max_tenure_months"`
	ApplicationDeadline string   `json:"application_deadline" gorm:"type:date"`
	IsActive            bool     `json:"is_active" gorm:"type:boolean;not null;default:true"`
	CreatedBy           int      `json:"created_by"`

	Base
	Users User `json:"users" gorm:"foreignKey:CreatedBy;references:ID"`
}
