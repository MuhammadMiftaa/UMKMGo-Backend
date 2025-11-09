package model

type ApplicationDocuments struct {
	ID            int    `json:"id" gorm:"primary_key"`
	ApplicationID int    `json:"application_id" gorm:"not null"`
	Type          string `json:"type" gorm:"type:document_type;not null"`
	File          string `json:"file" gorm:"type:text;not null"`
	Base
}
