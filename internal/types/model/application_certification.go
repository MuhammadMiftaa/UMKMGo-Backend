package model

type CertificationApplication struct {
	ID                  int    `json:"id" gorm:"primary_key"`
	ApplicationID       int    `json:"application_id" gorm:"not null"`
	BusinessSector      string `json:"business_sector" gorm:"type:varchar(100);not null"`
	ProductOrService    string `json:"product_or_service" gorm:"type:varchar(255);not null"`
	BusinessDescription string `json:"business_description" gorm:"type:text;not null"`
	YearsOperating      *int   `json:"years_operating"`
	CurrentStandards    string `json:"current_standards" gorm:"type:text"`
	CertificationGoals  string `json:"certification_goals" gorm:"type:text;not null"`
	Base

	Application Application `json:"application" gorm:"foreignKey:ApplicationID"`
}
