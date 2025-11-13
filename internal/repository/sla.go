package repository

import (
	"context"
	"errors"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type SLARepository interface {
	GetSLAByStatus(ctx context.Context, status string) (model.SLA, error)
	UpdateSLA(ctx context.Context, sla model.SLA) (model.SLA, error)
	GetApplicationsForExport(ctx context.Context, applicationType string) ([]model.Application, error)
	GetProgramsForExport(ctx context.Context, applicationType string) ([]model.Program, error)
}

type slaRepository struct {
	db *gorm.DB
}

func NewSLARepository(db *gorm.DB) SLARepository {
	return &slaRepository{db}
}

func (repo *slaRepository) GetSLAByStatus(ctx context.Context, status string) (model.SLA, error) {
	var sla model.SLA
	err := repo.db.WithContext(ctx).Where("status = ? AND deleted_at IS NULL", status).First(&sla).Error
	if err != nil {
		return model.SLA{}, errors.New("SLA not found")
	}
	return sla, nil
}

func (repo *slaRepository) UpdateSLA(ctx context.Context, sla model.SLA) (model.SLA, error) {
	err := repo.db.WithContext(ctx).Save(&sla).Error
	if err != nil {
		return model.SLA{}, errors.New("failed to update SLA")
	}
	return sla, nil
}

func (repo *slaRepository) GetApplicationsForExport(ctx context.Context, applicationType string) ([]model.Application, error) {
	var applications []model.Application
	query := repo.db.WithContext(ctx).
		Preload("Program").
		Preload("UMKM.User").
		Preload("UMKM.City.Province").
		Where("applications.deleted_at IS NULL")

	if applicationType != "all" {
		query = query.Where("type = ?", applicationType)
	}

	err := query.Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (repo *slaRepository) GetProgramsForExport(ctx context.Context, applicationType string) ([]model.Program, error) {
	var programs []model.Program
	query := repo.db.WithContext(ctx).
		Preload("Users").
		Where("deleted_at IS NULL")

	if applicationType != "all" {
		query = query.Where("type = ?", applicationType)
	}

	err := query.Find(&programs).Error
	if err != nil {
		return nil, err
	}

	return programs, nil
}
