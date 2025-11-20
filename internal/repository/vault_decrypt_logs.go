package repository

import (
	"context"

	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type VaultDecryptLogRepository interface {
	LogDecrypt(ctx context.Context, log model.VaultDecryptLog) error
	GetLogsByUserID(ctx context.Context, userID int, limit, offset int) ([]model.VaultDecryptLog, error)
	GetLogsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.VaultDecryptLog, error)
}

type vaultDecryptLogRepository struct {
	db *gorm.DB
}

func NewVaultDecryptLogRepository(db *gorm.DB) VaultDecryptLogRepository {
	return &vaultDecryptLogRepository{db}
}

func (r *vaultDecryptLogRepository) LogDecrypt(ctx context.Context, log model.VaultDecryptLog) error {
	return r.db.WithContext(ctx).Create(&log).Error
}

func (r *vaultDecryptLogRepository) GetLogsByUserID(ctx context.Context, userID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	var logs []model.VaultDecryptLog
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("decrypted_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error
	return logs, err
}

func (r *vaultDecryptLogRepository) GetLogsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	var logs []model.VaultDecryptLog
	err := r.db.WithContext(ctx).
		Where("umkm_id = ?", umkmID).
		Order("decrypted_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error
	return logs, err
}
