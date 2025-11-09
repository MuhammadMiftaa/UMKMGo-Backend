package repository

import (
	"context"
	"errors"
	"time"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type ApplicationsRepository interface {
	GetAllApplications(ctx context.Context, filterType string) ([]model.Applications, error)
	GetApplicationByID(ctx context.Context, id int) (model.Applications, error)
	CreateApplication(ctx context.Context, application model.Applications) (model.Applications, error)
	UpdateApplication(ctx context.Context, application model.Applications) (model.Applications, error)
	DeleteApplication(ctx context.Context, application model.Applications) (model.Applications, error)

	// Documents
	CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocuments) error
	GetApplicationDocuments(ctx context.Context, applicationID int) ([]model.ApplicationDocuments, error)
	DeleteApplicationDocuments(ctx context.Context, applicationID int) error

	// Histories
	CreateApplicationHistory(ctx context.Context, history model.ApplicationHistories) error
	GetApplicationHistories(ctx context.Context, applicationID int) ([]model.ApplicationHistories, error)

	// Validations
	GetProgramByID(ctx context.Context, id int) (model.Programs, error)
	GetUMKMByUserID(ctx context.Context, userID int) (model.UMKMS, error)
	IsApplicationExists(ctx context.Context, umkmID, programID int) bool
}

type applicationsRepository struct {
	db *gorm.DB
}

func NewApplicationsRepository(db *gorm.DB) ApplicationsRepository {
	return &applicationsRepository{db}
}

func (repo *applicationsRepository) GetAllApplications(ctx context.Context, filterType string) ([]model.Applications, error) {
	var applications []model.Applications
	query := repo.db.WithContext(ctx).Preload("Program").Preload("UMKM.User").Preload("Documents").Preload("Histories.User").Where("applications.deleted_at IS NULL")

	if filterType != "" {
		query = query.Where("type = ?", filterType)
	}

	err := query.Find(&applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (repo *applicationsRepository) GetApplicationByID(ctx context.Context, id int) (model.Applications, error) {
	var application model.Applications
	err := repo.db.WithContext(ctx).Preload("Program").Preload("UMKM.User").Preload("Documents").Preload("Histories.User").Where("applications.id = ? AND applications.deleted_at IS NULL", id).First(&application).Error
	if err != nil {
		return model.Applications{}, errors.New("application not found")
	}
	return application, nil
}

func (repo *applicationsRepository) CreateApplication(ctx context.Context, application model.Applications) (model.Applications, error) {
	application.SubmittedAt = time.Now()
	application.ExpiredAt = time.Now().AddDate(0, 0, 30) // 30 days from now

	err := repo.db.WithContext(ctx).Create(&application).Error
	if err != nil {
		return model.Applications{}, errors.New("failed to create application")
	}
	return application, nil
}

func (repo *applicationsRepository) UpdateApplication(ctx context.Context, application model.Applications) (model.Applications, error) {
	err := repo.db.WithContext(ctx).Save(&application).Error
	if err != nil {
		return model.Applications{}, errors.New("failed to update application")
	}
	return application, nil
}

func (repo *applicationsRepository) DeleteApplication(ctx context.Context, application model.Applications) (model.Applications, error) {
	err := repo.db.WithContext(ctx).Delete(&application).Error
	if err != nil {
		return model.Applications{}, errors.New("failed to delete application")
	}
	return application, nil
}

func (repo *applicationsRepository) CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocuments) error {
	if len(documents) == 0 {
		return nil
	}
	err := repo.db.WithContext(ctx).Create(&documents).Error
	if err != nil {
		return errors.New("failed to create application documents")
	}
	return nil
}

func (repo *applicationsRepository) GetApplicationDocuments(ctx context.Context, applicationID int) ([]model.ApplicationDocuments, error) {
	var documents []model.ApplicationDocuments
	err := repo.db.WithContext(ctx).Where("application_id = ? AND deleted_at IS NULL", applicationID).Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (repo *applicationsRepository) DeleteApplicationDocuments(ctx context.Context, applicationID int) error {
	err := repo.db.WithContext(ctx).Where("application_id = ?", applicationID).Unscoped().Delete(&model.ApplicationDocuments{}).Error
	if err != nil {
		return errors.New("failed to delete application documents")
	}
	return nil
}

func (repo *applicationsRepository) CreateApplicationHistory(ctx context.Context, history model.ApplicationHistories) error {
	history.ActionedAt = time.Now()
	err := repo.db.WithContext(ctx).Create(&history).Error
	if err != nil {
		return errors.New("failed to create application history")
	}
	return nil
}

func (repo *applicationsRepository) GetApplicationHistories(ctx context.Context, applicationID int) ([]model.ApplicationHistories, error) {
	var histories []model.ApplicationHistories
	err := repo.db.WithContext(ctx).Preload("User").Where("application_id = ? AND deleted_at IS NULL", applicationID).Order("actioned_at DESC").Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (repo *applicationsRepository) GetProgramByID(ctx context.Context, id int) (model.Programs, error) {
	var program model.Programs
	err := repo.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("program not found")
	}
	return program, nil
}

func (repo *applicationsRepository) GetUMKMByUserID(ctx context.Context, userID int) (model.UMKMS, error) {
	var umkm model.UMKMS
	err := repo.db.WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userID).First(&umkm).Error
	if err != nil {
		return model.UMKMS{}, errors.New("UMKM not found")
	}
	return umkm, nil
}

func (repo *applicationsRepository) IsApplicationExists(ctx context.Context, umkmID, programID int) bool {
	var count int64
	repo.db.WithContext(ctx).Model(&model.Applications{}).Where("umkm_id = ? AND program_id = ? AND deleted_at IS NULL AND status NOT IN ('rejected')", umkmID, programID).Count(&count)
	return count > 0
}
