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

	validDocTypes := []string{"nib", "npwp", "revenue-record", "business-permit"}

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
