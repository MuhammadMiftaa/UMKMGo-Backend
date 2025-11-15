package model

import "time"

type OTP struct {
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	OTPCode     string    `json:"otp_code"`
	TempToken   *string   `json:"temp_token"`
	Status      string    `json:"status"`
	ExpiresAt   time.Time `json:"expires_at"`
}
