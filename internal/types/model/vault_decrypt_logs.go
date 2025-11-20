package model

import "time"

type VaultDecryptLog struct {
	ID           int64     `json:"id" gorm:"primary_key"`
	UserID       int       `json:"user_id" gorm:"not null"`
	UMKMID       *int      `json:"umkm_id"`
	FieldName    string    `json:"field_name" gorm:"type:varchar(100);not null"`
	TableName    string    `json:"table_name" gorm:"type:varchar(100);not null"`
	RecordID     int       `json:"record_id" gorm:"not null"`
	Purpose      string    `json:"purpose" gorm:"type:decrypt_purpose;not null"`
	IPAddress    string    `json:"ip_address" gorm:"type:inet"`
	UserAgent    string    `json:"user_agent" gorm:"type:text"`
	RequestID    string    `json:"request_id" gorm:"type:varchar(50)"`
	Success      bool      `json:"success" gorm:"default:true"`
	ErrorMessage string    `json:"error_message" gorm:"type:text"`
	DecryptedAt  time.Time `json:"decrypted_at" gorm:"default:NOW()"`
}
