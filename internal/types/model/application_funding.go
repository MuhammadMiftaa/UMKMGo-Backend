package model

type FundingApplication struct {
	ID                    int      `json:"id" gorm:"primary_key"`
	ApplicationID         int      `json:"application_id" gorm:"not null"`
	BusinessSector        string   `json:"business_sector" gorm:"type:varchar(100);not null"`
	BusinessDescription   string   `json:"business_description" gorm:"type:text;not null"`
	YearsOperating        *int     `json:"years_operating"`
	RequestedAmount       float64  `json:"requested_amount" gorm:"type:numeric(15,2);not null"`
	FundPurpose           string   `json:"fund_purpose" gorm:"type:text;not null"`
	BusinessPlan          string   `json:"business_plan" gorm:"type:text"`
	RevenueProjection     *float64 `json:"revenue_projection" gorm:"type:numeric(15,2)"`
	MonthlyRevenue        *float64 `json:"monthly_revenue" gorm:"type:numeric(15,2)"`
	RequestedTenureMonths int      `json:"requested_tenure_months" gorm:"not null"`
	CollateralDescription string   `json:"collateral_description" gorm:"type:text"`
	Base

	Application Application `json:"application" gorm:"foreignKey:ApplicationID"`
}
