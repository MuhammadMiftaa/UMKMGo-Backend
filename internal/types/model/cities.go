package model

type City struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar(100);not null"`
	ProvinceID int    `json:"province_id" gorm:"not null"`
	Base

	UMKMs    []UMKM   `json:"umkms" gorm:"foreignKey:CityID"`
	Province Province `json:"province" gorm:"foreignKey:ProvinceID"`
}
