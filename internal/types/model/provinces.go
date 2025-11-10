package model

type Province struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Base

	UMKMs  []UMKM `json:"umkms" gorm:"foreignKey:ProvinceID"`
	Cities []City `json:"cities" gorm:"foreignKey:ProvinceID"`
}
