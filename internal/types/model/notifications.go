package model

import "time"

type Notification struct {
	ID            int        `json:"id" gorm:"primary_key"`
	UMKMID        int        `json:"umkm_id" gorm:"not null"`
	ApplicationID *int       `json:"application_id"`
	Type          string     `json:"type" gorm:"type:notification_type;not null"`
	Title         string     `json:"title" gorm:"type:varchar(255);not null"`
	Message       string     `json:"message" gorm:"type:text;not null"`
	IsRead        bool       `json:"is_read" gorm:"default:false"`
	ReadAt        *time.Time `json:"read_at"`
	Metadata      string     `json:"metadata" gorm:"type:jsonb"` // Store as JSON string
	Base

	UMKM        UMKM         `json:"umkm" gorm:"foreignKey:UMKMID"`
	Application *Application `json:"application" gorm:"foreignKey:ApplicationID"`
}
