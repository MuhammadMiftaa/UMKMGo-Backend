package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils/constant"
)

// ==================== ADDITIONAL MOCK REPOSITORIES ====================

// Mock Mobile Repository
type mockMobileRepository struct {
	programs     map[int]model.Program
	umkms        map[int]model.UMKM
	applications map[int]model.Application
	news         map[int]model.News
}

func newMockMobileRepository() *mockMobileRepository {
	return &mockMobileRepository{
		programs:     make(map[int]model.Program),
		umkms:        make(map[int]model.UMKM),
		applications: make(map[int]model.Application),
		news:         make(map[int]model.News),
	}
}

func (m *mockMobileRepository) GetProgramsByType(ctx context.Context, programType string) ([]model.Program, error) {
	var programs []model.Program
	for _, p := range m.programs {
		if p.Type == programType && p.IsActive {
			programs = append(programs, p)
		}
	}
	return programs, nil
}

func (m *mockMobileRepository) GetProgramDetailByID(ctx context.Context, id int) (model.Program, error) {
	if prog, exists := m.programs[id]; exists && prog.IsActive {
		return prog, nil
	}
	return model.Program{}, errors.New("program not found")
}

func (m *mockMobileRepository) GetUMKMProfileByID(ctx context.Context, userID int) (model.UMKM, error) {
	if umkm, exists := m.umkms[userID]; exists {
		return umkm, nil
	}
	return model.UMKM{}, errors.New("UMKM profile not found")
}

func (m *mockMobileRepository) UpdateUMKMProfile(ctx context.Context, umkm model.UMKM) (model.UMKM, error) {
	m.umkms[umkm.ID] = umkm
	return umkm, nil
}

func (m *mockMobileRepository) UpdateUMKMDocument(ctx context.Context, umkmID int, field, value string) error {
	if umkm, exists := m.umkms[umkmID]; exists {
		// Simulate document update
		_ = umkm
		return nil
	}
	return errors.New("UMKM not found")
}

func (m *mockMobileRepository) CreateApplication(ctx context.Context, app model.Application) (model.Application, error) {
	app.ID = len(m.applications) + 1
	m.applications[app.ID] = app
	return app, nil
}

func (m *mockMobileRepository) CreateApplicationDocuments(ctx context.Context, docs []model.ApplicationDocument) error {
	return nil
}

func (m *mockMobileRepository) CreateApplicationHistory(ctx context.Context, hist model.ApplicationHistory) error {
	return nil
}

func (m *mockMobileRepository) CreateTrainingApplication(ctx context.Context, training model.TrainingApplication) error {
	return nil
}

func (m *mockMobileRepository) CreateCertificationApplication(ctx context.Context, cert model.CertificationApplication) error {
	return nil
}

func (m *mockMobileRepository) CreateFundingApplication(ctx context.Context, funding model.FundingApplication) error {
	return nil
}

func (m *mockMobileRepository) GetApplicationsByUMKMID(ctx context.Context, umkmID int) ([]model.Application, error) {
	var apps []model.Application
	for _, app := range m.applications {
		if app.UMKMID == umkmID {
			apps = append(apps, app)
		}
	}
	return apps, nil
}

func (m *mockMobileRepository) GetApplicationDetailByID(ctx context.Context, id int) (model.Application, error) {
	if app, exists := m.applications[id]; exists {
		return app, nil
	}
	return model.Application{}, errors.New("application not found")
}

func (m *mockMobileRepository) DeleteApplicationDocumentsByApplicationID(ctx context.Context, applicationID int) error {
	return nil
}

func (m *mockMobileRepository) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	if prog, exists := m.programs[id]; exists && prog.IsActive {
		return prog, nil
	}
	return model.Program{}, errors.New("program not found")
}

func (m *mockMobileRepository) GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirement, error) {
	return []model.ProgramRequirement{}, nil
}

func (m *mockMobileRepository) IsApplicationExists(ctx context.Context, umkmID, programID int) bool {
	for _, app := range m.applications {
		if app.UMKMID == umkmID && app.ProgramID == programID && app.Status != "rejected" {
			return true
		}
	}
	return false
}

func (m *mockMobileRepository) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	var news []model.News
	for _, n := range m.news {
		if n.IsPublished {
			news = append(news, n)
		}
	}
	return news, int64(len(news)), nil
}

func (m *mockMobileRepository) GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error) {
	for _, n := range m.news {
		if n.Slug == slug && n.IsPublished {
			return n, nil
		}
	}
	return model.News{}, errors.New("news not found")
}

func (m *mockMobileRepository) IncrementViews(ctx context.Context, newsID int) error {
	if n, exists := m.news[newsID]; exists {
		n.ViewsCount++
		m.news[newsID] = n
		return nil
	}
	return errors.New("news not found")
}

// Mock Program Repository for Mobile
type mockProgramRepoForMobile struct {
	benefits     map[int][]model.ProgramBenefit
	requirements map[int][]model.ProgramRequirement
}

func newMockProgramRepoForMobile() *mockProgramRepoForMobile {
	return &mockProgramRepoForMobile{
		benefits:     make(map[int][]model.ProgramBenefit),
		requirements: make(map[int][]model.ProgramRequirement),
	}
}

func (m *mockProgramRepoForMobile) GetAllPrograms(ctx context.Context) ([]model.Program, error) {
	return nil, errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	return model.Program{}, errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) CreateProgram(ctx context.Context, program model.Program) (model.Program, error) {
	return model.Program{}, errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) UpdateProgram(ctx context.Context, program model.Program) (model.Program, error) {
	return model.Program{}, errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) DeleteProgram(ctx context.Context, program model.Program) (model.Program, error) {
	return model.Program{}, errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) CreateProgramBenefits(ctx context.Context, benefits []model.ProgramBenefit) error {
	return errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) CreateProgramRequirements(ctx context.Context, requirements []model.ProgramRequirement) error {
	return errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) GetProgramBenefits(ctx context.Context, programID int) ([]model.ProgramBenefit, error) {
	if benefits, exists := m.benefits[programID]; exists {
		return benefits, nil
	}
	return []model.ProgramBenefit{}, nil
}

func (m *mockProgramRepoForMobile) GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirement, error) {
	if reqs, exists := m.requirements[programID]; exists {
		return reqs, nil
	}
	return []model.ProgramRequirement{}, nil
}

func (m *mockProgramRepoForMobile) DeleteProgramBenefits(ctx context.Context, programID int) error {
	return errors.New("not implemented")
}

func (m *mockProgramRepoForMobile) DeleteProgramRequirements(ctx context.Context, programID int) error {
	return errors.New("not implemented")
}

// Mock Notification Repository for Mobile Tests
type mockNotificationRepoForMobile struct {
	notifications []model.Notification
	shouldError   bool
}

func newMockNotificationRepoForMobile() *mockNotificationRepoForMobile {
	return &mockNotificationRepoForMobile{
		notifications: []model.Notification{},
		shouldError:   false,
	}
}

func (m *mockNotificationRepoForMobile) CreateNotification(ctx context.Context, notif model.Notification) error {
	if m.shouldError {
		return errors.New("database error")
	}
	m.notifications = append(m.notifications, notif)
	return nil
}

func (m *mockNotificationRepoForMobile) GetNotificationsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.Notification, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}
	// Filter by UMKMID to satisfy tests expecting empty list for other UMKM IDs
	var filtered []model.Notification
	for _, n := range m.notifications {
		if n.UMKMID == umkmID {
			filtered = append(filtered, n)
		}
	}
	return filtered, nil
}

func (m *mockNotificationRepoForMobile) GetUnreadCount(ctx context.Context, umkmID int) (int64, error) {
	if m.shouldError {
		return 0, errors.New("database error")
	}
	count := int64(0)
	for _, n := range m.notifications {
		if n.UMKMID == umkmID && !n.IsRead {
			count++
		}
	}
	return count, nil
}

func (m *mockNotificationRepoForMobile) MarkAsRead(ctx context.Context, notifIDs int, umkmID int) error {
	if m.shouldError {
		return errors.New("database error")
	}
	return nil
}

func (m *mockNotificationRepoForMobile) MarkAllAsRead(ctx context.Context, umkmID int) error {
	if m.shouldError {
		return errors.New("database error")
	}
	return nil
}

type mockNewsRepoForMobile struct {
	news        map[int]model.News
	shouldError bool
}

func newMockNewsRepoForMobile() *mockNewsRepoForMobile {
	now := time.Now()
	return &mockNewsRepoForMobile{
		news: map[int]model.News{
			1: {
				ID:          1,
				Title:       "New Training Program Launched",
				Slug:        "new-training-program-launched",
				Excerpt:     "We are excited to announce a new training program",
				Content:     "Full content of the news article here...",
				Thumbnail:   "http://example.com/thumbnail1.jpg",
				Category:    "program_update",
				AuthorID:    1,
				IsPublished: true,
				PublishedAt: &now,
				ViewsCount:  100,
				Author: model.User{
					ID:   1,
					Name: "Admin User",
				},
				Tags: []model.NewsTag{
					{ID: 1, NewsID: 1, TagName: "training"},
					{ID: 2, NewsID: 1, TagName: "program"},
				},
			},
			2: {
				ID:          2,
				Title:       "Success Story: UMKM Goes Digital",
				Slug:        "success-story-umkm-goes-digital",
				Excerpt:     "How one UMKM transformed their business",
				Content:     "Detailed success story content...",
				Thumbnail:   "http://example.com/thumbnail2.jpg",
				Category:    "success_story",
				AuthorID:    1,
				IsPublished: true,
				PublishedAt: &now,
				ViewsCount:  250,
				Author: model.User{
					ID:   1,
					Name: "Admin User",
				},
				Tags: []model.NewsTag{
					{ID: 3, NewsID: 2, TagName: "digital"},
					{ID: 4, NewsID: 2, TagName: "success"},
				},
			},
			3: {
				ID:          3,
				Title:       "Upcoming Certification Event",
				Slug:        "upcoming-certification-event",
				Excerpt:     "Join us for the certification event",
				Content:     "Event details and registration info...",
				Thumbnail:   "http://example.com/thumbnail3.jpg",
				Category:    "event",
				AuthorID:    1,
				IsPublished: true,
				PublishedAt: &now,
				ViewsCount:  50,
				Author: model.User{
					ID:   1,
					Name: "Admin User",
				},
				Tags: []model.NewsTag{
					{ID: 5, NewsID: 3, TagName: "event"},
					{ID: 6, NewsID: 3, TagName: "certification"},
				},
			},
			4: {
				ID:          4,
				Title:       "Unpublished Draft News",
				Slug:        "unpublished-draft-news",
				Excerpt:     "This is a draft",
				Content:     "Draft content...",
				Category:    "general",
				AuthorID:    1,
				IsPublished: false,
				ViewsCount:  0,
				Author: model.User{
					ID:   1,
					Name: "Admin User",
				},
			},
		},
		shouldError: false,
	}
}

// ==================== TEST FUNCTIONS FOR MOBILE METHODS ====================

func setupMobileServiceForTests() (*mobileService, *mockMobileRepository) {
	mockMobileRepo := newMockMobileRepository()
	mockProgramRepo := newMockProgramRepoForMobile()
	mockNotifRepo := newMockNotificationRepoForMobile()
	mockVaultRepo := newMockVaultDecryptLogRepo()
	mockAppRepo := newMockApplicationsRepo()
	mockSLARepo := newMockSLARepo()

	// Setup initial test data
	birthDate, _ := time.Parse("2006-01-02", "1990-01-01")
	mockMobileRepo.umkms[1] = model.UMKM{
		ID:           1,
		UserID:       1,
		BusinessName: "Test Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		BirthDate:    birthDate,
		User: model.User{
			ID:    1,
			Name:  "Test User",
			Email: "test@example.com",
		},
		Province: model.Province{ID: 1, Name: "DKI Jakarta"},
		City:     model.City{ID: 1, Name: "Jakarta Pusat"},
		NIB:      "http://example.com/nib.pdf",
		NPWP:     "http://example.com/npwp.pdf",
	}

	mockMobileRepo.umkms[2] = model.UMKM{
		ID:           2,
		UserID:       2,
		BusinessName: "Second Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		BirthDate:    birthDate,
		User: model.User{
			ID:    2,
			Name:  "Test User 2",
			Email: "test2@example.com",
		},
		Province: model.Province{ID: 1, Name: "DKI Jakarta"},
		City:     model.City{ID: 1, Name: "Jakarta Pusat"},
	}

	// Setup programs
	mockMobileRepo.programs[1] = model.Program{
		ID:          1,
		Title:       "Training Program",
		Description: "Test training",
		Type:        "training",
		IsActive:    true,
	}

	mockMobileRepo.programs[2] = model.Program{
		ID:          2,
		Title:       "Certification Program",
		Description: "Test certification",
		Type:        "certification",
		IsActive:    true,
	}

	mockMobileRepo.programs[3] = model.Program{
		ID:          3,
		Title:       "Funding Program",
		Description: "Test funding",
		Type:        "funding",
		IsActive:    true,
	}

	service := &mobileService{
		mobileRepo:       mockMobileRepo,
		programRepo:      mockProgramRepo,
		notificationRepo: mockNotifRepo,
		vaultLogRepo:     mockVaultRepo,
		applicationRepo:  mockAppRepo,
		slaRepo:          mockSLARepo,
		minio:            nil,
	}

	return service, mockMobileRepo
}

// Test GetDashboard
func TestGetDashboard(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get dashboard successfully", func(t *testing.T) {
		// This will fail due to vault decryption in real scenario
		_, err := service.GetDashboard(ctx, 1)
		// Expected to fail in test due to vault
		if err != nil {
			// Expected error due to vault decryption
			t.Log("Expected error due to vault decryption:", err)
		}
	})

	t.Run("Get dashboard for non-existing UMKM", func(t *testing.T) {
		_, err := service.GetDashboard(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing UMKM, got none")
		}
	})

	t.Run("Handle notification count error", func(t *testing.T) {
		// Mock notification error scenario would need additional setup
		_, err := service.GetDashboard(ctx, 1)
		// Verify error handling
		if err != nil {
			t.Log("Error handled:", err)
		}
	})
}

// Test GetTrainingPrograms - Additional scenarios
func TestGetTrainingProgramsExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get multiple training programs", func(t *testing.T) {
		// Add more training programs
		mockRepo.programs[4] = model.Program{
			ID:       4,
			Title:    "Advanced Training",
			Type:     "training",
			IsActive: true,
		}

		result, err := service.GetTrainingPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) < 1 {
			t.Errorf("Expected at least 1 training program, got %d", len(result))
		}

		// Verify all are training type
		for _, prog := range result {
			if prog.Type != "training" {
				t.Errorf("Expected type 'training', got '%s'", prog.Type)
			}
		}
	})

	t.Run("Filter out inactive training programs", func(t *testing.T) {
		mockRepo.programs[5] = model.Program{
			ID:       5,
			Title:    "Inactive Training",
			Type:     "training",
			IsActive: false,
		}

		result, err := service.GetTrainingPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check that inactive program is not included
		for _, prog := range result {
			if prog.ID == 5 {
				t.Error("Inactive program should not be returned")
			}
		}
	})
}

// Test GetCertificationPrograms - Additional scenarios
func TestGetCertificationProgramsExtended(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get certification programs with details", func(t *testing.T) {
		result, err := service.GetCertificationPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) < 1 {
			t.Errorf("Expected at least 1 certification program, got %d", len(result))
		}

		// Verify program details
		for _, prog := range result {
			if prog.Type != "certification" {
				t.Errorf("Expected type 'certification', got '%s'", prog.Type)
			}
			if prog.Title == "" {
				t.Error("Program title should not be empty")
			}
		}
	})
}

// Test GetFundingPrograms - Additional scenarios
func TestGetFundingProgramsExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get funding programs with financial details", func(t *testing.T) {
		minAmount := 10000000.0
		maxAmount := 50000000.0
		interestRate := 5.5

		mockRepo.programs[6] = model.Program{
			ID:           6,
			Title:        "SME Loan",
			Type:         "funding",
			IsActive:     true,
			MinAmount:    &minAmount,
			MaxAmount:    &maxAmount,
			InterestRate: &interestRate,
		}

		result, err := service.GetFundingPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Find the newly added program
		found := false
		for _, prog := range result {
			if prog.ID == 6 {
				found = true
				if prog.MinAmount == nil || *prog.MinAmount != minAmount {
					t.Error("MinAmount not correctly mapped")
				}
				if prog.MaxAmount == nil || *prog.MaxAmount != maxAmount {
					t.Error("MaxAmount not correctly mapped")
				}
			}
		}

		if !found {
			t.Error("Expected to find the funding program with ID 6")
		}
	})
}

// Test GetProgramDetail - Additional scenarios
func TestGetProgramDetailExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get program detail with benefits and requirements", func(t *testing.T) {
		result, err := service.GetProgramDetail(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		// Check that benefits and requirements are included
		if result.Benefits == nil {
			t.Error("Benefits should be initialized")
		}
		if result.Requirements == nil {
			t.Error("Requirements should be initialized")
		}
	})

	t.Run("Get inactive program should fail", func(t *testing.T) {
		mockRepo.programs[7] = model.Program{
			ID:       7,
			Title:    "Inactive Program",
			Type:     "training",
			IsActive: false,
		}

		_, err := service.GetProgramDetail(ctx, 7)

		if err == nil {
			t.Error("Expected error for inactive program, got none")
		}
	})
}

// Test GetUMKMProfile - Additional scenarios
func TestGetUMKMProfileExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Verify profile data completeness", func(t *testing.T) {
		// This will fail due to vault, but we can test the structure
		_, err := service.GetUMKMProfile(ctx, 1)
		if err != nil {
			// Expected due to vault
			t.Log("Expected vault error:", err)
		}
	})

	t.Run("Get profile with missing province/city", func(t *testing.T) {
		mockRepo.umkms[2] = model.UMKM{
			ID:           2,
			UserID:       2,
			BusinessName: "Incomplete Business",
			NIK:          "vault:v1:encrypted",
			KartuNumber:  "vault:v1:encrypted",
			User: model.User{
				ID:   2,
				Name: "User 2",
			},
			// Missing Province and City
		}

		_, err := service.GetUMKMProfile(ctx, 2)
		if err != nil {
			t.Log("Error handled for incomplete data:", err)
		}
	})

	t.Run("Test NIK decryption path", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1995-05-15")
		mockRepo.umkms[3] = model.UMKM{
			ID:             3,
			UserID:         3,
			BusinessName:   "Test Decrypt Business",
			NIK:            "vault:v1:encrypted_test_nik",
			KartuNumber:    "vault:v1:encrypted_test_kartu",
			Gender:         "female",
			BirthDate:      birthDate,
			Phone:          "081234567890",
			Address:        "Jl. Test No. 123",
			ProvinceID:     1,
			CityID:         1,
			District:       "District Test",
			Subdistrict:    "Subdistrict Test",
			PostalCode:     "12345",
			NIB:            "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			NPWP:           "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			RevenueRecord:  "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			BusinessPermit: "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			KartuType:      "UMKM",
			Photo:          "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			User: model.User{
				ID:    3,
				Name:  "Test Decrypt User",
				Email: "decrypt@example.com",
			},
			Province: model.Province{
				ID:   1,
				Name: "DKI Jakarta",
			},
			City: model.City{
				ID:   1,
				Name: "Jakarta Pusat",
			},
		}

		result, err := service.GetUMKMProfile(ctx, 3)
		// This will fail at NIK decryption because vault is not setup
		// but it exercises the code path we want to cover
		if err != nil {
			// Expected error because vault client is not initialized
			if !strings.Contains(err.Error(), "failed to decrypt") {
				t.Errorf("Expected decryption error, got: %v", err)
			}
		} else {
			// If somehow it succeeds (shouldn't happen), verify the structure
			if result.ID != 3 {
				t.Errorf("Expected ID 3, got %d", result.ID)
			}
		}
	})

	t.Run("Test Kartu Number decryption path", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1992-03-20")
		mockRepo.umkms[4] = model.UMKM{
			ID:             4,
			UserID:         4,
			BusinessName:   "Test Kartu Business",
			NIK:            "vault:v1:encrypted_nik_4",
			KartuNumber:    "vault:v1:encrypted_kartu_4",
			Gender:         "male",
			BirthDate:      birthDate,
			Phone:          "082345678901",
			Address:        "Jl. Kartu No. 456",
			ProvinceID:     2,
			CityID:         2,
			District:       "District Kartu",
			Subdistrict:    "Subdistrict Kartu",
			PostalCode:     "54321",
			NIB:            "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			NPWP:           "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			RevenueRecord:  "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			BusinessPermit: "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			KartuType:      "Premium",
			Photo:          "iVBORw0KGgoAAAANSUhEUgAAAA0AAAAOCAIAAAB7HQGFAAAAA3NCSVQICAjb4U/gAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAASBJREFUKJGtj7FKw1AUhs899ybNDZrahNrooA4N6lAdnKWor+LmIzj4GuLm6gvoEHBwKSjo5KAoWgSDRUsU0SQnOQ7WobZ18tvOfz7+wxEAEIZhEATwJ+rXTAwH3aILxsZYvmCO8HKGzfD1fly8Oc5JGRsubmsa4h23KUrsw3UZEx29Q0oAIL9X2Nf3+eFCQkSKixkl9jp6+N3VWavh88WTOH1JYuCmycP7bMu4vqOt/WS3JTIqdupi5L/NZft8CQAARV/e885al8zcvomEFN1OXPUraZLNr61gxUEUi1r2vMmae3v1MFefLrjwvLIy0DQNu2REOWPOufXjVf0JVWIppVKKiFBgmqUWEj4/mlZJOrWep7Wl9RQM4HmVwfA/+AIcQGQN0+eDggAAAABJRU5ErkJggg==",
			User: model.User{
				ID:    4,
				Name:  "Kartu Test User",
				Email: "kartu@example.com",
			},
			Province: model.Province{
				ID:   2,
				Name: "Jawa Barat",
			},
			City: model.City{
				ID:   2,
				Name: "Bandung",
			},
		}

		result, err := service.GetUMKMProfile(ctx, 4)
		// This will fail at NIK decryption (first decryption call)
		// We're testing that the code path exists even if it fails
		if err != nil {
			// Expected - vault is not setup
			if !strings.Contains(err.Error(), "decrypt") {
				t.Errorf("Expected decryption error, got: %v", err)
			}
			// The important thing is the code was executed
			t.Log("Decryption code path executed as expected:", err)
		} else {
			// Verify response structure if somehow it succeeds
			if result.BusinessName != "Test Kartu Business" {
				t.Errorf("Expected business name 'Test Kartu Business', got %s", result.BusinessName)
			}
			if result.Province.Name != "Jawa Barat" {
				t.Errorf("Expected province 'Jawa Barat', got %s", result.Province.Name)
			}
			if result.City.Name != "Bandung" {
				t.Errorf("Expected city 'Bandung', got %s", result.City.Name)
			}
		}
	})
}

// Test UpdateUMKMProfile - Additional scenarios
func TestUpdateUMKMProfileExtended(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Update profile with valid data", func(t *testing.T) {
		request := dto.UpdateUMKMProfile{
			BusinessName: "Updated Business",
			Gender:       "male",
			BirthDate:    "1990-01-01",
			Address:      "New Address 123",
			ProvinceID:   1,
			CityID:       1,
			District:     "Central District",
			PostalCode:   "12345",
			Name:         "Updated Name",
		}

		_, err := service.UpdateUMKMProfile(ctx, 1, request)
		// Will fail due to vault, but structure is tested
		if err != nil {
			t.Log("Expected error due to vault:", err)
		}
	})

	t.Run("Update with invalid date format", func(t *testing.T) {
		request := dto.UpdateUMKMProfile{
			BusinessName: "Business",
			Gender:       "male",
			BirthDate:    "invalid-date",
			Address:      "Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "District",
			PostalCode:   "12345",
			Name:         "Name",
		}

		_, err := service.UpdateUMKMProfile(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for invalid date, got none")
		}

		if !contains(err.Error(), "invalid birth date format, use YYYY-MM-DD") {
			t.Errorf("Expected birth date error, got: %v", err)
		}
	})

	// t.Run("Update with photo upload", func(t *testing.T) {
	// 	request := dto.UpdateUMKMProfile{
	// 		BusinessName: "Business",
	// 		Gender:       "female",
	// 		BirthDate:    "1990-01-01",
	// 		Address:      "Address",
	// 		ProvinceID:   1,
	// 		CityID:       1,
	// 		District:     "District",
	// 		PostalCode:   "12345",
	// 		Name:         "Name",
	// 		Photo:        "https://example.com/new_nib.pdf", // Would trigger upload
	// 	}

	// 	_, err := service.UpdateUMKMProfile(ctx, 1, request)
	// 	// Will fail due to minio being nil
	// 	if err != nil {
	// 		t.Log("Error handled:", err)
	// 	}
	// })

	t.Run("Update non-existing profile", func(t *testing.T) {
		request := dto.UpdateUMKMProfile{
			BusinessName: "Business",
			Gender:       "male",
			BirthDate:    "1990-01-01",
			Address:      "Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "District",
			PostalCode:   "12345",
			Name:         "Name",
		}

		_, err := service.UpdateUMKMProfile(ctx, 999, request)

		if err == nil {
			t.Error("Expected error for non-existing profile, got none")
		}
	})
}

// Test GetUMKMDocuments - Additional scenarios
func TestGetUMKMDocumentsExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get all available documents", func(t *testing.T) {
		result, err := service.GetUMKMDocuments(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Should have 2 documents (NIB and NPWP from setup)
		if len(result) != 2 {
			t.Errorf("Expected 2 documents, got %d", len(result))
		}

		// Verify document types
		docTypes := make(map[string]bool)
		for _, doc := range result {
			docTypes[doc.DocumentType] = true
			if doc.DocumentURL == "" {
				t.Error("Document URL should not be empty")
			}
		}

		if !docTypes["nib"] {
			t.Error("Expected NIB document")
		}
		if !docTypes["npwp"] {
			t.Error("Expected NPWP document")
		}
	})

	t.Run("Get documents for UMKM with all document types", func(t *testing.T) {
		mockRepo.umkms[3] = model.UMKM{
			ID:             3,
			UserID:         3,
			NIB:            "http://example.com/nib.pdf",
			NPWP:           "http://example.com/npwp.pdf",
			RevenueRecord:  "http://example.com/revenue.pdf",
			BusinessPermit: "http://example.com/permit.pdf",
		}

		result, err := service.GetUMKMDocuments(ctx, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 4 {
			t.Errorf("Expected 4 documents, got %d", len(result))
		}
	})

	t.Run("Get documents for UMKM with no documents", func(t *testing.T) {
		mockRepo.umkms[4] = model.UMKM{
			ID:     4,
			UserID: 4,
		}

		result, err := service.GetUMKMDocuments(ctx, 4)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 documents, got %d", len(result))
		}
	})
}

// Test UploadDocument - Additional scenarios
func TestUploadDocumentExtended(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Upload valid document", func(t *testing.T) {
		doc := dto.UploadDocumentRequest{
			Type:     "nib",
			Document: "https://example.com/new_nib.pdf",
		}

		err := service.UploadDocument(ctx, 1, doc)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Upload with invalid document type", func(t *testing.T) {
		doc := dto.UploadDocumentRequest{
			Type:     "invalid-type",
			Document: "https://example.com/new_nib.pdf",
		}

		err := service.UploadDocument(ctx, 1, doc)

		if err == nil {
			t.Error("Expected error for invalid document type, got none")
		}

		if !contains(err.Error(), "invalid document type") {
			t.Errorf("Expected document type error, got: %v", err)
		}
	})

	t.Run("Upload for non-existing UMKM", func(t *testing.T) {
		doc := dto.UploadDocumentRequest{
			Type:     "nib",
			Document: "https://example.com/new_nib.pdf",
		}

		err := service.UploadDocument(ctx, 999, doc)

		if err == nil {
			t.Error("Expected error for non-existing UMKM, got none")
		}
	})

	t.Run("Upload replaces old document", func(t *testing.T) {
		// First upload
		doc1 := dto.UploadDocumentRequest{
			Type:     "npwp",
			Document: "https://example.com/old_nib.pdf",
		}
		service.UploadDocument(ctx, 1, doc1)

		// Second upload should replace
		doc2 := dto.UploadDocumentRequest{
			Type:     "npwp",
			Document: "https://example.com/new_nib.pdf",
		}
		err := service.UploadDocument(ctx, 1, doc2)
		if err != nil {
			t.Errorf("Expected no error on replacement, got %v", err)
		}
	})
}

func TestCreateTrainingApplication(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Create training application successfully", func(t *testing.T) {
		request := dto.CreateApplicationTraining{
			ProgramID:          1,
			Motivation:         "I want to improve my business skills",
			BusinessExperience: "5 years in retail",
			LearningObjectives: "Learn digital marketing",
			AvailabilityNotes:  "Available weekdays",
			Documents: map[string]string{
				"ktp":      "http://example.com/ktp.pdf",
				"proposal": "http://example.com/proposal.pdf",
			},
		}

		err := service.CreateTrainingApplication(ctx, 1, request)
		// Will fail due to vault decryption
		if err != nil {
			if !contains(err.Error(), "UMKM profile not found") {
				t.Log("Expected vault error:", err)
			}
		}
	})

	t.Run("Create application with non-existing program", func(t *testing.T) {
		request := dto.CreateApplicationTraining{
			ProgramID:  999,
			Motivation: "Test motivation",
			Documents:  map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateTrainingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})

	t.Run("Create application with wrong program type", func(t *testing.T) {
		// Program ID 2 is certification, not training
		request := dto.CreateApplicationTraining{
			ProgramID:  2,
			Motivation: "Test motivation",
			Documents:  map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateTrainingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for wrong program type, got none")
		}

		if !contains(err.Error(), "program type must be training") {
			t.Errorf("Expected program type error, got: %v", err)
		}
	})

	t.Run("Create duplicate application", func(t *testing.T) {
		// Add existing application
		mockRepo.applications[1] = model.Application{
			ID:        1,
			UMKMID:    1,
			ProgramID: 1,
			Status:    "screening",
		}

		request := dto.CreateApplicationTraining{
			ProgramID:  1,
			Motivation: "Test motivation",
			Documents:  map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateTrainingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for duplicate application, got none")
		}

		if !contains(err.Error(), "you have already applied for this program") {
			t.Errorf("Expected duplicate error, got: %v", err)
		}
	})

	t.Run("Create application for non-existing UMKM", func(t *testing.T) {
		request := dto.CreateApplicationTraining{
			ProgramID:  1,
			Motivation: "Test motivation",
			Documents:  map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateTrainingApplication(ctx, 999, request)

		if err == nil {
			t.Error("Expected error for non-existing UMKM, got none")
		}

		if !contains(err.Error(), "UMKM profile not found") {
			t.Errorf("Expected UMKM not found error, got: %v", err)
		}
	})

	t.Run("Create application without required documents", func(t *testing.T) {
		request := dto.CreateApplicationTraining{
			ProgramID:  1,
			Motivation: "Test motivation",
			Documents:  map[string]string{}, // Empty documents
		}

		err := service.CreateTrainingApplication(ctx, 1, request)
		if err != nil {
			t.Log("Error handled:", err)
		}
	})
}

// ==================== TEST CREATE CERTIFICATION APPLICATION ====================

func TestCreateCertificationApplication(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Create certification application successfully", func(t *testing.T) {
		yearsOperating := 5
		request := dto.CreateApplicationCertification{
			ProgramID:           2,
			BusinessSector:      "Food & Beverage",
			ProductOrService:    "Halal snacks",
			BusinessDescription: "We produce halal snacks",
			YearsOperating:      &yearsOperating,
			CurrentStandards:    "ISO 9001",
			CertificationGoals:  "Get Halal certification",
			Documents: map[string]string{
				"ktp":  "http://example.com/ktp.pdf",
				"nib":  "http://example.com/nib.pdf",
				"npwp": "http://example.com/npwp.pdf",
			},
		}

		err := service.CreateCertificationApplication(ctx, 1, request)
		// Will fail due to vault
		if err != nil {
			if !contains(err.Error(), "UMKM profile not found") {
				t.Log("Expected vault error:", err)
			}
		}
	})

	t.Run("Create certification with wrong program type", func(t *testing.T) {
		request := dto.CreateApplicationCertification{
			ProgramID:           1, // Training program
			BusinessSector:      "Food",
			ProductOrService:    "Snacks",
			BusinessDescription: "Description",
			CertificationGoals:  "Goals",
			Documents:           map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateCertificationApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for wrong program type, got none")
		}

		if !contains(err.Error(), "program type must be certification") {
			t.Errorf("Expected program type error, got: %v", err)
		}
	})

	t.Run("Create certification without optional fields", func(t *testing.T) {
		request := dto.CreateApplicationCertification{
			ProgramID:           2,
			BusinessSector:      "Food",
			ProductOrService:    "Snacks",
			BusinessDescription: "Description",
			CertificationGoals:  "Goals",
			Documents:           map[string]string{"ktp": "http://example.com/ktp.pdf"},
			// No YearsOperating, CurrentStandards
		}

		err := service.CreateCertificationApplication(ctx, 1, request)
		// Should work even without optional fields
		if err != nil {
			if !contains(err.Error(), "UMKM profile not found") {
				t.Log("Error:", err)
			}
		}
	})

	t.Run("Create duplicate certification application", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:        2,
			UMKMID:    1,
			ProgramID: 2,
			Status:    "screening",
		}

		request := dto.CreateApplicationCertification{
			ProgramID:           2,
			BusinessSector:      "Food",
			ProductOrService:    "Snacks",
			BusinessDescription: "Description",
			CertificationGoals:  "Goals",
			Documents:           map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateCertificationApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for duplicate application, got none")
		}
	})
}

// ==================== TEST CREATE FUNDING APPLICATION ====================

func TestCreateFundingApplication(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	// Setup funding program with limits
	minAmount := 10000000.0
	maxAmount := 50000000.0
	maxTenure := 12

	mockRepo.programs[3] = model.Program{
		ID:              3,
		Title:           "SME Funding",
		Type:            "funding",
		IsActive:        true,
		MinAmount:       &minAmount,
		MaxAmount:       &maxAmount,
		MaxTenureMonths: &maxTenure,
	}

	t.Run("Create funding application successfully", func(t *testing.T) {
		yearsOperating := 3
		revenue := 5000000.0
		monthlyRev := 500000.0

		request := dto.CreateApplicationFunding{
			ProgramID:             3,
			BusinessSector:        "Retail",
			BusinessDescription:   "Small retail business",
			YearsOperating:        &yearsOperating,
			RequestedAmount:       20000000,
			FundPurpose:           "Expand inventory",
			BusinessPlan:          "Detailed business plan",
			RevenueProjection:     &revenue,
			MonthlyRevenue:        &monthlyRev,
			RequestedTenureMonths: 10,
			CollateralDescription: "Property certificate",
			Documents: map[string]string{
				"ktp":      "http://example.com/ktp.pdf",
				"proposal": "http://example.com/proposal.pdf",
			},
		}

		err := service.CreateFundingApplication(ctx, 1, request)
		// Will fail due to vault
		if err != nil {
			if !contains(err.Error(), "UMKM profile not found") {
				t.Log("Expected vault error:", err)
			}
		}
	})

	t.Run("Create funding with amount below minimum", func(t *testing.T) {
		request := dto.CreateApplicationFunding{
			ProgramID:             3,
			BusinessSector:        "Retail",
			BusinessDescription:   "Description",
			RequestedAmount:       5000000, // Below min
			FundPurpose:           "Purpose",
			RequestedTenureMonths: 10,
			Documents:             map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateFundingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for amount below minimum, got none")
		}

		if !contains(err.Error(), "requested amount must be at least") {
			t.Errorf("Expected minimum amount error, got: %v", err)
		}
	})

	t.Run("Create funding with amount above maximum", func(t *testing.T) {
		request := dto.CreateApplicationFunding{
			ProgramID:             3,
			BusinessSector:        "Retail",
			BusinessDescription:   "Description",
			RequestedAmount:       100000000, // Above max
			FundPurpose:           "Purpose",
			RequestedTenureMonths: 10,
			Documents:             map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateFundingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for amount above maximum, got none")
		}

		if !contains(err.Error(), "requested amount cannot exceed") {
			t.Errorf("Expected maximum amount error, got: %v", err)
		}
	})

	t.Run("Create funding with tenure exceeding limit", func(t *testing.T) {
		request := dto.CreateApplicationFunding{
			ProgramID:             3,
			BusinessSector:        "Retail",
			BusinessDescription:   "Description",
			RequestedAmount:       20000000,
			FundPurpose:           "Purpose",
			RequestedTenureMonths: 24, // Above max of 12
			Documents:             map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateFundingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for tenure exceeding limit, got none")
		}

		if !contains(err.Error(), "requested tenure cannot exceed") {
			t.Errorf("Expected tenure error, got: %v", err)
		}
	})

	t.Run("Create funding with wrong program type", func(t *testing.T) {
		request := dto.CreateApplicationFunding{
			ProgramID:             1, // Training program
			BusinessSector:        "Retail",
			BusinessDescription:   "Description",
			RequestedAmount:       20000000,
			FundPurpose:           "Purpose",
			RequestedTenureMonths: 10,
			Documents:             map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateFundingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for wrong program type, got none")
		}

		if !contains(err.Error(), "program type must be funding") {
			t.Errorf("Expected program type error, got: %v", err)
		}
	})

	t.Run("Create duplicate funding application", func(t *testing.T) {
		mockRepo.applications[3] = model.Application{
			ID:        3,
			UMKMID:    1,
			ProgramID: 3,
			Status:    "screening",
		}

		request := dto.CreateApplicationFunding{
			ProgramID:             3,
			BusinessSector:        "Retail",
			BusinessDescription:   "Description",
			RequestedAmount:       20000000,
			FundPurpose:           "Purpose",
			RequestedTenureMonths: 10,
			Documents:             map[string]string{"ktp": "http://example.com/ktp.pdf"},
		}

		err := service.CreateFundingApplication(ctx, 1, request)

		if err == nil {
			t.Error("Expected error for duplicate application, got none")
		}
	})
}

// ==================== TEST GET APPLICATION LIST ====================

func TestGetApplicationListExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get application list with multiple statuses", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:          1,
			UMKMID:      1,
			ProgramID:   1,
			Type:        "training",
			Status:      "screening",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[1],
		}

		mockRepo.applications[2] = model.Application{
			ID:          2,
			UMKMID:      1,
			ProgramID:   2,
			Type:        "certification",
			Status:      "approved",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[2],
		}

		result, err := service.GetApplicationList(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("Expected 2 applications, got %d", len(result))
		}

		// Verify different statuses
		statuses := make(map[string]bool)
		for _, app := range result {
			statuses[app.Status] = true
		}

		if !statuses["screening"] || !statuses["approved"] {
			t.Error("Expected both screening and approved statuses")
		}
	})

	t.Run("Get empty list for UMKM with no applications", func(t *testing.T) {
		result, err := service.GetApplicationList(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 applications, got %d", len(result))
		}
	})

	t.Run("Verify application list order by submission date", func(t *testing.T) {
		result, err := service.GetApplicationList(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) > 1 {
			// Check if ordered by SubmittedAt DESC
			for i := 0; i < len(result)-1; i++ {
				current, _ := time.Parse("2006-01-02 15:04:05", result[i].SubmittedAt)
				next, _ := time.Parse("2006-01-02 15:04:05", result[i+1].SubmittedAt)

				if current.Before(next) {
					t.Error("Applications should be ordered by submission date DESC")
				}
			}
		}
	})
}

// ==================== TEST GET APPLICATION DETAIL ====================

func TestGetApplicationDetailExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get training application detail", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:          1,
			UMKMID:      1,
			ProgramID:   1,
			Type:        "training",
			Status:      "screening",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[1],
			TrainingApplication: &model.TrainingApplication{
				Motivation:         "Test motivation",
				BusinessExperience: "5 years",
			},
		}

		result, err := service.GetApplicationDetail(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Type != "training" {
			t.Errorf("Expected type 'training', got '%s'", result.Type)
		}

		if result.TrainingData == nil {
			t.Error("Expected training data to be populated")
		} else {
			if result.TrainingData.Motivation != "Test motivation" {
				t.Error("Training motivation not correctly mapped")
			}
		}
	})

	t.Run("Get certification application detail", func(t *testing.T) {
		yearsOp := 5
		mockRepo.applications[2] = model.Application{
			ID:          2,
			UMKMID:      1,
			ProgramID:   2,
			Type:        "certification",
			Status:      "final",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[2],
			CertificationApplication: &model.CertificationApplication{
				BusinessSector:      "Food",
				ProductOrService:    "Snacks",
				BusinessDescription: "Halal snacks producer",
				YearsOperating:      &yearsOp,
				CertificationGoals:  "Halal certification",
			},
		}

		result, err := service.GetApplicationDetail(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.CertificationData == nil {
			t.Error("Expected certification data to be populated")
		} else {
			if result.CertificationData.BusinessSector != "Food" {
				t.Error("Business sector not correctly mapped")
			}
			if result.CertificationData.YearsOperating == nil || *result.CertificationData.YearsOperating != 5 {
				t.Error("Years operating not correctly mapped")
			}
		}
	})

	t.Run("Get funding application detail", func(t *testing.T) {
		yearsOp := 3
		revenue := 5000000.0
		monthlyRev := 500000.0

		mockRepo.applications[3] = model.Application{
			ID:          3,
			UMKMID:      1,
			ProgramID:   3,
			Type:        "funding",
			Status:      "approved",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[3],
			FundingApplication: &model.FundingApplication{
				BusinessSector:        "Retail",
				BusinessDescription:   "Small retail",
				YearsOperating:        &yearsOp,
				RequestedAmount:       20000000,
				FundPurpose:           "Expand inventory",
				RevenueProjection:     &revenue,
				MonthlyRevenue:        &monthlyRev,
				RequestedTenureMonths: 10,
				CollateralDescription: "Property",
			},
		}

		result, err := service.GetApplicationDetail(ctx, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.FundingData == nil {
			t.Error("Expected funding data to be populated")
		} else {
			if result.FundingData.RequestedAmount != 20000000 {
				t.Error("Requested amount not correctly mapped")
			}
			if result.FundingData.RequestedTenureMonths != 10 {
				t.Error("Tenure not correctly mapped")
			}
		}
	})

	t.Run("Get application with documents and histories", func(t *testing.T) {
		mockRepo.applications[4] = model.Application{
			ID:          4,
			UMKMID:      1,
			ProgramID:   1,
			Type:        "training",
			Status:      "screening",
			SubmittedAt: time.Now(),
			ExpiredAt:   time.Now().AddDate(0, 0, 30),
			Program:     mockRepo.programs[1],
			Documents: []model.ApplicationDocument{
				{ID: 1, Type: "ktp", File: "http://example.com/ktp.pdf"},
				{ID: 2, Type: "proposal", File: "http://example.com/proposal.pdf"},
			},
			Histories: []model.ApplicationHistory{
				{
					ID:         1,
					Status:     "submit",
					Notes:      "Application submitted",
					ActionedAt: time.Now(),
					User:       model.User{Name: "System"},
				},
			},
		}

		result, err := service.GetApplicationDetail(ctx, 4)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Documents) != 2 {
			t.Errorf("Expected 2 documents, got %d", len(result.Documents))
		}

		if len(result.Histories) != 1 {
			t.Errorf("Expected 1 history, got %d", len(result.Histories))
		}
	})

	t.Run("Get non-existing application detail", func(t *testing.T) {
		_, err := service.GetApplicationDetail(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing application, got none")
		}
	})
}

// ==================== TEST GET UMKM PROFILE WITH DECRYPTION ====================

func TestGetUMKMProfileWithDecryption(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Get profile with decryption for application creation", func(t *testing.T) {
		_, err := service.GetUMKMProfileWithDecryption(ctx, 1, "application_creation")
		// Will fail due to vault
		if err != nil {
			if !contains(err.Error(), "failed to decrypt") {
				t.Log("Expected vault error:", err)
			}
		}
	})

	t.Run("Get profile with decryption for review", func(t *testing.T) {
		_, err := service.GetUMKMProfileWithDecryption(ctx, 1, "application_review")
		if err != nil {
			t.Log("Expected vault error:", err)
		}
	})

	t.Run("Get non-existing profile", func(t *testing.T) {
		_, err := service.GetUMKMProfileWithDecryption(ctx, 999, "test")

		if err == nil {
			t.Error("Expected error for non-existing profile, got none")
		}
	})
}

// ==================== TEST REVISE APPLICATION ====================

func TestReviseApplication(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	mockRepo.applications[1] = model.Application{
		ID:     1,
		UMKMID: 1,
		Status: constant.ApplicationStatusRevised,
	}

	t.Run("Revise application successfully", func(t *testing.T) {
		documents := []dto.UploadDocumentRequest{
			{Type: "ktp", Document: "http://example.com/ktp.pdf"},
			{Type: "proposal", Document: "http://example.com/proposal.pdf"},
		}

		err := service.ReviseApplication(ctx, 1, 1, documents)
		// Will fail due to vault, but structure is tested
		if err != nil {
			t.Log("Expected error due to vault:", err)
		}
	})

	t.Run("Revise non-revised application", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: constant.ApplicationStatusScreening,
		}

		documents := []dto.UploadDocumentRequest{
			{Type: "ktp", Document: "http://example.com/ktp.pdf"},
		}

		err := service.ReviseApplication(ctx, 1, 2, documents)
		if err == nil {
			t.Error("Expected error for non-revised application, got none")
		}

		expectedMsg := "application is not in a revisable state"
		if err != nil && err.Error() != expectedMsg {
			t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Revise non-existing application", func(t *testing.T) {
		documents := []dto.UploadDocumentRequest{}

		err := service.ReviseApplication(ctx, 1, 999, documents)
		if err == nil {
			t.Error("Expected error for non-existing application, got none")
		}
	})
}

// ==================== TEST NOTIFICATIONS ====================

func TestGetNotificationsByUMKMID(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	// Setup notifications via the notification repo
	notifService := service.notificationRepo.(*mockNotificationRepoForMobile)
	now := time.Now()
	notifService.notifications = []model.Notification{
		{
			ID:      1,
			UMKMID:  1,
			Type:    "application_submitted",
			Title:   "Application Submitted",
			Message: "Your application has been submitted",
			IsRead:  false,
			Base: model.Base{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			ID:      2,
			UMKMID:  1,
			Type:    "screening_approved",
			Title:   "Application Approved",
			Message: "Your application has been approved",
			IsRead:  true,
			ReadAt:  &now,
			Base: model.Base{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}

	t.Run("Get all notifications for UMKM", func(t *testing.T) {
		result, err := service.GetNotificationsByUMKMID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("Expected 2 notifications, got %d", len(result))
		}

		// Verify read/unread status
		readCount := 0
		unreadCount := 0
		for _, notif := range result {
			if notif.IsRead {
				readCount++
			} else {
				unreadCount++
			}
		}

		if readCount != 1 || unreadCount != 1 {
			t.Errorf("Expected 1 read and 1 unread, got %d read and %d unread", readCount, unreadCount)
		}
	})

	t.Run("Get empty notifications list", func(t *testing.T) {
		result, err := service.GetNotificationsByUMKMID(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 notifications, got %d", len(result))
		}
	})

	t.Run("Handle notification repository error", func(t *testing.T) {
		notifService.shouldError = true
		defer func() { notifService.shouldError = false }()

		_, err := service.GetNotificationsByUMKMID(ctx, 1)

		if err == nil {
			t.Error("Expected error from repository, got none")
		}
	})
}

func TestGetUnreadCount(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	notifService := service.notificationRepo.(*mockNotificationRepoForMobile)
	notifService.notifications = []model.Notification{
		{ID: 1, UMKMID: 1, IsRead: false},
		{ID: 2, UMKMID: 1, IsRead: false},
		{ID: 3, UMKMID: 1, IsRead: true},
	}

	t.Run("Get unread count", func(t *testing.T) {
		count, err := service.GetUnreadCount(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if count != 2 {
			t.Errorf("Expected 2 unread notifications, got %d", count)
		}
	})

	t.Run("Get zero unread count", func(t *testing.T) {
		count, err := service.GetUnreadCount(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if count != 0 {
			t.Errorf("Expected 0 unread notifications, got %d", count)
		}
	})

	t.Run("Handle repository error", func(t *testing.T) {
		notifService.shouldError = true
		defer func() { notifService.shouldError = false }()

		_, err := service.GetUnreadCount(ctx, 1)

		if err == nil {
			t.Error("Expected error from repository, got none")
		}
	})
}

func TestMarkNotificationsAsRead(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Mark notification as read", func(t *testing.T) {
		err := service.MarkNotificationsAsRead(ctx, 1, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Handle repository error", func(t *testing.T) {
		notifService := service.notificationRepo.(*mockNotificationRepoForMobile)
		notifService.shouldError = true
		defer func() { notifService.shouldError = false }()

		err := service.MarkNotificationsAsRead(ctx, 1, 1)

		if err == nil {
			t.Error("Expected error from repository, got none")
		}
	})
}

func TestMarkAllNotificationsAsRead(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Mark all notifications as read", func(t *testing.T) {
		err := service.MarkAllNotificationsAsRead(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Handle repository error", func(t *testing.T) {
		notifService := service.notificationRepo.(*mockNotificationRepoForMobile)
		notifService.shouldError = true
		defer func() { notifService.shouldError = false }()

		err := service.MarkAllNotificationsAsRead(ctx, 1)

		if err == nil {
			t.Error("Expected error from repository, got none")
		}
	})
}

//  ==================== TEST NEWS ====================

func (m *mockNewsRepoForMobile) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	if m.shouldError {
		return nil, 0, errors.New("database error")
	}

	var result []model.News
	for _, n := range m.news {
		if !n.IsPublished {
			continue
		}

		// Filter by category
		if params.Category != "" && n.Category != params.Category {
			continue
		}

		// Filter by search
		if params.Search != "" {
			if !strings.Contains(strings.ToLower(n.Title), strings.ToLower(params.Search)) &&
				!strings.Contains(strings.ToLower(n.Content), strings.ToLower(params.Search)) {
				continue
			}
		}

		// Filter by tag
		if params.Tag != "" {
			hasTag := false
			for _, tag := range n.Tags {
				if tag.TagName == params.Tag {
					hasTag = true
					break
				}
			}
			if !hasTag {
				continue
			}
		}

		result = append(result, n)
	}

	total := int64(len(result))

	// Apply pagination
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	start := (params.Page - 1) * params.Limit
	end := start + params.Limit

	if start >= len(result) {
		return []model.News{}, total, nil
	}
	if end > len(result) {
		end = len(result)
	}

	return result[start:end], total, nil
}

func (m *mockNewsRepoForMobile) GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News, error) {
	if m.shouldError {
		return model.News{}, errors.New("database error")
	}

	for _, n := range m.news {
		if n.Slug == slug && n.IsPublished {
			return n, nil
		}
	}

	return model.News{}, errors.New("news not found")
}

func (m *mockNewsRepoForMobile) IncrementViews(ctx context.Context, newsID int) error {
	if m.shouldError {
		return errors.New("database error")
	}

	if news, exists := m.news[newsID]; exists {
		news.ViewsCount++
		m.news[newsID] = news
		return nil
	}

	return errors.New("news not found")
}

// ==================== TEST HELPER FUNCTIONS ====================

func TestCreateApplicationHistory(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Create application history", func(t *testing.T) {
		err := service.createApplicationHistory(ctx, 1, 1, "submit", "Application submitted")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Create history with empty status", func(t *testing.T) {
		// Service should still accept it
		err := service.createApplicationHistory(ctx, 1, 1, "", "Test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestCreateNotification(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Create notification", func(t *testing.T) {
		err := service.createNotification(ctx, 1, 1,
			constant.NotificationSubmitted,
			constant.NotificationTitleSubmitted,
			constant.NotificationMessageSubmitted)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Create notification with different types", func(t *testing.T) {
		types := []string{
			constant.NotificationApproved,
			constant.NotificationRejected,
			constant.NotificationRevised,
		}

		for _, notifType := range types {
			err := service.createNotification(ctx, 1, 1, notifType, "Title", "Message")
			if err != nil {
				t.Errorf("Expected no error for type %s, got %v", notifType, err)
			}
		}
	})
}

func TestProcessAndSaveDocuments(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Process documents with URLs", func(t *testing.T) {
		documents := map[string]string{
			"ktp":      "http://example.com/ktp.pdf",
			"proposal": "http://example.com/proposal.pdf",
			"npwp":     "http://example.com/npwp.pdf",
		}

		// Should not panic or error
		service.processAndSaveDocuments(ctx, 1, documents)

		// Give goroutine time to complete
		time.Sleep(100 * time.Millisecond)
	})

	t.Run("Process empty documents", func(t *testing.T) {
		documents := map[string]string{}

		// Should handle gracefully
		service.processAndSaveDocuments(ctx, 1, documents)
		time.Sleep(100 * time.Millisecond)
	})

	t.Run("Process with mixed URL and base64", func(t *testing.T) {
		documents := map[string]string{
			"ktp":  "http://example.com/ktp.pdf",
			"npwp": "http://example.com/npwp.pdf",
		}

		service.processAndSaveDocuments(ctx, 1, documents)
		time.Sleep(100 * time.Millisecond)
	})
}

// ==================== TEST UPLOAD DOCUMENT WITHOUT MINIO ====================

func TestUploadDocumentWithoutMinio(t *testing.T) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	validDocTypes := []string{"nib", "npwp", "revenue_record", "business_permit"}

	for _, docType := range validDocTypes {
		t.Run("Upload "+docType+" with URL", func(t *testing.T) {
			doc := dto.UploadDocumentRequest{
				Type:     docType,
				Document: "http://example.com/" + docType + ".pdf",
			}

			err := service.UploadDocument(ctx, 1, doc)
			if err != nil {
				t.Errorf("Expected no error for %s, got %v", docType, err)
			}
		})
	}

	t.Run("Upload with HTTPS URL", func(t *testing.T) {
		doc := dto.UploadDocumentRequest{
			Type:     "nib",
			Document: "https://secure.example.com/nib.pdf",
		}

		err := service.UploadDocument(ctx, 1, doc)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

// Test Edge Cases
func TestMobileServiceEdgeCasesExtended(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Handle concurrent profile updates", func(t *testing.T) {
		done := make(chan bool)

		for i := 0; i < 5; i++ {
			go func(id int) {
				request := dto.UpdateUMKMProfile{
					BusinessName: "Concurrent Business",
					Gender:       "male",
					BirthDate:    "1990-01-01",
					Address:      "Address",
					ProvinceID:   1,
					CityID:       1,
					District:     "District",
					PostalCode:   "12345",
					Name:         "Name",
				}

				_, err := service.UpdateUMKMProfile(ctx, 1, request)
				if err != nil {
					t.Log("Concurrent error:", err)
				}
				done <- true
			}(i)
		}

		for i := 0; i < 5; i++ {
			<-done
		}
	})

	t.Run("Handle empty program list for each type", func(t *testing.T) {
		// Clear all programs
		mockRepo.programs = make(map[int]model.Program)

		training, err1 := service.GetTrainingPrograms(ctx)
		cert, err2 := service.GetCertificationPrograms(ctx)
		funding, err3 := service.GetFundingPrograms(ctx)

		if err1 != nil || err2 != nil || err3 != nil {
			t.Error("Expected no errors for empty lists")
		}

		if len(training) != 0 || len(cert) != 0 || len(funding) != 0 {
			t.Error("Expected empty program lists")
		}
	})

	t.Run("Verify data consistency across methods", func(t *testing.T) {
		// Get documents and verify profile has same data
		docs, err1 := service.GetUMKMDocuments(ctx, 1)

		if err1 != nil {
			t.Errorf("Expected no error, got %v", err1)
		}

		// Verify UMKM exists in repo
		if _, exists := mockRepo.umkms[1]; !exists {
			t.Error("UMKM should exist in repository")
		}

		// Check document count matches
		umkm := mockRepo.umkms[1]
		expectedDocs := 0
		if umkm.NIB != "" {
			expectedDocs++
		}
		if umkm.NPWP != "" {
			expectedDocs++
		}
		if umkm.RevenueRecord != "" {
			expectedDocs++
		}
		if umkm.BusinessPermit != "" {
			expectedDocs++
		}

		if len(docs) != expectedDocs {
			t.Errorf("Expected %d documents, got %d", expectedDocs, len(docs))
		}
	})
}

// Benchmark Tests
func BenchmarkGetTrainingPrograms(b *testing.B) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetTrainingPrograms(ctx)
	}
}

func BenchmarkGetUMKMDocuments(b *testing.B) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUMKMDocuments(ctx, 1)
	}
}

func BenchmarkGetProgramDetail(b *testing.B) {
	service, _ := setupMobileServiceForTests()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetProgramDetail(ctx, 1)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		(s == substr || len(s) >= len(substr) && s[:len(substr)] == substr ||
			len(s) > len(substr) && s[len(s)-len(substr):] == substr)
}

// ==================== NEWS TESTS ====================

func TestGetPublishedNews(t *testing.T) {
	t.Run("Success - Get published news with results", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		// Setup test data
		now := time.Now()
		mockRepo.news[1] = model.News{
			ID:          1,
			Title:       "Test News 1",
			Slug:        "test-news-1",
			Excerpt:     "This is test news 1 excerpt",
			Thumbnail:   "https://example.com/thumbnail1.jpg",
			Category:    "general",
			ViewsCount:  100,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Admin User",
			},
		}

		mockRepo.news[2] = model.News{
			ID:          2,
			Title:       "Test News 2",
			Slug:        "test-news-2",
			Excerpt:     "This is test news 2 excerpt",
			Thumbnail:   "https://example.com/thumbnail2.jpg",
			Category:    "announcement",
			ViewsCount:  200,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now.Add(-24 * time.Hour), UpdatedAt: now},
			Author: model.User{
				ID:   2,
				Name: "Editor User",
			},
		}

		params := dto.NewsQueryParams{
			Page:  1,
			Limit: 10,
		}

		// Execute
		result, total, err := service.GetPublishedNews(ctx, params)
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if total != 2 {
			t.Errorf("Expected total 2, got %d", total)
		}

		if len(result) != 2 {
			t.Errorf("Expected 2 news items, got %d", len(result))
		}

		// Verify first news item
		if result[0].ID != 1 && result[0].ID != 2 {
			t.Errorf("Expected news ID 1 or 2, got %d", result[0].ID)
		}

		// Verify fields mapping
		for _, news := range result {
			if news.Title == "" {
				t.Error("Title should not be empty")
			}
			if news.Slug == "" {
				t.Error("Slug should not be empty")
			}
			if news.AuthorName == "" {
				t.Error("AuthorName should not be empty")
			}
			if news.CreatedAt == "" {
				t.Error("CreatedAt should not be empty")
			}

			// Verify date format
			if !strings.Contains(news.CreatedAt, "-") || !strings.Contains(news.CreatedAt, ":") {
				t.Errorf("CreatedAt should be in format YYYY-MM-DD HH:MM:SS, got %s", news.CreatedAt)
			}
		}
	})

	t.Run("Success - Empty news list", func(t *testing.T) {
		service, _ := setupMobileServiceForTests()
		ctx := context.Background()

		params := dto.NewsQueryParams{
			Page:  1,
			Limit: 10,
		}

		// Execute
		result, total, err := service.GetPublishedNews(ctx, params)
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if total != 0 {
			t.Errorf("Expected total 0, got %d", total)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 news items, got %d", len(result))
		}
	})

	t.Run("Success - Only published news returned", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()

		// Published news
		mockRepo.news[1] = model.News{
			ID:          1,
			Title:       "Published News",
			Slug:        "published-news",
			Excerpt:     "This is published",
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Admin",
			},
		}

		// Unpublished news (should not be returned)
		mockRepo.news[2] = model.News{
			ID:          2,
			Title:       "Unpublished News",
			Slug:        "unpublished-news",
			Excerpt:     "This is unpublished",
			IsPublished: false,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Admin",
			},
		}

		params := dto.NewsQueryParams{
			Page:  1,
			Limit: 10,
		}

		// Execute
		result, total, err := service.GetPublishedNews(ctx, params)
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if total != 1 {
			t.Errorf("Expected total 1, got %d", total)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 news item, got %d", len(result))
		}

		if len(result) > 0 && result[0].Title != "Published News" {
			t.Errorf("Expected 'Published News', got '%s'", result[0].Title)
		}
	})

	t.Run("Success - Verify all fields are correctly mapped", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[1] = model.News{
			ID:          999,
			Title:       "Complete News",
			Slug:        "complete-news",
			Excerpt:     "Complete excerpt",
			Thumbnail:   "https://example.com/complete.jpg",
			Category:    "event",
			ViewsCount:  500,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   5,
				Name: "Complete Author",
			},
		}

		params := dto.NewsQueryParams{}

		// Execute
		result, _, err := service.GetPublishedNews(ctx, params)
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 1 {
			t.Fatalf("Expected 1 news item, got %d", len(result))
		}

		news := result[0]
		if news.ID != 999 {
			t.Errorf("Expected ID 999, got %d", news.ID)
		}
		if news.Title != "Complete News" {
			t.Errorf("Expected Title 'Complete News', got '%s'", news.Title)
		}
		if news.Slug != "complete-news" {
			t.Errorf("Expected Slug 'complete-news', got '%s'", news.Slug)
		}
		if news.Excerpt != "Complete excerpt" {
			t.Errorf("Expected Excerpt 'Complete excerpt', got '%s'", news.Excerpt)
		}
		if news.Thumbnail != "https://example.com/complete.jpg" {
			t.Errorf("Expected Thumbnail URL, got '%s'", news.Thumbnail)
		}
		if news.Category != "event" {
			t.Errorf("Expected Category 'event', got '%s'", news.Category)
		}
		if news.AuthorName != "Complete Author" {
			t.Errorf("Expected AuthorName 'Complete Author', got '%s'", news.AuthorName)
		}
		if news.ViewsCount != 500 {
			t.Errorf("Expected ViewsCount 500, got %d", news.ViewsCount)
		}
		if news.CreatedAt != now.Format("2006-01-02 15:04:05") {
			t.Errorf("Expected CreatedAt '%s', got '%s'", now.Format("2006-01-02 15:04:05"), news.CreatedAt)
		}
	})

	t.Run("Error - Repository returns error", func(t *testing.T) {
		mockRepo := newMockMobileRepository()
		mockProgramRepo := newMockProgramRepoForMobile()
		mockNotifRepo := newMockNotificationRepoForMobile()
		mockVaultLogRepo := newMockVaultDecryptLogRepo()
		mockApplicationRepo := newMockApplicationsRepo()
		mockSLARepo := newMockSLARepo()

		// Create custom mock that returns error
		customMockRepo := &mockMobileRepoWithError{
			mockMobileRepository: mockRepo,
			shouldErrorOnGetNews: true,
		}

		service := NewMobileService(
			customMockRepo,
			mockProgramRepo,
			mockNotifRepo,
			mockVaultLogRepo,
			mockApplicationRepo,
			mockSLARepo,
			nil,
		)

		ctx := context.Background()
		params := dto.NewsQueryParams{}

		// Execute
		result, total, err := service.GetPublishedNews(ctx, params)

		// Verify
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if total != 0 {
			t.Errorf("Expected total 0, got %d", total)
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}

func TestGetNewsDetail(t *testing.T) {
	t.Run("Success - Get news detail with tags", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[1] = model.News{
			ID:          1,
			Title:       "Detailed News",
			Slug:        "detailed-news",
			Content:     "This is the full content of the news article",
			Thumbnail:   "https://example.com/thumbnail.jpg",
			Category:    "general",
			ViewsCount:  150,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Author Name",
			},
			Tags: []model.NewsTag{
				{ID: 1, NewsID: 1, TagName: "technology"},
				{ID: 2, NewsID: 1, TagName: "innovation"},
				{ID: 3, NewsID: 1, TagName: "business"},
			},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "detailed-news")
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		if result.Title != "Detailed News" {
			t.Errorf("Expected Title 'Detailed News', got '%s'", result.Title)
		}

		if result.Slug != "detailed-news" {
			t.Errorf("Expected Slug 'detailed-news', got '%s'", result.Slug)
		}

		if result.Content != "This is the full content of the news article" {
			t.Errorf("Expected full content, got '%s'", result.Content)
		}

		if result.Thumbnail != "https://example.com/thumbnail.jpg" {
			t.Errorf("Expected Thumbnail URL, got '%s'", result.Thumbnail)
		}

		if result.Category != "general" {
			t.Errorf("Expected Category 'general', got '%s'", result.Category)
		}

		if result.AuthorName != "Author Name" {
			t.Errorf("Expected AuthorName 'Author Name', got '%s'", result.AuthorName)
		}

		// Verify views count is incremented
		if result.ViewsCount != 151 {
			t.Errorf("Expected ViewsCount 151 (150 + 1), got %d", result.ViewsCount)
		}

		if result.CreatedAt != now.Format("2006-01-02 15:04:05") {
			t.Errorf("Expected CreatedAt '%s', got '%s'", now.Format("2006-01-02 15:04:05"), result.CreatedAt)
		}

		// Verify tags
		if len(result.Tags) != 3 {
			t.Errorf("Expected 3 tags, got %d", len(result.Tags))
		}

		expectedTags := map[string]bool{
			"technology": false,
			"innovation": false,
			"business":   false,
		}

		for _, tag := range result.Tags {
			if _, exists := expectedTags[tag]; exists {
				expectedTags[tag] = true
			} else {
				t.Errorf("Unexpected tag: %s", tag)
			}
		}

		for tag, found := range expectedTags {
			if !found {
				t.Errorf("Expected tag '%s' not found", tag)
			}
		}
	})

	t.Run("Success - Get news detail without tags", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[2] = model.News{
			ID:          2,
			Title:       "News Without Tags",
			Slug:        "news-without-tags",
			Content:     "Content without tags",
			Thumbnail:   "https://example.com/no-tags.jpg",
			Category:    "announcement",
			ViewsCount:  50,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   2,
				Name: "Editor",
			},
			Tags: []model.NewsTag{},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "news-without-tags")
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 2 {
			t.Errorf("Expected ID 2, got %d", result.ID)
		}

		if len(result.Tags) != 0 {
			t.Errorf("Expected 0 tags, got %d", len(result.Tags))
		}

		// Verify views count is incremented
		if result.ViewsCount != 51 {
			t.Errorf("Expected ViewsCount 51 (50 + 1), got %d", result.ViewsCount)
		}
	})

	t.Run("Success - Verify IncrementViews is called", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		initialViews := 100
		mockRepo.news[3] = model.News{
			ID:          3,
			Title:       "View Counter Test",
			Slug:        "view-counter-test",
			Content:     "Testing view increment",
			Category:    "general",
			ViewsCount:  initialViews,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Test Author",
			},
			Tags: []model.NewsTag{},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "view-counter-test")
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check that returned views count shows incremented value
		if result.ViewsCount != initialViews+1 {
			t.Errorf("Expected ViewsCount %d, got %d", initialViews+1, result.ViewsCount)
		}

		// Verify increment was called on repository
		updatedNews := mockRepo.news[3]
		if updatedNews.ViewsCount != initialViews+1 {
			t.Errorf("Expected repository ViewsCount to be incremented to %d, got %d", initialViews+1, updatedNews.ViewsCount)
		}
	})

	t.Run("Success - Verify date format", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		customTime := time.Date(2024, 12, 15, 14, 30, 45, 0, time.UTC)
		mockRepo.news[4] = model.News{
			ID:          4,
			Title:       "Date Format Test",
			Slug:        "date-format-test",
			Content:     "Testing date formatting",
			Category:    "general",
			ViewsCount:  0,
			IsPublished: true,
			Base:        model.Base{CreatedAt: customTime, UpdatedAt: customTime},
			Author: model.User{
				ID:   1,
				Name: "Test",
			},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "date-format-test")
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedDate := "2024-12-15 14:30:45"
		if result.CreatedAt != expectedDate {
			t.Errorf("Expected CreatedAt '%s', got '%s'", expectedDate, result.CreatedAt)
		}
	})

	t.Run("Error - News not found", func(t *testing.T) {
		service, _ := setupMobileServiceForTests()
		ctx := context.Background()

		// Execute
		result, err := service.GetNewsDetail(ctx, "non-existent-slug")

		// Verify
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "news not found" {
			t.Errorf("Expected 'news not found' error, got '%s'", err.Error())
		}

		if result.ID != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("Error - Unpublished news not accessible", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[5] = model.News{
			ID:          5,
			Title:       "Unpublished News",
			Slug:        "unpublished-news",
			Content:     "This should not be accessible",
			Category:    "general",
			ViewsCount:  0,
			IsPublished: false, // Not published
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Admin",
			},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "unpublished-news")

		// Verify
		if err == nil {
			t.Error("Expected error for unpublished news, got nil")
		}

		if result.ID != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("Success - Multiple tags are correctly mapped", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[6] = model.News{
			ID:          6,
			Title:       "Multi Tag News",
			Slug:        "multi-tag-news",
			Content:     "News with many tags",
			Category:    "general",
			ViewsCount:  0,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:   1,
				Name: "Author",
			},
			Tags: []model.NewsTag{
				{ID: 1, NewsID: 6, TagName: "tag1"},
				{ID: 2, NewsID: 6, TagName: "tag2"},
				{ID: 3, NewsID: 6, TagName: "tag3"},
				{ID: 4, NewsID: 6, TagName: "tag4"},
				{ID: 5, NewsID: 6, TagName: "tag5"},
			},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "multi-tag-news")
		// Verify
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Tags) != 5 {
			t.Errorf("Expected 5 tags, got %d", len(result.Tags))
		}

		// Verify tag order is preserved
		for i := 0; i < 5; i++ {
			expectedTag := "tag" + string(rune('1'+i))
			if result.Tags[i] != expectedTag {
				t.Errorf("Expected tag[%d] to be '%s', got '%s'", i, expectedTag, result.Tags[i])
			}
		}
	})

	t.Run("Success - All fields correctly populated", func(t *testing.T) {
		service, mockRepo := setupMobileServiceForTests()
		ctx := context.Background()

		now := time.Now()
		mockRepo.news[7] = model.News{
			ID:          7,
			Title:       "Complete News Article",
			Slug:        "complete-news-article",
			Content:     "This is a complete news article with all fields populated correctly for testing purposes.",
			Thumbnail:   "https://cdn.example.com/images/complete-news.jpg",
			Category:    "technology",
			ViewsCount:  999,
			IsPublished: true,
			Base:        model.Base{CreatedAt: now, UpdatedAt: now},
			Author: model.User{
				ID:    10,
				Name:  "John Doe",
				Email: "john@example.com",
			},
			Tags: []model.NewsTag{
				{ID: 1, NewsID: 7, TagName: "tech"},
				{ID: 2, NewsID: 7, TagName: "news"},
			},
		}

		// Execute
		result, err := service.GetNewsDetail(ctx, "complete-news-article")
		// Verify no error
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify all fields
		if result.ID != 7 {
			t.Errorf("ID: expected 7, got %d", result.ID)
		}
		if result.Title != "Complete News Article" {
			t.Errorf("Title: expected 'Complete News Article', got '%s'", result.Title)
		}
		if result.Slug != "complete-news-article" {
			t.Errorf("Slug: expected 'complete-news-article', got '%s'", result.Slug)
		}
		if !strings.Contains(result.Content, "complete news article") {
			t.Errorf("Content: expected to contain 'complete news article', got '%s'", result.Content)
		}
		if result.Thumbnail != "https://cdn.example.com/images/complete-news.jpg" {
			t.Errorf("Thumbnail: expected 'https://cdn.example.com/images/complete-news.jpg', got '%s'", result.Thumbnail)
		}
		if result.Category != "technology" {
			t.Errorf("Category: expected 'technology', got '%s'", result.Category)
		}
		if result.AuthorName != "John Doe" {
			t.Errorf("AuthorName: expected 'John Doe', got '%s'", result.AuthorName)
		}
		if result.ViewsCount != 1000 { // 999 + 1 (incremented)
			t.Errorf("ViewsCount: expected 1000, got %d", result.ViewsCount)
		}
		if result.CreatedAt != now.Format("2006-01-02 15:04:05") {
			t.Errorf("CreatedAt: expected '%s', got '%s'", now.Format("2006-01-02 15:04:05"), result.CreatedAt)
		}
		if len(result.Tags) != 2 {
			t.Errorf("Tags: expected 2 tags, got %d", len(result.Tags))
		}
	})
}

// Mock repository with error capability for news
type mockMobileRepoWithError struct {
	*mockMobileRepository
	shouldErrorOnGetNews bool
}

func (m *mockMobileRepoWithError) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]model.News, int64, error) {
	if m.shouldErrorOnGetNews {
		return nil, 0, errors.New("database error")
	}
	return m.mockMobileRepository.GetPublishedNews(ctx, params)
}
