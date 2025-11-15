package dto

import "UMKMGo-backend/internal/types/model"

type RegisterMobile struct {
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	OTPCode string `json:"otp_code,omitempty"`
}

type ResetPasswordMobile struct {
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type UMKMMobile struct {
	ID           int    `json:"id,omitempty"`
	UserID       int    `json:"user_id,omitempty"`
	Fullname     string `json:"fullname,omitempty"`
	BusinessName string `json:"business_name,omitempty"`
	NIK          string `json:"nik,omitempty"`
	Gender       string `json:"gender,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	Password     string `json:"password,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Address      string `json:"address,omitempty"`
	ProvinceID   int    `json:"province_id,omitempty"`
	CityID       int    `json:"city_id,omitempty"`
	District     string `json:"district,omitempty"`
	Subdistrict  string `json:"subdistrict,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	NIB          string `json:"nib,omitempty"`
	NPWP         string `json:"npwp,omitempty"`
	KartuType    string `json:"kartu_type,omitempty"`
	KartuNumber  string `json:"kartu_number,omitempty"`
}

type UMKMWeb struct {
	ID           int      `json:"id,omitempty"`
	UserID       int      `json:"user_id,omitempty"`
	BusinessName string   `json:"business_name,omitempty"`
	NIK          string   `json:"nik,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	BirthDate    string   `json:"birth_date,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	Address      string   `json:"address,omitempty"`
	ProvinceID   int      `json:"province_id,omitempty"`
	CityID       int      `json:"city_id,omitempty"`
	District     string   `json:"district,omitempty"`
	Subdistrict  string   `json:"subdistrict,omitempty"`
	PostalCode   string   `json:"postal_code,omitempty"`
	NIB          string   `json:"nib,omitempty"`
	NPWP         string   `json:"npwp,omitempty"`
	KartuType    string   `json:"kartu_type,omitempty"`
	KartuNumber  string   `json:"kartu_number,omitempty"`
	User         User     `json:"user"`
	Province     Province `json:"province"`
	City         City     `json:"city"`
}

type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

type City struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ProvinceID int    `json:"province_id,omitempty"`
}

type Province struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RegisterOTP struct {
	State string     `json:"state,omitempty"`
	UMKM  model.UMKM `json:"umkm,omitempty"`
	User  model.User `json:"user,omitempty"`
}

type MetaCityAndProvince struct {
	Provinces []Province `json:"provinces,omitempty"`
	Cities    []City     `json:"cities,omitempty"`
}