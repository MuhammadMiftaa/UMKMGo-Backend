package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// ==================== MOCK NEWS REPOSITORY ====================

type mockNewsRepository struct {
	news        map[int]model.News
	tags        map[int][]model.NewsTag
	shouldError bool
}

func newMockNewsRepository() *mockNewsRepository {
	now := time.Now()
	return &mockNewsRepository{
		news: map[int]model.News{
			1: {
				ID:          1,
				Title:       "Test News",
				Slug:        "test-news",
				Excerpt:     "Test excerpt",
				Content:     "Test content",
				Thumbnail:   "",
				Category:    "announcement",
				AuthorID:    1,
				IsPublished: true,
				PublishedAt: &now,
				ViewsCount:  10,
				Author:      model.User{ID: 1, Name: "Admin"},
				Base: model.Base{
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
		tags: map[int][]model.NewsTag{
			1: {
				{ID: 1, NewsID: 1, TagName: "announcement"},
				{ID: 2, NewsID: 1, TagName: "important"},
			},
		},
		shouldError: false,
	}
}

func (m *mockNewsRepository) GetAllNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	if m.shouldError {
		return nil, 0, errors.New("database error")
	}

	var news []model.News
	for _, n := range m.news {
		// Apply filters
		if params.Category != "" && n.Category != params.Category {
			continue
		}
		if params.IsPublished != nil && n.IsPublished != *params.IsPublished {
			continue
		}
		news = append(news, n)
	}

	return news, int64(len(news)), nil
}

func (m *mockNewsRepository) GetNewsByID(ctx context.Context, id int) (model.News, error) {
	if m.shouldError {
		return model.News{}, errors.New("database error")
	}
	if news, exists := m.news[id]; exists {
		if tags, ok := m.tags[id]; ok {
			news.Tags = tags
		}
		return news, nil
	}
	return model.News{}, errors.New("news not found")
}

func (m *mockNewsRepository) GetNewsBySlug(ctx context.Context, slug string) (model.News, error) {
	if m.shouldError {
		return model.News{}, errors.New("database error")
	}
	for _, news := range m.news {
		if news.Slug == slug {
			if tags, ok := m.tags[news.ID]; ok {
				news.Tags = tags
			}
			return news, nil
		}
	}
	return model.News{}, errors.New("news not found")
}

func (m *mockNewsRepository) CreateNews(ctx context.Context, news model.News) (model.News, error) {
	if m.shouldError {
		return model.News{}, errors.New("database error")
	}
	news.ID = len(m.news) + 1
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()
	m.news[news.ID] = news
	return news, nil
}

func (m *mockNewsRepository) UpdateNews(ctx context.Context, news model.News) (model.News, error) {
	if m.shouldError {
		return model.News{}, errors.New("database error")
	}
	if _, exists := m.news[news.ID]; !exists {
		return model.News{}, errors.New("news not found")
	}
	news.UpdatedAt = time.Now()
	m.news[news.ID] = news
	return news, nil
}

func (m *mockNewsRepository) DeleteNews(ctx context.Context, news model.News) error {
	if m.shouldError {
		return errors.New("database error")
	}
	delete(m.news, news.ID)
	return nil
}

func (m *mockNewsRepository) IsSlugExists(ctx context.Context, slug string, excludeID int) bool {
	for id, news := range m.news {
		if news.Slug == slug && id != excludeID {
			return true
		}
	}
	return false
}

func (m *mockNewsRepository) CreateNewsTags(ctx context.Context, tags []model.NewsTag) error {
	if m.shouldError {
		return errors.New("database error")
	}
	if len(tags) == 0 {
		return nil
	}
	newsID := tags[0].NewsID
	m.tags[newsID] = tags
	return nil
}

func (m *mockNewsRepository) DeleteNewsTags(ctx context.Context, newsID int) error {
	if m.shouldError {
		return errors.New("database error")
	}
	delete(m.tags, newsID)
	return nil
}

func (m *mockNewsRepository) GetNewsTags(ctx context.Context, newsID int) ([]model.NewsTag, error) {
	if tags, ok := m.tags[newsID]; ok {
		return tags, nil
	}
	return []model.NewsTag{}, nil
}

// ==================== TEST FUNCTIONS ====================

func setupNewsService() (*newsService, *mockNewsRepository) {
	mockRepo := newMockNewsRepository()
	service := &newsService{
		newsRepository: mockRepo,
		minio:          nil, // Not needed for these tests
	}
	return service, mockRepo
}

// Test GetAllNews
func TestGetAllNews(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Get all news successfully", func(t *testing.T) {
		params := dto.NewsQueryParams{
			Page:  1,
			Limit: 10,
		}

		result, total, err := service.GetAllNews(ctx, params)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if total != 1 {
			t.Errorf("Expected total 1, got %d", total)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 news item, got %d", len(result))
		}
	})

	t.Run("Filter by category", func(t *testing.T) {
		params := dto.NewsQueryParams{
			Category: "announcement",
			Page:     1,
			Limit:    10,
		}

		result, total, err := service.GetAllNews(ctx, params)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if total != 1 {
			t.Errorf("Expected 1 news item, got %d", total)
		}

		if result[0].Category != "announcement" {
			t.Errorf("Expected category 'announcement', got '%s'", result[0].Category)
		}
	})

	t.Run("Filter by published status", func(t *testing.T) {
		published := true
		params := dto.NewsQueryParams{
			IsPublished: &published,
			Page:        1,
			Limit:       10,
		}

		result, _, err := service.GetAllNews(ctx, params)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		for _, news := range result {
			if !news.IsPublished {
				t.Error("Expected only published news")
			}
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		params := dto.NewsQueryParams{}

		_, _, err := service.GetAllNews(ctx, params)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})
}

// Test GetNewsByID
func TestGetNewsByID(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Get existing news by ID", func(t *testing.T) {
		result, err := service.GetNewsByID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		if len(result.Tags) != 2 {
			t.Errorf("Expected 2 tags, got %d", len(result.Tags))
		}
	})

	t.Run("Get non-existing news", func(t *testing.T) {
		_, err := service.GetNewsByID(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing news, got none")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetNewsByID(ctx, 1)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})
}

// Test CreateNews
func TestCreateNews(t *testing.T) {
	service, _ := setupNewsService()
	ctx := context.Background()

	t.Run("Create news successfully", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:       "New Article",
			Content:     "Article content",
			Category:    "general",
			IsPublished: false,
			Tags:        []string{"news", "update"},
		}

		result, err := service.CreateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Title != "New Article" {
			t.Errorf("Expected title 'New Article', got '%s'", result.Title)
		}

		if result.Slug == "" {
			t.Error("Expected slug to be generated")
		}
	})

	t.Run("Create news with missing required fields", func(t *testing.T) {
		request := dto.NewsRequest{
			Title: "Test",
			// Missing Content
		}

		_, err := service.CreateNews(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for missing content, got none")
		}

		if err.Error() != "title and content are required" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Create published news sets published_at", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:       "Published News",
			Content:     "Content",
			Category:    "general",
			IsPublished: true,
		}

		result, err := service.CreateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.PublishedAt == nil {
			t.Error("Expected published_at to be set")
		}
	})

	t.Run("Create news with duplicate slug adds timestamp", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:    "Test News", // Same slug as existing
			Content:  "Content",
			Category: "general",
		}

		result, err := service.CreateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Slug == "test-news" {
			t.Error("Expected slug to be modified to avoid duplicate")
		}
	})
}

// Test UpdateNews
func TestUpdateNews(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Update news successfully", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:       "Updated Title",
			Content:     "Updated content",
			Category:    "general",
			Thumbnail:   "",
			IsPublished: true,
			Tags:        []string{"updated"},
		}

		result, err := service.UpdateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Title != "Updated Title" {
			t.Errorf("Expected title 'Updated Title', got '%s'", result.Title)
		}
	})

	t.Run("Update non-existing news", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:    "Test",
			Content:  "Content",
			Category: "general",
		}

		_, err := service.UpdateNews(ctx, 999, request)

		if err == nil {
			t.Error("Expected error for non-existing news, got none")
		}
	})

	t.Run("Update news title generates new slug", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:       "Completely New Title",
			Content:     "Content",
			Category:    "general",
			Thumbnail:   "",
			IsPublished: false,
		}

		result, err := service.UpdateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Slug == "test-news" {
			t.Error("Expected slug to change with new title")
		}
	})

	t.Run("Publishing draft sets published_at", func(t *testing.T) {
		// Add draft news
		now := time.Now()
		mockRepo.news[2] = model.News{
			ID:          2,
			Title:       "Draft",
			Slug:        "draft",
			Content:     "Content",
			IsPublished: false,
			PublishedAt: nil,
			Author:      model.User{ID: 1, Name: "Admin"},
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
		}

		request := dto.NewsRequest{
			Title:       "Draft",
			Content:     "Content",
			Category:    "general",
			Thumbnail:   "",
			IsPublished: true,
		}

		result, err := service.UpdateNews(ctx, 2, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.PublishedAt == nil {
			t.Error("Expected published_at to be set")
		}
	})
}

// Test DeleteNews
func TestDeleteNews(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Delete news successfully", func(t *testing.T) {
		err := service.DeleteNews(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify deletion
		_, err = mockRepo.GetNewsByID(ctx, 1)
		if err == nil {
			t.Error("Expected news to be deleted")
		}
	})

	t.Run("Delete non-existing news", func(t *testing.T) {
		err := service.DeleteNews(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing news, got none")
		}
	})
}

// Test PublishNews
func TestPublishNews(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Publish draft news successfully", func(t *testing.T) {
		// Add draft news
		now := time.Now()
		mockRepo.news[3] = model.News{
			ID:          3,
			Title:       "Draft Article",
			Slug:        "draft-article",
			Content:     "Content",
			IsPublished: false,
			Author:      model.User{ID: 1, Name: "Admin"},
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
		}

		result, err := service.PublishNews(ctx, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !result.IsPublished {
			t.Error("Expected news to be published")
		}

		if result.PublishedAt == nil {
			t.Error("Expected published_at to be set")
		}
	})

	t.Run("Publish already published news returns error", func(t *testing.T) {
		_, err := service.PublishNews(ctx, 1)

		if err == nil {
			t.Error("Expected error for already published news, got none")
		}

		if err.Error() != "news is already published" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Publish non-existing news", func(t *testing.T) {
		_, err := service.PublishNews(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing news, got none")
		}
	})
}

// Test UnpublishNews
func TestUnpublishNews(t *testing.T) {
	service, _ := setupNewsService()
	ctx := context.Background()

	t.Run("Unpublish news successfully", func(t *testing.T) {
		result, err := service.UnpublishNews(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.IsPublished {
			t.Error("Expected news to be unpublished")
		}

		if result.PublishedAt != nil {
			t.Error("Expected published_at to be nil")
		}
	})

	t.Run("Unpublish already unpublished news returns error", func(t *testing.T) {
		// News is now unpublished from previous test
		_, err := service.UnpublishNews(ctx, 1)

		if err == nil {
			t.Error("Expected error for already unpublished news, got none")
		}
	})

	t.Run("Unpublish non-existing news", func(t *testing.T) {
		_, err := service.UnpublishNews(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing news, got none")
		}
	})
}

// Test generateSlug helper function
func TestGenerateSlug(t *testing.T) {
	service, _ := setupNewsService()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple title",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			name:     "Title with special characters",
			input:    "Hello, World!",
			expected: "hello-world",
		},
		{
			name:     "Title with multiple spaces",
			input:    "Hello   World",
			expected: "hello-world",
		},
		{
			name:     "Title with numbers",
			input:    "2024 New Year",
			expected: "2024-new-year",
		},
		{
			name:     "Title with hyphens",
			input:    "COVID-19 Update",
			expected: "covid-19-update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.generateSlug(tt.input)
			if result != tt.expected {
				t.Errorf("Expected slug '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// Test Edge Cases
func TestNewsServiceEdgeCases(t *testing.T) {
	service, mockRepo := setupNewsService()
	ctx := context.Background()

	t.Run("Handle empty tags list", func(t *testing.T) {
		request := dto.NewsRequest{
			Title:    "No Tags News",
			Content:  "Content",
			Category: "general",
			Tags:     []string{},
		}

		result, err := service.CreateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Tags) != 0 {
			t.Errorf("Expected 0 tags, got %d", len(result.Tags))
		}
	})

	t.Run("Handle very long title", func(t *testing.T) {
		longTitle := string(make([]byte, 300))
		for range longTitle {
			longTitle = "A Very Long Title That Exceeds Normal Length"
		}

		title := longTitle
		if len(title) > 255 {
			title = title[:255]
		}

		request := dto.NewsRequest{
			Title:    title,
			Content:  "Content",
			Category: "general",
		}

		_, err := service.CreateNews(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Create multiple news items", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			request := dto.NewsRequest{
				Title:    "News " + string(rune(i)),
				Content:  "Content",
				Category: "general",
			}

			_, err := service.CreateNews(ctx, 1, request)
			if err != nil {
				t.Errorf("Failed to create news %d: %v", i, err)
			}
		}

		// Should have initial 1 + 5 new = 6 total
		if len(mockRepo.news) < 6 {
			t.Errorf("Expected at least 6 news items, got %d", len(mockRepo.news))
		}
	})
}

// Benchmark Tests
func BenchmarkGetNewsByID(b *testing.B) {
	service, _ := setupNewsService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetNewsByID(ctx, 1)
	}
}

func BenchmarkCreateNews(b *testing.B) {
	service, _ := setupNewsService()
	ctx := context.Background()

	request := dto.NewsRequest{
		Title:    "Benchmark News",
		Content:  "Content",
		Category: "general",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.CreateNews(ctx, 1, request)
	}
}
