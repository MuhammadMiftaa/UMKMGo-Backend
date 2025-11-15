package otp

import (
	"log"

	"UMKMGo-backend/config/env"
)

func SendOTP(phone, otpCode string) (string, error) {
	vendor, err := InitVendor(OTPVendor{
		Title:  VENDOR_FONNTE,
		APIKey: env.Cfg.Fonnte.Token,
	})
	if err != nil {
		log.Println("Error when initialize vendor, Err : ", err.Error())
		return "", err
	}

	resp, err := vendor.SendOtp(phone, otpCode)
	if err != nil {
		log.Println("error when send OTP . Err : ", err.Error())
		return "", err
	}
	return resp, nil
}
