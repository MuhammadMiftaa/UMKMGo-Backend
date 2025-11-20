package dto

// Program List Response
type ProgramListMobile struct {
	ID                  int      `json:"id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	Banner              string   `json:"banner"`
	Provider            string   `json:"provider"`
	ProviderLogo        string   `json:"provider_logo"`
	Type                string   `json:"type"`
	TrainingType        *string  `json:"training_type,omitempty"`
	Batch               *int     `json:"batch,omitempty"`
	BatchStartDate      *string  `json:"batch_start_date,omitempty"`
	BatchEndDate        *string  `json:"batch_end_date,omitempty"`
	Location            *string  `json:"location,omitempty"`
	MinAmount           *float64 `json:"min_amount,omitempty"`
	MaxAmount           *float64 `json:"max_amount,omitempty"`
	InterestRate        *float64 `json:"interest_rate,omitempty"`
	MaxTenureMonths     *int     `json:"max_tenure_months,omitempty"`
	ApplicationDeadline string   `json:"application_deadline"`
	IsActive            bool     `json:"is_active"`
}

// Program Detail Response
type ProgramDetailMobile struct {
	ProgramListMobile
	Benefits     []string `json:"benefits"`
	Requirements []string `json:"requirements"`
}

// UMKM Profile Response
type UMKMProfile struct {
	ID             int      `json:"id"`
	UserID         int      `json:"user_id"`
	BusinessName   string   `json:"business_name"`
	NIK            string   `json:"nik"`
	Gender         string   `json:"gender"`
	BirthDate      string   `json:"birth_date"`
	Phone          string   `json:"phone"`
	Address        string   `json:"address"`
	ProvinceID     int      `json:"province_id"`
	CityID         int      `json:"city_id"`
	District       string   `json:"district"`
	Subdistrict    string   `json:"subdistrict"`
	PostalCode     string   `json:"postal_code"`
	NIB            string   `json:"nib,omitempty"`
	NPWP           string   `json:"npwp,omitempty"`
	RevenueRecord  string   `json:"revenue_record,omitempty"`
	BusinessPermit string   `json:"business_permit,omitempty"`
	KartuType      string   `json:"kartu_type"`
	KartuNumber    string   `json:"kartu_number"`
	Province       Province `json:"province"`
	City           City     `json:"city"`
	User           User     `json:"user"`
}

// Update UMKM Profile Request
type UpdateUMKMProfile struct {
	BusinessName string `json:"business_name" validate:"required"`
	Gender       string `json:"gender" validate:"required,oneof=male female other"`
	BirthDate    string `json:"birth_date" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Address      string `json:"address" validate:"required"`
	ProvinceID   int    `json:"province_id" validate:"required"`
	CityID       int    `json:"city_id" validate:"required"`
	District     string `json:"district" validate:"required"`
	Subdistrict  string `json:"subdistrict" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	KartuType    string `json:"kartu_type" validate:"required,oneof=produktif afirmatif"`
}

// Upload Document Request
type UploadDocumentRequest struct {
	Document string `json:"document" validate:"required"`
}

// Create Application Request - Training
type CreateApplicationTraining struct {
	ProgramID          int               `json:"program_id" validate:"required"`
	Motivation         string            `json:"motivation" validate:"required"`
	BusinessExperience string            `json:"business_experience"`
	LearningObjectives string            `json:"learning_objectives"`
	AvailabilityNotes  string            `json:"availability_notes"`
	Documents          map[string]string `json:"documents" validate:"required"`
}

// Create Application Request - Certification
type CreateApplicationCertification struct {
	ProgramID           int               `json:"program_id" validate:"required"`
	BusinessSector      string            `json:"business_sector" validate:"required"`
	ProductOrService    string            `json:"product_or_service" validate:"required"`
	BusinessDescription string            `json:"business_description" validate:"required"`
	YearsOperating      *int              `json:"years_operating"`
	CurrentStandards    string            `json:"current_standards"`
	CertificationGoals  string            `json:"certification_goals" validate:"required"`
	Documents           map[string]string `json:"documents" validate:"required"`
}

// Create Application Request - Funding
type CreateApplicationFunding struct {
	ProgramID             int               `json:"program_id" validate:"required"`
	BusinessSector        string            `json:"business_sector" validate:"required"`
	BusinessDescription   string            `json:"business_description" validate:"required"`
	YearsOperating        *int              `json:"years_operating"`
	RequestedAmount       float64           `json:"requested_amount" validate:"required"`
	FundPurpose           string            `json:"fund_purpose" validate:"required"`
	BusinessPlan          string            `json:"business_plan"`
	RevenueProjection     *float64          `json:"revenue_projection"`
	MonthlyRevenue        *float64          `json:"monthly_revenue"`
	RequestedTenureMonths int               `json:"requested_tenure_months" validate:"required"`
	CollateralDescription string            `json:"collateral_description"`
	Documents             map[string]string `json:"documents" validate:"required"`
}

// Application List Response
type ApplicationListMobile struct {
	ID          int    `json:"id"`
	ProgramID   int    `json:"program_id"`
	ProgramName string `json:"program_name"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	SubmittedAt string `json:"submitted_at"`
	ExpiredAt   string `json:"expired_at"`
}

// Application Detail Response
type ApplicationDetailMobile struct {
	ID          int                    `json:"id"`
	UMKMID      int                    `json:"umkm_id"`
	ProgramID   int                    `json:"program_id"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	SubmittedAt string                 `json:"submitted_at"`
	ExpiredAt   string                 `json:"expired_at"`
	Documents   []ApplicationDocuments `json:"documents"`
	Histories   []ApplicationHistories `json:"histories"`
	Program     ProgramDetailMobile    `json:"program"`
}

// Notification Response
type NotificationResponse struct {
	ID            int                    `json:"id"`
	Type          string                 `json:"type"`
	Title         string                 `json:"title"`
	Message       string                 `json:"message"`
	IsRead        bool                   `json:"is_read"`
	ReadAt        *string                `json:"read_at,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt     string                 `json:"created_at"`
	ApplicationID *int                   `json:"application_id,omitempty"`
}

// Notification Mark as Read Request
type MarkNotificationReadRequest struct {
	NotificationIDs []int `json:"notification_ids" validate:"required"`
}
