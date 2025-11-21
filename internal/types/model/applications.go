package model

import "time"

type Application struct {
	ID          int       `json:"id" gorm:"primary_key"`
	UMKMID      int       `json:"umkm_id" gorm:"not null"`
	ProgramID   int       `json:"program_id" gorm:"not null"`
	Type        string    `json:"type" gorm:"type:application_type;not null"`
	Status      string    `json:"status" gorm:"type:application_status;not null;default:'screening'"`
	SubmittedAt time.Time `json:"submitted_at" gorm:"not null;default:NOW()"`
	ExpiredAt   time.Time `json:"expired_at" gorm:"not null"`
	Base

	Documents                []ApplicationDocument     `json:"documents" gorm:"foreignKey:ApplicationID"`
	Histories                []ApplicationHistory      `json:"histories" gorm:"foreignKey:ApplicationID"`
	Program                  Program                   `json:"program" gorm:"foreignKey:ProgramID"`
	UMKM                     UMKM                      `json:"umkm" gorm:"foreignKey:UMKMID"`
	TrainingApplication      *TrainingApplication      `json:"training_application,omitempty" gorm:"foreignKey:ApplicationID"`
	CertificationApplication *CertificationApplication `json:"certification_application,omitempty" gorm:"foreignKey:ApplicationID"`
	FundingApplication       *FundingApplication       `json:"funding_application,omitempty" gorm:"foreignKey:ApplicationID"`
}
