package repository

import (
	"context"

	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type OTPRepository interface {
	CreateOTP(ctx context.Context, otp model.OTP) error
	GetOTPByPhone(ctx context.Context, phone string) (*model.OTP, error)
	GetOTPByTempToken(ctx context.Context, tempToken string) (*model.OTP, error)
	UpdateOTP(ctx context.Context, otp model.OTP) error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) CreateOTP(ctx context.Context, otp model.OTP) error {
	if err := r.db.Create(&otp).Error; err != nil {
		return err
	}
	return nil
}

func (r *otpRepository) GetOTPByPhone(ctx context.Context, phone string) (*model.OTP, error) {
	var otp model.OTP
	if err := r.db.Where("phone_number = ?", phone).Order("created_at desc").Take(&otp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No record found
		}
		return nil, err // Other error
	}
	return &otp, nil
}

func (r *otpRepository) GetOTPByTempToken(ctx context.Context, tempToken string) (*model.OTP, error) {
	var otp model.OTP
	if err := r.db.Where("temp_token = ?", tempToken).First(&otp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No record found
		}
		return nil, err // Other error
	}
	return &otp, nil
}

func (r *otpRepository) UpdateOTP(ctx context.Context, otp model.OTP) error {
	if err := r.db.Model(&model.OTP{}).Where("phone_number = ? AND otp_code = ?", otp.PhoneNumber, otp.OTPCode).Updates(otp).Error; err != nil {
		return err
	}
	return nil
}
