package otp

import (
	"errors"
)

const (
	MESSAGE_TEMPLATE    = "Welcome to TEST. Your request " + MESSAGE_ACTION_CODE + " requires OTP.  Here is your OTP CODE: " + MESSAGE_OTP_CODE
	MESSAGE_OTP_CODE    = "[OTP_CODE]"
	MESSAGE_ACTION_CODE = "[ACTION_NAME]"
	MESSAGE_CLIENT_CODE = "[CLIENT_NAME]"

	VENDOR_WATZAP = "watzap"
	VENDOR_FONNTE = "fonnte"

	COUNTRY_PHONE_SEPARATOR = "."
	COUNTRY_PHONE_PREFIX    = "+"

	ERR_NOT_FOUND_VENDOR      = "not found vendor"
	ERR_PHONE_EMPTY           = "phone empty"
	ERR_NOT_AVAILABLE_SERVICE = "not available service"

	COUNTRY_CODE_INDONESIA = "62"
)

type OTPVendor struct {
	Title          string `db:"title"               form:"title"               json:"title"`
	APIKey         string `db:"api_key"             form:"api_key"             json:"api_key"`
	Url            string `db:"url"                 form:"url"                 json:"url"`
	NumberKey      string `db:"number_key"          form:"number_key"          json:"number_key"`
	DefaultMessage string `db:"default_message"     form:"default_message"     json:"default_message"`
}

type OTPProviderAPI interface {
	SendOtp(phone, otpCode string) (string, error)
}

func InitVendor(vendor OTPVendor) (OTPProviderAPI, error) {
	var otp OTPProviderAPI
	defMessage := MESSAGE_TEMPLATE
	if vendor.DefaultMessage != "" {
		defMessage = vendor.DefaultMessage
	}
	switch vendor.Title {
	case VENDOR_WATZAP:
		var otpW WatzapOtp
		otpW.Vendor = vendor
		otpW.Request.ApiKey = vendor.APIKey
		otpW.Request.NumberKey = vendor.NumberKey
		otpW.Request.Message = defMessage
		otp = &otpW
	case VENDOR_FONNTE:
		var otpF FonnteOtp
		otpF.Vendor = vendor
		otpF.Request.Message = defMessage
		otpF.Request.CountryCode = COUNTRY_CODE_INDONESIA
		otp = &otpF

	default:
		return nil, errors.New(ERR_NOT_FOUND_VENDOR)
	}

	return otp, nil
}
