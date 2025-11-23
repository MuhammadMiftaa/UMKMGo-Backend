package model

import "time"

type UMKM struct {
	ID             int       `json:"id" gorm:"primary_key"`
	UserID         int       `json:"user_id" gorm:"not null"`
	BusinessName   string    `json:"business_name" gorm:"type:varchar(100);not null"`
	NIK            string    `json:"nik" gorm:"type:text;not null"`
	Gender         string    `json:"gender" gorm:"type:gender;not null;default:'other'"`
	BirthDate      time.Time `json:"birth_date" gorm:"type:date"`
	Phone          string    `json:"phone" gorm:"type:varchar(15)"`
	Address        string    `json:"address" gorm:"type:text"`
	ProvinceID     int       `json:"province_id" gorm:"not null"`
	CityID         int       `json:"city_id" gorm:"not null"`
	District       string    `json:"district" gorm:"not null"`
	Subdistrict    string    `json:"subdistrict" gorm:"not null"`
	PostalCode     string    `json:"postal_code" gorm:"type:varchar(10)"`
	NIB            string    `json:"nib" gorm:"type:text"`
	NPWP           string    `json:"npwp" gorm:"type:text"`
	RevenueRecord  string    `json:"revenue_record" gorm:"type:text"`
	BusinessPermit string    `json:"business_permit" gorm:"type:text"`
	KartuType      string    `json:"kartu_type" gorm:"type:card_type"`
	KartuNumber    string    `json:"kartu_number" gorm:"type:text"`
	Photo          string    `json:"photo" gorm:"type:text"`
	QRCode         string    `json:"qr_code" gorm:"type:text"`
	Base

	User         User          `json:"user" gorm:"foreignKey:UserID"`
	Province     Province      `json:"province" gorm:"foreignKey:ProvinceID"`
	City         City          `json:"city" gorm:"foreignKey:CityID"`
	Applications []Application `json:"applications" gorm:"foreignKey:UMKMID"`
}
