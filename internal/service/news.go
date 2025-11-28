package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"
)

type NewsService interface {
	GetAllNews(ctx context.Context, params dto.NewsQueryParams) ([]dto.NewsListResponse, int64, error)
	GetNewsByID(ctx context.Context, id int) (dto.NewsResponse, error)
	CreateNews(ctx context.Context, authorID int, request dto.NewsRequest) (dto.NewsResponse, error)
	UpdateNews(ctx context.Context, id int, request dto.NewsRequest) (dto.NewsResponse, error)
	DeleteNews(ctx context.Context, id int) error
	PublishNews(ctx context.Context, id int) (dto.NewsResponse, error)
	UnpublishNews(ctx context.Context, id int) (dto.NewsResponse, error)
}

type newsService struct {
	newsRepository repository.NewsRepository
	minio          *storage.MinIOManager
}

func NewNewsService(newsRepo repository.NewsRepository, minio *storage.MinIOManager) NewsService {
	return &newsService{
		newsRepository: newsRepo,
		minio:          minio,
	}
}

func (s *newsService) GetAllNews(ctx context.Context, params dto.NewsQueryParams) ([]dto.NewsListResponse, int64, error) {
	news, total, err := s.newsRepository.GetAllNews(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	var response []dto.NewsListResponse
	for _, n := range news {
		var publishedAt *string
		if n.PublishedAt != nil {
			pub := n.PublishedAt.Format("2006-01-02 15:04:05")
			publishedAt = &pub
		}

		response = append(response, dto.NewsListResponse{
			ID:          n.ID,
			Title:       n.Title,
			Slug:        n.Slug,
			Excerpt:     n.Excerpt,
			Thumbnail:   n.Thumbnail,
			Category:    n.Category,
			AuthorName:  n.Author.Name,
			IsPublished: n.IsPublished,
			PublishedAt: publishedAt,
			ViewsCount:  n.ViewsCount,
			CreatedAt:   n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response, total, nil
}

func (s *newsService) GetNewsByID(ctx context.Context, id int) (dto.NewsResponse, error) {
	news, err := s.newsRepository.GetNewsByID(ctx, id)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	var publishedAt *string
	if news.PublishedAt != nil {
		pub := news.PublishedAt.Format("2006-01-02 15:04:05")
		publishedAt = &pub
	}

	var tags []string
	for _, tag := range news.Tags {
		tags = append(tags, tag.TagName)
	}

	return dto.NewsResponse{
		ID:          news.ID,
		Title:       news.Title,
		Slug:        news.Slug,
		Excerpt:     news.Excerpt,
		Content:     news.Content,
		Thumbnail:   news.Thumbnail,
		Category:    news.Category,
		AuthorID:    news.AuthorID,
		AuthorName:  news.Author.Name,
		IsPublished: news.IsPublished,
		PublishedAt: publishedAt,
		ViewsCount:  news.ViewsCount,
		CreatedAt:   news.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   news.UpdatedAt.Format("2006-01-02 15:04:05"),
		Tags:        tags,
	}, nil
}

func (s *newsService) CreateNews(ctx context.Context, authorID int, request dto.NewsRequest) (dto.NewsResponse, error) {
	// Validation
	if request.Title == "" || request.Content == "" {
		return dto.NewsResponse{}, errors.New("title and content are required")
	}

	// Generate slug
	slug := s.generateSlug(request.Title)

	// Check if slug exists
	if s.newsRepository.IsSlugExists(ctx, slug, 0) {
		slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix())
	}

	// Upload thumbnail if provided
	thumbnailURL := request.Thumbnail
	if request.Thumbnail != "" && !(strings.HasPrefix(request.Thumbnail, "http") || strings.HasPrefix(request.Thumbnail, "https")) {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: request.Thumbnail,
			BucketName: storage.NewsBucket,
			Prefix:     utils.GenerateFileName(request.Title, "news_thumbnail_"),
			Validation: storage.CreateImageValidationConfig(),
		})
		if err != nil {
			return dto.NewsResponse{}, fmt.Errorf("failed to upload thumbnail: %w", err)
		}
		thumbnailURL = res.URL
	}

	// Create news
	news := model.News{
		Title:       request.Title,
		Slug:        slug,
		Excerpt:     request.Excerpt,
		Content:     request.Content,
		Thumbnail:   thumbnailURL,
		Category:    request.Category,
		AuthorID:    authorID,
		IsPublished: request.IsPublished,
		ViewsCount:  0,
	}

	if request.IsPublished {
		now := time.Now()
		news.PublishedAt = &now
	}

	createdNews, err := s.newsRepository.CreateNews(ctx, news)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	// Create tags
	if len(request.Tags) > 0 {
		var tags []model.NewsTag
		for _, tagName := range request.Tags {
			tags = append(tags, model.NewsTag{
				NewsID:  createdNews.ID,
				TagName: strings.ToLower(strings.TrimSpace(tagName)),
			})
		}
		if err := s.newsRepository.CreateNewsTags(ctx, tags); err != nil {
			return dto.NewsResponse{}, err
		}
	}

	return s.GetNewsByID(ctx, createdNews.ID)
}

func (s *newsService) UpdateNews(ctx context.Context, id int, request dto.NewsRequest) (dto.NewsResponse, error) {
	// Get existing news
	news, err := s.newsRepository.GetNewsByID(ctx, id)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	// Validation
	if request.Title == "" || request.Content == "" {
		return dto.NewsResponse{}, errors.New("title and content are required")
	}

	// Generate new slug if title changed
	newSlug := s.generateSlug(request.Title)
	if newSlug != news.Slug {
		if s.newsRepository.IsSlugExists(ctx, newSlug, id) {
			newSlug = fmt.Sprintf("%s-%d", newSlug, time.Now().Unix())
		}
		news.Slug = newSlug
	}

	// Upload new thumbnail if provided
	if request.Thumbnail != "" && !(strings.HasPrefix(request.Thumbnail, "http") || strings.HasPrefix(request.Thumbnail, "https")) {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: request.Thumbnail,
			BucketName: storage.NewsBucket,
			Prefix:     utils.GenerateFileName(request.Title, "news_thumbnail_"),
			Validation: storage.CreateImageValidationConfig(),
		})
		if err != nil {
			return dto.NewsResponse{}, fmt.Errorf("failed to upload thumbnail: %w", err)
		}

		// Delete old thumbnail
		if news.Thumbnail != "" {
			oldObjectName := storage.ExtractObjectNameFromURL(news.Thumbnail)
			if oldObjectName != "" {
				s.minio.DeleteFile(ctx, storage.NewsBucket, oldObjectName)
			}
		}

		news.Thumbnail = res.URL
	} else if request.Thumbnail != "" {
		news.Thumbnail = request.Thumbnail
	}

	// Update fields
	news.Title = request.Title
	news.Excerpt = request.Excerpt
	news.Content = request.Content
	news.Category = request.Category

	// Handle publishing
	wasPublished := news.IsPublished
	news.IsPublished = request.IsPublished

	if request.IsPublished && !wasPublished {
		now := time.Now()
		news.PublishedAt = &now
	} else if !request.IsPublished && wasPublished {
		news.PublishedAt = nil
	}

	updatedNews, err := s.newsRepository.UpdateNews(ctx, news)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	// Update tags
	if err := s.newsRepository.DeleteNewsTags(ctx, id); err != nil {
		return dto.NewsResponse{}, err
	}

	if len(request.Tags) > 0 {
		var tags []model.NewsTag
		for _, tagName := range request.Tags {
			tags = append(tags, model.NewsTag{
				NewsID:  updatedNews.ID,
				TagName: strings.ToLower(strings.TrimSpace(tagName)),
			})
		}
		if err := s.newsRepository.CreateNewsTags(ctx, tags); err != nil {
			return dto.NewsResponse{}, err
		}
	}

	return s.GetNewsByID(ctx, updatedNews.ID)
}

func (s *newsService) DeleteNews(ctx context.Context, id int) error {
	news, err := s.newsRepository.GetNewsByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete thumbnail from storage
	if news.Thumbnail != "" {
		objectName := storage.ExtractObjectNameFromURL(news.Thumbnail)
		if objectName != "" {
			s.minio.DeleteFile(ctx, storage.NewsBucket, objectName)
		}
	}

	return s.newsRepository.DeleteNews(ctx, news)
}

func (s *newsService) PublishNews(ctx context.Context, id int) (dto.NewsResponse, error) {
	news, err := s.newsRepository.GetNewsByID(ctx, id)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	if news.IsPublished {
		return dto.NewsResponse{}, errors.New("news is already published")
	}

	now := time.Now()
	news.IsPublished = true
	news.PublishedAt = &now

	_, err = s.newsRepository.UpdateNews(ctx, news)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	return s.GetNewsByID(ctx, id)
}

func (s *newsService) UnpublishNews(ctx context.Context, id int) (dto.NewsResponse, error) {
	news, err := s.newsRepository.GetNewsByID(ctx, id)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	if !news.IsPublished {
		return dto.NewsResponse{}, errors.New("news is already unpublished")
	}

	news.IsPublished = false
	news.PublishedAt = nil

	_, err = s.newsRepository.UpdateNews(ctx, news)
	if err != nil {
		return dto.NewsResponse{}, err
	}

	return s.GetNewsByID(ctx, id)
}

// Helper function
func (s *newsService) generateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}
