package service

import (
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/model"
	"context"
)

type VaultDecryptLogService interface {
	GetLogs(ctx context.Context) ([]model.VaultDecryptLog, error)
	GetLogsByUserID(ctx context.Context, userID int) ([]model.VaultDecryptLog, error)
	GetLogsByUMKMID(ctx context.Context, umkmID int) ([]model.VaultDecryptLog, error)
}

const (
	defaultLimit  = 100
	defaultOffset = 0
)

type vaultDecryptLogService struct {
	vaultDecryptLogRepo repository.VaultDecryptLogRepository
}

func NewVaultDecryptLogService(vaultDecryptLogRepo repository.VaultDecryptLogRepository) VaultDecryptLogService {
	return &vaultDecryptLogService{
		vaultDecryptLogRepo: vaultDecryptLogRepo,
	}
}

func (s *vaultDecryptLogService) GetLogs(ctx context.Context) ([]model.VaultDecryptLog, error) {
	return s.vaultDecryptLogRepo.GetLogs(ctx, defaultLimit, defaultOffset)
}

func (s *vaultDecryptLogService) GetLogsByUserID(ctx context.Context, userID int) ([]model.VaultDecryptLog, error) {
	return s.vaultDecryptLogRepo.GetLogsByUserID(ctx, userID, defaultLimit, defaultOffset)
}

func (s *vaultDecryptLogService) GetLogsByUMKMID(ctx context.Context, umkmID int) ([]model.VaultDecryptLog, error) {
	return s.vaultDecryptLogRepo.GetLogsByUMKMID(ctx, umkmID, defaultLimit, defaultOffset)
}