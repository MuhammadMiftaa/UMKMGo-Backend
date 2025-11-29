package repository

import (
	"context"
	"errors"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type MobileRepository interface {
	// Programs
	GetProgramsByType(ctx context.Context, programType string) ([]model.Program, error)
	GetProgramDetailByID(ctx context.Context, id int) (model.Program, error)

	// UMKM Profile
	GetUMKMProfileByID(ctx context.Context, userID int) (model.UMKM, error)
	UpdateUMKMProfile(ctx context.Context, umkm model.UMKM) (model.UMKM, error)

	// Documents
	UpdateUMKMDocument(ctx context.Context, umkmID int, field, value string) error

	// Applications
	CreateApplication(ctx context.Context, application model.Application) (model.Application, error)
	CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocument) error
	CreateApplicationHistory(ctx context.Context, history model.ApplicationHistory) error
	CreateTrainingApplication(ctx context.Context, training model.TrainingApplication) error
	CreateCertificationApplication(ctx context.Context, certification model.CertificationApplication) error
	CreateFundingApplication(ctx context.Context, funding model.FundingApplication) error
	GetApplicationsByUMKMID(ctx context.Context, umkmID int) ([]model.Application, error)
	GetApplicationDetailByID(ctx context.Context, id int) (model.Application, error)
	DeleteApplicationDocumentsByApplicationID(ctx context.Context, applicationID int) error

	// Validations
	GetProgramByID(ctx context.Context, id int) (model.Program, error)
	GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirement, error)
	IsApplicationExists(ctx context.Context, umkmID, programID int) bool

	// News
	GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error)
	GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error)
	IncrementViews(ctx context.Context, newsID int) error
}

type mobileRepository struct {
	db *gorm.DB
}

func NewMobileRepository(db *gorm.DB) MobileRepository {
	return &mobileRepository{db}
}

// Programs
func (r *mobileRepository) GetProgramsByType(ctx context.Context, programType string) ([]model.Program, error) {
	var programs []model.Program
	err := r.db.WithContext(ctx).
		Where("type = ? AND is_active = ? AND deleted_at IS NULL", programType, true).
		Order("created_at DESC").
		Find(&programs).Error
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (r *mobileRepository) GetProgramDetailByID(ctx context.Context, id int) (model.Program, error) {
	var program model.Program
	err := r.db.WithContext(ctx).
		Preload("Users").
		Where("id = ? AND is_active = ? AND deleted_at IS NULL", id, true).
		First(&program).Error
	if err != nil {
		return model.Program{}, errors.New("program not found")
	}
	return program, nil
}

// UMKM Profile
func (r *mobileRepository) GetUMKMProfileByID(ctx context.Context, userID int) (model.UMKM, error) {
	var umkm model.UMKM
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Province").
		Preload("City").
		Where("id = ? AND deleted_at IS NULL", userID).
		First(&umkm).Error
	if err != nil {
		return model.UMKM{}, errors.New("UMKM profile not found")
	}
	return umkm, nil
}

func (r *mobileRepository) UpdateUMKMProfile(ctx context.Context, umkm model.UMKM) (model.UMKM, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Save(&umkm).Error; err != nil {
			return errors.New("failed to update UMKM profile")
		}

		// Update user name if provided
		if umkm.User.Name != "" {
			if err := tx.Model(&umkm.User).Where("id = ?", umkm.User.ID).Update("name", umkm.User.Name).Error; err != nil {
				return errors.New("failed to update user name")
			}
		}

		return nil
	})
	if err != nil {
		return model.UMKM{}, errors.New("failed to update UMKM profile")
	}
	return umkm, nil
}

// Documents
func (r *mobileRepository) UpdateUMKMDocument(ctx context.Context, umkmID int, field, value string) error {
	err := r.db.WithContext(ctx).
		Model(&model.UMKM{}).
		Where("id = ?", umkmID).
		Update(field, value).Error
	if err != nil {
		return errors.New("failed to update document")
	}
	return nil
}

// Applications
func (r *mobileRepository) CreateApplication(ctx context.Context, application model.Application) (model.Application, error) {
	err := r.db.WithContext(ctx).Create(&application).Error
	if err != nil {
		return model.Application{}, errors.New("failed to create application")
	}
	return application, nil
}

func (r *mobileRepository) CreateApplicationDocuments(ctx context.Context, documents []model.ApplicationDocument) error {
	if len(documents) == 0 {
		return nil
	}
	err := r.db.WithContext(ctx).Create(&documents).Error
	if err != nil {
		return errors.New("failed to create application documents")
	}
	return nil
}

func (r *mobileRepository) CreateApplicationHistory(ctx context.Context, history model.ApplicationHistory) error {
	err := r.db.WithContext(ctx).Create(&history).Error
	if err != nil {
		return errors.New("failed to create application history")
	}
	return nil
}

func (r *mobileRepository) CreateTrainingApplication(ctx context.Context, training model.TrainingApplication) error {
	err := r.db.WithContext(ctx).Create(&training).Error
	if err != nil {
		return errors.New("failed to create training application")
	}
	return nil
}

func (r *mobileRepository) CreateCertificationApplication(ctx context.Context, certification model.CertificationApplication) error {
	err := r.db.WithContext(ctx).Create(&certification).Error
	if err != nil {
		return errors.New("failed to create certification application")
	}
	return nil
}

func (r *mobileRepository) CreateFundingApplication(ctx context.Context, funding model.FundingApplication) error {
	err := r.db.WithContext(ctx).Create(&funding).Error
	if err != nil {
		return errors.New("failed to create funding application")
	}
	return nil
}

func (r *mobileRepository) GetApplicationsByUMKMID(ctx context.Context, umkmID int) ([]model.Application, error) {
	var applications []model.Application
	err := r.db.WithContext(ctx).
		Preload("Program").
		Preload("TrainingApplication").
		Preload("CertificationApplication").
		Preload("FundingApplication").
		Where("umkm_id = ? AND deleted_at IS NULL", umkmID).
		Order("submitted_at DESC").
		Find(&applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *mobileRepository) GetApplicationDetailByID(ctx context.Context, id int) (model.Application, error) {
	var application model.Application
	err := r.db.WithContext(ctx).
		Preload("Program").
		Preload("Documents").
		Preload("Histories.User").
		Preload("TrainingApplication").
		Preload("CertificationApplication").
		Preload("FundingApplication").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&application).Error
	if err != nil {
		return model.Application{}, errors.New("application not found")
	}
	return application, nil
}

func (r *mobileRepository) DeleteApplicationDocumentsByApplicationID(ctx context.Context, applicationID int) error {
	err := r.db.WithContext(ctx).
		Where("application_id = ?", applicationID).
		Delete(&model.ApplicationDocument{}).Error
	if err != nil {
		return errors.New("failed to delete application documents")
	}
	return nil
}

// Validations
func (r *mobileRepository) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	var program model.Program
	err := r.db.WithContext(ctx).
		Where("id = ? AND is_active = ? AND deleted_at IS NULL", id, true).
		First(&program).Error
	if err != nil {
		return model.Program{}, errors.New("program not found")
	}
	return program, nil
}

func (r *mobileRepository) GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirement, error) {
	var requirements []model.ProgramRequirement
	err := r.db.WithContext(ctx).
		Where("program_id = ? AND deleted_at IS NULL", programID).
		Find(&requirements).Error
	if err != nil {
		return nil, err
	}
	return requirements, nil
}

func (r *mobileRepository) IsApplicationExists(ctx context.Context, umkmID, programID int) bool {
	var count int64
	r.db.WithContext(ctx).
		Model(&model.Application{}).
		Where("umkm_id = ? AND program_id = ? AND deleted_at IS NULL AND status NOT IN ('rejected')", umkmID, programID).
		Count(&count)
	return count > 0
}


func (r *mobileRepository) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	var news []model.News
	var total int64

	query := r.db.WithContext(ctx).
		Model(&model.News{}).
		Preload("Author").
		Preload("Tags").
		Where("is_published = ? AND deleted_at IS NULL", true)

	// Filter by category
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	// Search by title
	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("title ILIKE ?", searchPattern)
	}

	// Filter by tag
	if params.Tag != "" {
		query = query.Joins("JOIN news_tags ON news_tags.news_id = news.id").
			Where("news_tags.tag_name = ?", params.Tag)
	}

	// Count total
	query.Count(&total)

	// Pagination
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}
	offset := (params.Page - 1) * params.Limit

	err := query.Order("published_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&news).Error

	return news, total, err
}

func (r *mobileRepository) GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error) {
	var news model.News
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Tags").
		Where("slug = ? AND is_published = ? AND deleted_at IS NULL", slug, true).
		First(&news).Error
	if err != nil {
		return model.News{}, errors.New("news not found")
	}
	return news, nil
}

func (r *mobileRepository) IncrementViews(ctx context.Context, newsID int) error {
	return r.db.WithContext(ctx).
		Model(&model.News{}).
		Where("id = ?", newsID).
		UpdateColumn("views_count", gorm.Expr("views_count + ?", 1)).
		Error
}