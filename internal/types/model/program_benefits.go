package model

type ProgramBenefits struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ProgramID int    `json:"program_id" gorm:"not null"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Base
}