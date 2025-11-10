package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type ApplicationsRepository interface {
	GetAllApplications(ctx context.Context, filterType string) ([]model.Application, error)
	GetApplicationByID(ctx context.Context, id int) (model.Application, error)
	CreateApplication(ctx context.Context, application model.Application) (model.Application, error)
	UpdateApplication(ctx context.Context, application model.Application) (model.Application, error)
	DeleteApplication(ctx context.Context, application model.Application) (model.Application, error)

	// Documents
	CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocument) error
	GetApplicationDocuments(ctx context.Context, applicationID int) ([]model.ApplicationDocument, error)
	DeleteApplicationDocuments(ctx context.Context, applicationID int) error

	// Histories
	CreateApplicationHistory(ctx context.Context, history model.ApplicationHistory) error
	GetApplicationHistories(ctx context.Context, applicationID int) ([]model.ApplicationHistory, error)

	// Validations
	GetProgramByID(ctx context.Context, id int) (model.Program, error)
	GetUMKMByUserID(ctx context.Context, userID int) (model.UMKM, error)
	IsApplicationExists(ctx context.Context, umkmID, programID int) bool
}

type applicationsRepository struct {
	db *gorm.DB
}

func NewApplicationsRepository(db *gorm.DB) ApplicationsRepository {
	return &applicationsRepository{db}
}

func (repo *applicationsRepository) GetAllApplications(ctx context.Context, filterType string) ([]model.Application, error) {
	var applications []model.Application
	query := repo.db.WithContext(ctx).
	Debug().
	Preload("Program").
	Preload("UMKM.User").
	Preload("UMKM.City.Province").
	Preload("Documents").
	Preload("Histories.User").
	Where("applications.deleted_at IS NULL")

	if filterType != "" {
		query = query.Where("type = ?", filterType)
	}
	log.Printf("Query: %s", query.Statement.SQL.String())
	err := query.WithContext(ctx).Find(&applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (repo *applicationsRepository) GetApplicationByID(ctx context.Context, id int) (model.Application, error) {
	var application model.Application
	err := repo.db.WithContext(ctx).
	Debug().
	Preload("Program").
	Preload("UMKM.User").
	Preload("UMKM.City.Province").
	Preload("Documents").
	Preload("Histories.User").
	Where("applications.id = ? AND applications.deleted_at IS NULL", id).
	First(&application).Error
	if err != nil {
		return model.Application{}, errors.New("application not found")
	}
	return application, nil
}

func (repo *applicationsRepository) CreateApplication(ctx context.Context, application model.Application) (model.Application, error) {
	application.SubmittedAt = time.Now()
	application.ExpiredAt = time.Now().AddDate(0, 0, 30) // 30 days from now

	err := repo.db.WithContext(ctx).Create(&application).Error
	if err != nil {
		return model.Application{}, errors.New("failed to create application")
	}
	return application, nil
}

func (repo *applicationsRepository) UpdateApplication(ctx context.Context, application model.Application) (model.Application, error) {
	err := repo.db.WithContext(ctx).Save(&application).Error
	if err != nil {
		return model.Application{}, errors.New("failed to update application")
	}
	return application, nil
}

func (repo *applicationsRepository) DeleteApplication(ctx context.Context, application model.Application) (model.Application, error) {
	err := repo.db.WithContext(ctx).Delete(&application).Error
	if err != nil {
		return model.Application{}, errors.New("failed to delete application")
	}
	return application, nil
}

func (repo *applicationsRepository) CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocument) error {
	if len(documents) == 0 {
		return nil
	}
	err := repo.db.WithContext(ctx).Create(&documents).Error
	if err != nil {
		return errors.New("failed to create application documents")
	}
	return nil
}

func (repo *applicationsRepository) GetApplicationDocuments(ctx context.Context, applicationID int) ([]model.ApplicationDocument, error) {
	var documents []model.ApplicationDocument
	err := repo.db.WithContext(ctx).Where("application_id = ? AND deleted_at IS NULL", applicationID).Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (repo *applicationsRepository) DeleteApplicationDocuments(ctx context.Context, applicationID int) error {
	err := repo.db.WithContext(ctx).Where("application_id = ?", applicationID).Unscoped().Delete(&model.ApplicationDocument{}).Error
	if err != nil {
		return errors.New("failed to delete application documents")
	}
	return nil
}

func (repo *applicationsRepository) CreateApplicationHistory(ctx context.Context, history model.ApplicationHistory) error {
	history.ActionedAt = time.Now()
	err := repo.db.WithContext(ctx).Create(&history).Error
	if err != nil {
		return errors.New("failed to create application history")
	}
	return nil
}

func (repo *applicationsRepository) GetApplicationHistories(ctx context.Context, applicationID int) ([]model.ApplicationHistory, error) {
	var histories []model.ApplicationHistory
	err := repo.db.WithContext(ctx).Preload("User").Where("application_id = ? AND deleted_at IS NULL", applicationID).Order("actioned_at DESC").Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (repo *applicationsRepository) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	var program model.Program
	err := repo.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&program).Error
	if err != nil {
		return model.Program{}, errors.New("program not found")
	}
	return program, nil
}

func (repo *applicationsRepository) GetUMKMByUserID(ctx context.Context, userID int) (model.UMKM, error) {
	var umkm model.UMKM
	err := repo.db.WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userID).First(&umkm).Error
	if err != nil {
		return model.UMKM{}, errors.New("UMKM not found")
	}
	return umkm, nil
}

func (repo *applicationsRepository) IsApplicationExists(ctx context.Context, umkmID, programID int) bool {
	var count int64
	repo.db.WithContext(ctx).Model(&model.Application{}).Where("umkm_id = ? AND program_id = ? AND deleted_at IS NULL AND status NOT IN ('rejected')", umkmID, programID).Count(&count)
	return count > 0
}
