package model

import "database/sql"

type UMKMS struct {
	ID           int            `json:"id" gorm:"primary_key"`
	UserID       int            `json:"user_id" gorm:"not null"`
	BusinessName string         `json:"business_name" gorm:"type:varchar(100);not null"`
	NIK          string         `json:"nik" gorm:"type:varchar(20);not null"`
	Gender       string         `json:"gender" gorm:"type:gender;not null;default:'other'"`
	BirthDate    sql.NullString `json:"birth_date" gorm:"type:date"`
	Phone        string         `json:"phone" gorm:"type:varchar(15)"`
	Address      string         `json:"address" gorm:"type:text"`
	ProvinceID   int            `json:"province_id" gorm:"not null"`
	CityID       int            `json:"city_id" gorm:"not null"`
	District     int            `json:"district" gorm:"not null"`
	Subdistrict  int            `json:"subdistrict" gorm:"not null"`
	PostalCode   string         `json:"postal_code" gorm:"type:varchar(10)"`
	NIB          string         `json:"nib" gorm:"type:varchar(50)"`
	NPWP         string         `json:"npwp" gorm:"type:varchar(50)"`
	KartuType    string         `json:"kartu_type" gorm:"type:card_type"`
	KartuNumber  string         `json:"kartu_number" gorm:"type:varchar(50)"`
	Base

	User Users `json:"user" gorm:"foreignKey:UserID"`
}
