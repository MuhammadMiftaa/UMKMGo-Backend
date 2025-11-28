package repository

import (
	"context"
	"errors"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type NewsRepository interface {
	// Web - Admin Management
	GetAllNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error)
	GetNewsByID(ctx context.Context, id int) (model.News, error)
	GetNewsBySlug(ctx context.Context, slug string) (model.News, error)
	CreateNews(ctx context.Context, news model.News) (model.News, error)
	UpdateNews(ctx context.Context, news model.News) (model.News, error)
	DeleteNews(ctx context.Context, news model.News) error
	IsSlugExists(ctx context.Context, slug string, excludeID int) bool

	// Tags Management
	CreateNewsTags(ctx context.Context, tags []model.NewsTag) error
	DeleteNewsTags(ctx context.Context, newsID int) error
	GetNewsTags(ctx context.Context, newsID int) ([]model.NewsTag, error)

	// Mobile - Public Access
	GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error)
	GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error)
	IncrementViews(ctx context.Context, newsID int) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db}
}

func (r *newsRepository) GetAllNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	var news []model.News
	var total int64

	query := r.db.WithContext(ctx).Model(&model.News{}).Preload("Author").Preload("Tags")

	// Filter by category
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	// Filter by published status
	if params.IsPublished != nil {
		query = query.Where("is_published = ?", *params.IsPublished)
	}

	// Search by title or content
	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ?", searchPattern, searchPattern)
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

	err := query.Order("created_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&news).Error

	return news, total, err
}

func (r *newsRepository) GetNewsByID(ctx context.Context, id int) (model.News, error) {
	var news model.News
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Tags").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&news).Error
	if err != nil {
		return model.News{}, errors.New("news not found")
	}
	return news, nil
}

func (r *newsRepository) GetNewsBySlug(ctx context.Context, slug string) (model.News, error) {
	var news model.News
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Tags").
		Where("slug = ? AND deleted_at IS NULL", slug).
		First(&news).Error
	if err != nil {
		return model.News{}, errors.New("news not found")
	}
	return news, nil
}

func (r *newsRepository) CreateNews(ctx context.Context, news model.News) (model.News, error) {
	err := r.db.WithContext(ctx).Create(&news).Error
	if err != nil {
		return model.News{}, errors.New("failed to create news")
	}
	return news, nil
}

func (r *newsRepository) UpdateNews(ctx context.Context, news model.News) (model.News, error) {
	err := r.db.WithContext(ctx).Save(&news).Error
	if err != nil {
		return model.News{}, errors.New("failed to update news")
	}
	return news, nil
}

func (r *newsRepository) DeleteNews(ctx context.Context, news model.News) error {
	err := r.db.WithContext(ctx).Delete(&news).Error
	if err != nil {
		return errors.New("failed to delete news")
	}
	return nil
}

func (r *newsRepository) IsSlugExists(ctx context.Context, slug string, excludeID int) bool {
	var count int64
	query := r.db.WithContext(ctx).Model(&model.News{}).Where("slug = ? AND deleted_at IS NULL", slug)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	query.Count(&count)
	return count > 0
}

func (r *newsRepository) CreateNewsTags(ctx context.Context, tags []model.NewsTag) error {
	if len(tags) == 0 {
		return nil
	}
	err := r.db.WithContext(ctx).Create(&tags).Error
	if err != nil {
		return errors.New("failed to create news tags")
	}
	return nil
}

func (r *newsRepository) DeleteNewsTags(ctx context.Context, newsID int) error {
	err := r.db.WithContext(ctx).Where("news_id = ?", newsID).Unscoped().Delete(&model.NewsTag{}).Error
	if err != nil {
		return errors.New("failed to delete news tags")
	}
	return nil
}

func (r *newsRepository) GetNewsTags(ctx context.Context, newsID int) ([]model.NewsTag, error) {
	var tags []model.NewsTag
	err := r.db.WithContext(ctx).Where("news_id = ?", newsID).Find(&tags).Error
	return tags, err
}

// Mobile - Public Access
func (r *newsRepository) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
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

func (r *newsRepository) GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error) {
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

func (r *newsRepository) IncrementViews(ctx context.Context, newsID int) error {
	return r.db.WithContext(ctx).
		Model(&model.News{}).
		Where("id = ?", newsID).
		UpdateColumn("views_count", gorm.Expr("views_count + ?", 1)).
		Error
}