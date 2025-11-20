package model

type TrainingApplication struct {
	ID                 int    `json:"id" gorm:"primary_key"`
	ApplicationID      int    `json:"application_id" gorm:"not null"`
	Motivation         string `json:"motivation" gorm:"type:text;not null"`
	BusinessExperience string `json:"business_experience" gorm:"type:text"`
	LearningObjectives string `json:"learning_objectives" gorm:"type:text"`
	AvailabilityNotes  string `json:"availability_notes" gorm:"type:text"`
	Base

	Application Application `json:"application" gorm:"foreignKey:ApplicationID"`
}