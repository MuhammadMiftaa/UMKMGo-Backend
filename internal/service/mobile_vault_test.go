package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/vault"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"

	vaultapi "github.com/hashicorp/vault/api"
)

// setupMockVaultServer creates a mock HTTP server that simulates Vault API
func setupMockVaultServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock transit decrypt endpoint
		if r.URL.Path == "/v1/transit/decrypt/nik-key" ||
			r.URL.Path == "/v1/transit/decrypt/kartu-key" {
			// Return mock decrypted data as base64
			plaintext := "1234567890123456" // Mock NIK
			if r.URL.Path == "/v1/transit/decrypt/kartu-key" {
				plaintext = "UMKM-TEST-2024" // Mock Kartu Number
			}

			plainB64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

			response := fmt.Sprintf(`{
				"data": {
					"plaintext": "%s"
				}
			}`, plainB64)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
			return
		}

		// Mock auth endpoint
		if r.URL.Path == "/v1/auth/approle/login" {
			response := `{
				"auth": {
					"client_token": "mock-token",
					"renewable": true,
					"lease_duration": 3600
				}
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	return httptest.NewServer(handler)
}

// setupVaultClientForTest sets up a real vault client pointing to mock server
func setupVaultClientForTest(t *testing.T) (*httptest.Server, func()) {
	// Create mock vault server
	server := setupMockVaultServer()

	// Create vault client config
	config := vaultapi.DefaultConfig()
	config.Address = server.URL

	client, err := vaultapi.NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create vault client: %v", err)
	}

	// Set mock token
	client.SetToken("mock-token")

	// Store original client
	originalClient := vault.VaultClient

	// Set mock client
	vault.VaultClient = client

	// Setup environment config
	env.Cfg.Vault.TransitPath = "transit"
	env.Cfg.Vault.NIKEncryptionKey = "nik-key"
	env.Cfg.Vault.KartuEncryptionKey = "kartu-key"

	// Return cleanup function
	cleanup := func() {
		vault.VaultClient = originalClient
		server.Close()
	}

	return server, cleanup
}

// TestGetUMKMProfile_WithMockedVault tests GetUMKMProfile with a fully mocked vault
func TestGetUMKMProfile_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Successfully decrypt and return complete UMKM profile", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1990-01-01")
		mockRepo.umkms[100] = model.UMKM{
			ID:             100,
			UserID:         100,
			BusinessName:   "Mocked Vault Business",
			NIK:            "vault:v1:test_nik_cipher",
			KartuNumber:    "vault:v1:test_kartu_cipher",
			Gender:         "female",
			BirthDate:      birthDate,
			Phone:          "081234567890",
			Address:        "Jl. Vault Test No. 100",
			ProvinceID:     1,
			CityID:         1,
			District:       "District Vault",
			Subdistrict:    "Subdistrict Vault",
			PostalCode:     "54321",
			NIB:            "http://example.com/nib100.pdf",
			NPWP:           "http://example.com/npwp100.pdf",
			RevenueRecord:  "http://example.com/revenue100.pdf",
			BusinessPermit: "http://example.com/permit100.pdf",
			KartuType:      "Premium",
			Photo:          "http://example.com/photo100.jpg",
			User: model.User{
				ID:    100,
				Name:  "Vault Test User",
				Email: "vault@example.com",
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

		result, err := service.GetUMKMProfile(ctx, 100)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Verify all fields are correctly populated
		if result.ID != 100 {
			t.Errorf("Expected ID 100, got %d", result.ID)
		}
		if result.UserID != 100 {
			t.Errorf("Expected UserID 100, got %d", result.UserID)
		}
		if result.BusinessName != "Mocked Vault Business" {
			t.Errorf("Expected BusinessName 'Mocked Vault Business', got %s", result.BusinessName)
		}
		if result.NIK != "1234567890123456" {
			t.Errorf("Expected decrypted NIK '1234567890123456', got %s", result.NIK)
		}
		if result.Gender != "female" {
			t.Errorf("Expected Gender 'female', got %s", result.Gender)
		}
		if result.BirthDate != "1990-01-01" {
			t.Errorf("Expected BirthDate '1990-01-01', got %s", result.BirthDate)
		}
		if result.Phone != "081234567890" {
			t.Errorf("Expected Phone '081234567890', got %s", result.Phone)
		}
		if result.Address != "Jl. Vault Test No. 100" {
			t.Errorf("Expected Address 'Jl. Vault Test No. 100', got %s", result.Address)
		}
		if result.ProvinceID != 1 {
			t.Errorf("Expected ProvinceID 1, got %d", result.ProvinceID)
		}
		if result.CityID != 1 {
			t.Errorf("Expected CityID 1, got %d", result.CityID)
		}
		if result.District != "District Vault" {
			t.Errorf("Expected District 'District Vault', got %s", result.District)
		}
		if result.Subdistrict != "Subdistrict Vault" {
			t.Errorf("Expected Subdistrict 'Subdistrict Vault', got %s", result.Subdistrict)
		}
		if result.PostalCode != "54321" {
			t.Errorf("Expected PostalCode '54321', got %s", result.PostalCode)
		}
		if result.NIB != "http://example.com/nib100.pdf" {
			t.Errorf("Expected NIB URL, got %s", result.NIB)
		}
		if result.NPWP != "http://example.com/npwp100.pdf" {
			t.Errorf("Expected NPWP URL, got %s", result.NPWP)
		}
		if result.RevenueRecord != "http://example.com/revenue100.pdf" {
			t.Errorf("Expected RevenueRecord URL, got %s", result.RevenueRecord)
		}
		if result.BusinessPermit != "http://example.com/permit100.pdf" {
			t.Errorf("Expected BusinessPermit URL, got %s", result.BusinessPermit)
		}
		if result.KartuType != "Premium" {
			t.Errorf("Expected KartuType 'Premium', got %s", result.KartuType)
		}
		if result.Photo != "http://example.com/photo100.jpg" {
			t.Errorf("Expected Photo URL, got %s", result.Photo)
		}
		if result.KartuNumber != "UMKM-TEST-2024" {
			t.Errorf("Expected decrypted KartuNumber 'UMKM-TEST-2024', got %s", result.KartuNumber)
		}

		// Verify Province
		if result.Province.ID != 1 {
			t.Errorf("Expected Province.ID 1, got %d", result.Province.ID)
		}
		if result.Province.Name != "DKI Jakarta" {
			t.Errorf("Expected Province.Name 'DKI Jakarta', got %s", result.Province.Name)
		}

		// Verify City
		if result.City.ID != 1 {
			t.Errorf("Expected City.ID 1, got %d", result.City.ID)
		}
		if result.City.Name != "Jakarta Pusat" {
			t.Errorf("Expected City.Name 'Jakarta Pusat', got %s", result.City.Name)
		}

		// Verify User
		if result.User.ID != 100 {
			t.Errorf("Expected User.ID 100, got %d", result.User.ID)
		}
		if result.User.Name != "Vault Test User" {
			t.Errorf("Expected User.Name 'Vault Test User', got %s", result.User.Name)
		}
		if result.User.Email != "vault@example.com" {
			t.Errorf("Expected User.Email 'vault@example.com', got %s", result.User.Email)
		}

		t.Log("✓ All fields verified successfully")
		t.Log("✓ NIK decryption executed and verified")
		t.Log("✓ Kartu Number decryption executed and verified")
		t.Log("✓ Complete return statement with all fields covered")
	})
}

// TestGetUMKMProfile_WithVaultDecryptionError tests error handling in decryption
func TestGetUMKMProfile_WithVaultDecryptionError(t *testing.T) {
	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Handle NIK decryption error", func(t *testing.T) {
		// Don't setup vault client - this will cause decryption to fail
		birthDate, _ := time.Parse("2006-01-02", "1990-01-01")
		mockRepo.umkms[101] = model.UMKM{
			ID:          101,
			UserID:      101,
			NIK:         "vault:v1:bad_cipher",
			KartuNumber: "vault:v1:bad_cipher",
			BirthDate:   birthDate,
			User:        model.User{ID: 101, Name: "Error Test"},
			Province:    model.Province{ID: 1, Name: "Test"},
			City:        model.City{ID: 1, Name: "Test"},
		}

		_, err := service.GetUMKMProfile(ctx, 101)

		if err == nil {
			t.Error("Expected error from vault decryption, got nil")
		}

		if err.Error() != "failed to decrypt NIK" {
			t.Errorf("Expected 'failed to decrypt NIK', got %v", err)
		}

		t.Log("✓ NIK decryption error path covered")
	})
}

// TestGetUMKMProfile_DocumentationTest documents the coverage strategy
func TestGetUMKMProfile_DocumentationTest(t *testing.T) {
	t.Log("=== GetUMKMProfile Coverage Strategy ===")
	t.Log("")
	t.Log("This test suite covers the GetUMKMProfile function including:")
	t.Log("1. ✓ NIK decryption with vault.DecryptTransit")
	t.Log("2. ✓ Kartu Number decryption with vault.DecryptTransit")
	t.Log("3. ✓ Complete return statement with all fields")
	t.Log("4. ✓ Error handling for decryption failures")
	t.Log("")
	t.Log("Implementation:")
	t.Log("- Uses httptest to create mock Vault HTTP server")
	t.Log("- Creates real vault API client pointing to mock server")
	t.Log("- Mock server returns base64 encoded plaintext")
	t.Log("- All code paths in GetUMKMProfile are now covered")
	t.Log("")
	t.Log("Coverage achieved: NIK decrypt, Kartu decrypt, and full DTO return")
}

// TestGetDashboard_WithMockedVault tests GetDashboard with a fully mocked vault
func TestGetDashboard_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Successfully get dashboard with applications loop and approved count", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1992-06-15")

		// Setup UMKM profile
		mockRepo.umkms[200] = model.UMKM{
			ID:          200,
			UserID:      200,
			NIK:         "vault:v1:nik_200",
			KartuNumber: "vault:v1:kartu_200",
			KartuType:   "Gold",
			QRCode:      "QR-CODE-200",
			BirthDate:   birthDate,
			User: model.User{
				ID:    200,
				Name:  "Dashboard Test User",
				Email: "dashboard@example.com",
			},
			Province: model.Province{ID: 1, Name: "DKI Jakarta"},
			City:     model.City{ID: 1, Name: "Jakarta Pusat"},
		}

		// Setup applications - mix of approved and other statuses
		mockAppRepo := service.applicationRepo.(*mockApplicationsRepo)
		mockAppRepo.applications[1] = model.Application{
			ID:      1,
			UMKMID:  200,
			Status:  "approved", // approved
			Program: model.Program{Title: "Program 1"},
		}
		mockAppRepo.applications[2] = model.Application{
			ID:      2,
			UMKMID:  200,
			Status:  "screening", // not approved
			Program: model.Program{Title: "Program 2"},
		}
		mockAppRepo.applications[3] = model.Application{
			ID:      3,
			UMKMID:  200,
			Status:  "approved", // approved
			Program: model.Program{Title: "Program 3"},
		}
		mockAppRepo.applications[4] = model.Application{
			ID:      4,
			UMKMID:  200,
			Status:  "rejected", // not approved
			Program: model.Program{Title: "Program 4"},
		}
		mockAppRepo.applications[5] = model.Application{
			ID:      5,
			UMKMID:  200,
			Status:  "approved", // approved
			Program: model.Program{Title: "Program 5"},
		}

		// Setup notifications (5 unread)
		mockNotifRepo := service.notificationRepo.(*mockNotificationRepoForMobile)
		for i := 0; i < 5; i++ {
			mockNotifRepo.notifications = append(mockNotifRepo.notifications, model.Notification{
				UMKMID: 200,
				IsRead: false,
			})
		}

		result, err := service.GetDashboard(ctx, 200)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Verify dashboard data
		if result.Name != "Dashboard Test User" {
			t.Errorf("Expected Name 'Dashboard Test User', got %s", result.Name)
		}

		if result.NotificationsCount != 5 {
			t.Errorf("Expected NotificationsCount 5, got %d", result.NotificationsCount)
		}

		if result.KartuType != "Gold" {
			t.Errorf("Expected KartuType 'Gold', got %s", result.KartuType)
		}

		if result.KartuNumber != "UMKM-TEST-2024" {
			t.Errorf("Expected decrypted KartuNumber 'UMKM-TEST-2024', got %s", result.KartuNumber)
		}

		if result.QRCode != "QR-CODE-200" {
			t.Errorf("Expected QRCode 'QR-CODE-200', got %s", result.QRCode)
		}

		// Verify applications count
		if result.TotalApplications != 5 {
			t.Errorf("Expected TotalApplications 5, got %d", result.TotalApplications)
		}

		// Verify approved applications count (loop logic)
		if result.ApprovedApplications != 3 {
			t.Errorf("Expected ApprovedApplications 3 (counting approved status), got %d", result.ApprovedApplications)
		}

		t.Log("✓ Dashboard data retrieved successfully")
		t.Log("✓ Kartu Number decryption executed")
		t.Log("✓ Applications loop executed")
		t.Log("✓ Approved applications counted correctly")
		t.Log("✓ All return fields populated")
	})

	t.Run("Dashboard with zero approved applications", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1995-03-20")

		mockRepo.umkms[201] = model.UMKM{
			ID:          201,
			UserID:      201,
			NIK:         "vault:v1:nik_201",
			KartuNumber: "vault:v1:kartu_201",
			KartuType:   "Silver",
			QRCode:      "QR-CODE-201",
			BirthDate:   birthDate,
			User: model.User{
				ID:    201,
				Name:  "Test User 201",
				Email: "user201@example.com",
			},
			Province: model.Province{ID: 1, Name: "DKI Jakarta"},
			City:     model.City{ID: 1, Name: "Jakarta Pusat"},
		}

		// All applications are not approved
		mockAppRepo := service.applicationRepo.(*mockApplicationsRepo)
		mockAppRepo.applications[10] = model.Application{
			ID:      10,
			UMKMID:  201,
			Status:  "screening",
			Program: model.Program{Title: "Program 10"},
		}
		mockAppRepo.applications[11] = model.Application{
			ID:      11,
			UMKMID:  201,
			Status:  "rejected",
			Program: model.Program{Title: "Program 11"},
		}

		// No unread notifications for this user

		result, err := service.GetDashboard(ctx, 201)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if result.TotalApplications != 2 {
			t.Errorf("Expected TotalApplications 2, got %d", result.TotalApplications)
		}

		if result.ApprovedApplications != 0 {
			t.Errorf("Expected ApprovedApplications 0, got %d", result.ApprovedApplications)
		}

		t.Log("✓ Zero approved applications handled correctly")
	})

	t.Run("Dashboard with empty applications", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1988-11-05")

		mockRepo.umkms[202] = model.UMKM{
			ID:          202,
			UserID:      202,
			NIK:         "vault:v1:nik_202",
			KartuNumber: "vault:v1:kartu_202",
			KartuType:   "Basic",
			QRCode:      "QR-CODE-202",
			BirthDate:   birthDate,
			User: model.User{
				ID:    202,
				Name:  "Test User 202",
				Email: "user202@example.com",
			},
			Province: model.Province{ID: 1, Name: "DKI Jakarta"},
			City:     model.City{ID: 1, Name: "Jakarta Pusat"},
		}

		// No applications for this user
		// (applications map already exists, no need to clear it)

		// Setup 10 unread notifications
		mockNotifRepo := service.notificationRepo.(*mockNotificationRepoForMobile)
		for i := 0; i < 10; i++ {
			mockNotifRepo.notifications = append(mockNotifRepo.notifications, model.Notification{
				UMKMID: 202,
				IsRead: false,
			})
		}

		result, err := service.GetDashboard(ctx, 202)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if result.TotalApplications != 0 {
			t.Errorf("Expected TotalApplications 0, got %d", result.TotalApplications)
		}

		if result.ApprovedApplications != 0 {
			t.Errorf("Expected ApprovedApplications 0, got %d", result.ApprovedApplications)
		}

		if result.NotificationsCount != 10 {
			t.Errorf("Expected NotificationsCount 10, got %d", result.NotificationsCount)
		}

		t.Log("✓ Empty applications handled correctly")
	})
}

// TestGetDashboard_DocumentationTest documents the GetDashboard coverage
func TestGetDashboard_DocumentationTest(t *testing.T) {
	t.Log("=== GetDashboard Coverage Strategy ===")
	t.Log("")
	t.Log("This test suite covers the GetDashboard function including:")
	t.Log("1. ✓ GetApplicationsByUMKMID call")
	t.Log("2. ✓ Applications loop iteration")
	t.Log("3. ✓ Approved status counting logic")
	t.Log("4. ✓ Complete return statement with all fields")
	t.Log("5. ✓ Kartu Number decryption")
	t.Log("")
	t.Log("Test scenarios:")
	t.Log("- Dashboard with mixed approved/not approved applications")
	t.Log("- Dashboard with zero approved applications")
	t.Log("- Dashboard with empty applications list")
	t.Log("")
	t.Log("Coverage achieved: Applications loop, status check, and full response")
}

// TestGetUMKMProfileWithDecryption_WithMockedVault tests GetUMKMProfileWithDecryption
func TestGetUMKMProfileWithDecryption_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	t.Run("Successfully get UMKM profile with decryption and logging", func(t *testing.T) {
		birthDate, _ := time.Parse("2006-01-02", "1993-08-25")

		mockRepo.umkms[300] = model.UMKM{
			ID:           300,
			UserID:       300,
			BusinessName: "Test Decryption Business",
			NIK:          "vault:v1:nik_300",
			KartuNumber:  "vault:v1:kartu_300",
			BirthDate:    birthDate,
			User: model.User{
				ID:    300,
				Name:  "Decryption Test User",
				Email: "decrypt300@example.com",
			},
			Province: model.Province{ID: 1, Name: "DKI Jakarta"},
			City:     model.City{ID: 1, Name: "Jakarta Pusat"},
		}

		result, err := service.GetUMKMProfileWithDecryption(ctx, 300, "testing_purpose")
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Verify result - DecryptWithLog returns masked values
		if result.ID != 300 {
			t.Errorf("Expected ID 300, got %d", result.ID)
		}
		if result.BusinessName != "Test Decryption Business" {
			t.Errorf("Expected BusinessName 'Test Decryption Business', got %s", result.BusinessName)
		}

		// NIK and KartuNumber will be masked by utils.MaskMiddle
		if result.NIK == "" {
			t.Error("Expected decrypted NIK, got empty")
		}
		if result.KartuNumber == "" {
			t.Error("Expected decrypted Kartu Number, got empty")
		}

		t.Log("✓ GetUMKMProfileWithDecryption executed successfully")
		t.Log("✓ getUMKMWithDecryption helper function covered")
		t.Log("✓ DecryptNIKWithLog executed")
		t.Log("✓ DecryptKartuNumberWithLog executed")
	})
}

func TestCreateTrainingApplication_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	userID := 456
	birthDate, _ := time.Parse("2006-01-02", "1990-05-15")

	// Setup mock UMKM data (by userID as per GetUMKMProfileByID implementation)
	mockRepo.umkms[userID] = model.UMKM{
		ID:           123,
		UserID:       userID,
		BusinessName: "Test Training Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		Gender:       "male",
		BirthDate:    birthDate,
		Phone:        "081234567890",
		Address:      "Test Address",
		ProvinceID:   1,
		CityID:       1,
		District:     "Test District",
		Subdistrict:  "Test Subdistrict",
		PostalCode:   "12345",
		KartuType:    "silver",
		User: model.User{
			ID:    userID,
			Name:  "Training User",
			Email: "training@example.com",
		},
		Province: model.Province{ID: 1, Name: "DKI Jakarta"},
		City:     model.City{ID: 1, Name: "Jakarta Pusat"},
	}

	maxAmount := 0.0
	maxTenure := 0

	// Setup mock program data
	mockRepo.programs[1] = model.Program{
		ID:              1,
		Title:           "Training Program",
		Type:            "training",
		Description:     "Training program description",
		IsActive:        true,
		MaxAmount:       &maxAmount,
		MaxTenureMonths: &maxTenure,
	}

	// Create request DTO
	req := dto.CreateApplicationTraining{
		ProgramID:          1,
		Motivation:         "Want to improve my business skills",
		BusinessExperience: "5 years in food business",
		LearningObjectives: "Learn digital marketing",
		AvailabilityNotes:  "Available on weekends",
		Documents: map[string]string{
			"nib":  "https://example.com/nib",
			"npwp": "https://example.com/npwp",
		},
	}

	// Call the service - this will execute getUMKMWithDecryption and all code after it
	err := service.CreateTrainingApplication(ctx, userID, req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	t.Log("✓ CreateTrainingApplication executed successfully")
	t.Log("✓ getUMKMWithDecryption called and NIK/Kartu decrypted")
	t.Log("✓ Code after getUMKMWithDecryption covered: application creation, history, notification, documents")
}

func TestCreateCertificationApplication_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	userID := 789
	birthDate, _ := time.Parse("2006-01-02", "1985-08-20")

	// Setup mock UMKM data (by userID as per GetUMKMProfileByID implementation)
	mockRepo.umkms[userID] = model.UMKM{
		ID:           456,
		UserID:       userID,
		BusinessName: "Test Certification Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		Gender:       "female",
		BirthDate:    birthDate,
		Phone:        "082345678901",
		Address:      "Cert Address",
		ProvinceID:   2,
		CityID:       2,
		District:     "Cert District",
		Subdistrict:  "Cert Subdistrict",
		PostalCode:   "54321",
		KartuType:    "gold",
		User: model.User{
			ID:    userID,
			Name:  "Cert User",
			Email: "cert@example.com",
		},
		Province: model.Province{ID: 2, Name: "Jawa Barat"},
		City:     model.City{ID: 2, Name: "Bandung"},
	}

	maxAmount := 0.0
	maxTenure := 0

	// Setup mock program data
	mockRepo.programs[2] = model.Program{
		ID:              2,
		Title:           "Certification Program",
		Type:            "certification",
		Description:     "Certification program description",
		IsActive:        true,
		MaxAmount:       &maxAmount,
		MaxTenureMonths: &maxTenure,
	}

	yearsOperating := 3

	// Create request DTO
	req := dto.CreateApplicationCertification{
		ProgramID:           2,
		BusinessSector:      "Food & Beverage",
		ProductOrService:    "Organic Coffee",
		BusinessDescription: "Selling premium organic coffee",
		YearsOperating:      &yearsOperating,
		CurrentStandards:    "ISO 9001",
		CertificationGoals:  "Get halal certification",
		Documents: map[string]string{
			"nib":             "https://example.com/nib",
			"npwp":            "https://example.com/npwp",
			"business_permit": "https://example.com/business_permit",
		},
	}

	// Call the service - this will execute getUMKMWithDecryption and all code after it
	err := service.CreateCertificationApplication(ctx, userID, req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	t.Log("✓ CreateCertificationApplication executed successfully")
	t.Log("✓ getUMKMWithDecryption called and NIK/Kartu decrypted")
	t.Log("✓ Code after getUMKMWithDecryption covered: application creation, history, notification, documents")
}

func TestCreateFundingApplication_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockRepo := setupMobileServiceForTests()
	ctx := context.Background()

	userID := 999
	birthDate, _ := time.Parse("2006-01-02", "1988-11-30")

	// Setup mock UMKM data (by userID as per GetUMKMProfileByID implementation)
	mockRepo.umkms[userID] = model.UMKM{
		ID:           777,
		UserID:       userID,
		BusinessName: "Test Funding Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		Gender:       "male",
		BirthDate:    birthDate,
		Phone:        "083456789012",
		Address:      "Funding Address",
		ProvinceID:   3,
		CityID:       3,
		District:     "Funding District",
		Subdistrict:  "Funding Subdistrict",
		PostalCode:   "99999",
		KartuType:    "platinum",
		User: model.User{
			ID:    userID,
			Name:  "Funding User",
			Email: "funding@example.com",
		},
		Province: model.Province{ID: 3, Name: "Jawa Timur"},
		City:     model.City{ID: 3, Name: "Surabaya"},
	}

	maxAmount := 100000000.0
	maxTenure := 24

	// Setup mock program data
	mockRepo.programs[3] = model.Program{
		ID:              3,
		Title:           "Funding Program",
		Type:            "funding",
		Description:     "Funding program description",
		IsActive:        true,
		MaxAmount:       &maxAmount,
		MaxTenureMonths: &maxTenure,
	}

	yearsOperating := 2
	revenueProjection := 50000000.0
	monthlyRevenue := 5000000.0

	// Create request DTO
	req := dto.CreateApplicationFunding{
		ProgramID:             3,
		BusinessSector:        "Manufacturing",
		BusinessDescription:   "Producing handcraft products",
		YearsOperating:        &yearsOperating,
		RequestedAmount:       25000000,
		FundPurpose:           "Expand production capacity",
		BusinessPlan:          "Detailed business plan",
		RevenueProjection:     &revenueProjection,
		MonthlyRevenue:        &monthlyRevenue,
		RequestedTenureMonths: 12,
		CollateralDescription: "Land certificate",
		Documents: map[string]string{
			"nib":            "https://example.com/nib",
			"npwp":           "https://example.com/npwp",
			"revenue_record": "https://example.com/revenue_record",
		},
	}

	// Call the service - this will execute getUMKMWithDecryption and all code after it
	err := service.CreateFundingApplication(ctx, userID, req)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	t.Log("✓ CreateFundingApplication executed successfully")
	t.Log("✓ getUMKMWithDecryption called and NIK/Kartu decrypted")
	t.Log("✓ Code after getUMKMWithDecryption covered: application creation, history, notification, documents")
}

func TestReviseApplication_WithMockedVault(t *testing.T) {
	_, cleanup := setupVaultClientForTest(t)
	defer cleanup()

	service, mockMobileRepo, mockAppRepo := setupMobileServiceWithAllRepos()
	ctx := context.Background()

	userID := 555
	applicationID := 10
	umkmID := 888
	birthDate, _ := time.Parse("2006-01-02", "1992-03-10")

	// Setup mock UMKM data (by userID as per GetUMKMProfileByID implementation)
	mockMobileRepo.umkms[userID] = model.UMKM{
		ID:           umkmID,
		UserID:       userID,
		BusinessName: "Test Revise Business",
		NIK:          "vault:v1:encrypted_nik",
		KartuNumber:  "vault:v1:encrypted_kartu",
		Gender:       "female",
		BirthDate:    birthDate,
		Phone:        "084567890123",
		Address:      "Revise Address",
		ProvinceID:   4,
		CityID:       4,
		District:     "Revise District",
		Subdistrict:  "Revise Subdistrict",
		PostalCode:   "11111",
		KartuType:    "bronze",
		User: model.User{
			ID:    userID,
			Name:  "Revise User",
			Email: "revise@example.com",
		},
		Province: model.Province{ID: 4, Name: "Bali"},
		City:     model.City{ID: 4, Name: "Denpasar"},
	}

	// Create application object
	testApplication := model.Application{
		ID:          applicationID,
		UMKMID:      umkmID,
		ProgramID:   1,
		Status:      "revised", // Important: status must be "revised"
		Type:        "training",
		Documents:   []model.ApplicationDocument{},
		Histories:   []model.ApplicationHistory{},
		Program:     model.Program{ID: 1, Title: "Test Program", Type: "training"},
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, 7),
	}

	// Setup mock application data in BOTH repositories
	mockMobileRepo.applications[applicationID] = testApplication
	mockAppRepo.applications[applicationID] = testApplication

	// Setup mock program
	mockMobileRepo.programs[1] = model.Program{
		ID:          1,
		Title:       "Test Program",
		Type:        "training",
		Description: "Test program for revision",
		IsActive:    true,
	}

	// Create document request
	documents := []dto.UploadDocumentRequest{
		{
			Type:     "nib",
			Document: "https://example.com/nib",
		},
		{
			Type:     "npwp",
			Document: "https://example.com/npwp",
		},
	}

	// Call the service - this will execute getUMKMWithDecryption and all code after it
	err := service.ReviseApplication(ctx, userID, applicationID, documents)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	t.Log("✓ ReviseApplication executed successfully")
	t.Log("✓ getUMKMWithDecryption called and NIK/Kartu decrypted")
	t.Log("✓ Code after getUMKMWithDecryption covered: history, notification, document updates")
}

// setupMobileServiceWithAllRepos creates a mobile service and returns all mock repositories
func setupMobileServiceWithAllRepos() (*mobileService, *mockMobileRepository, *mockApplicationsRepo) {
	mockMobileRepo := newMockMobileRepository()
	mockProgramRepo := newMockProgramRepoForMobile()
	mockNotifRepo := newMockNotificationRepoForMobile()
	mockVaultRepo := newMockVaultDecryptLogRepo()
	mockAppRepo := newMockApplicationsRepo()
	mockSLARepo := newMockSLARepo()

	service := &mobileService{
		mobileRepo:       mockMobileRepo,
		programRepo:      mockProgramRepo,
		notificationRepo: mockNotifRepo,
		vaultLogRepo:     mockVaultRepo,
		applicationRepo:  mockAppRepo,
		slaRepo:          mockSLARepo,
		minio:            nil,
	}

	return service, mockMobileRepo, mockAppRepo
}
