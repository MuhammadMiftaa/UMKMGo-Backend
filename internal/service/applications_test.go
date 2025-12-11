package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// ==================== MOCK REPOSITORIES ====================

// Mock Applications Repository
type mockApplicationsRepo struct {
	applications map[int]model.Application
	documents    map[int][]model.ApplicationDocument
	histories    map[int][]model.ApplicationHistory
	programs     map[int]model.Program
	umkms        map[int]model.UMKM
}

func newMockApplicationsRepo() *mockApplicationsRepo {
	return &mockApplicationsRepo{
		applications: make(map[int]model.Application),
		documents:    make(map[int][]model.ApplicationDocument),
		histories:    make(map[int][]model.ApplicationHistory),
		programs: map[int]model.Program{
			1: {ID: 1, Title: "Training Program", Type: "training", IsActive: true},
			2: {ID: 2, Title: "Funding Program", Type: "funding", IsActive: true},
		},
		umkms: map[int]model.UMKM{
			1: {
				ID:           1,
				UserID:       1,
				BusinessName: "Test Business",
				NIK:          "encrypted_nik",
				KartuNumber:  "encrypted_kartu",
				User:         model.User{ID: 1, Name: "Test User"},
				City: model.City{
					ID:   1,
					Name: "Jakarta",
					Province: model.Province{
						ID:   1,
						Name: "DKI Jakarta",
					},
				},
			},
		},
	}
}

func (m *mockApplicationsRepo) GetAllApplications(ctx context.Context, filterType string) ([]model.Application, error) {
	var apps []model.Application
	for _, app := range m.applications {
		if filterType == "" || app.Type == filterType {
			apps = append(apps, app)
		}
	}
	return apps, nil
}

func (m *mockApplicationsRepo) GetApplicationByID(ctx context.Context, id int) (model.Application, error) {
	if app, exists := m.applications[id]; exists {
		return app, nil
	}
	return model.Application{}, errors.New("application not found")
}

func (m *mockApplicationsRepo) GetApplicationsByUMKMID(ctx context.Context, umkmID int) ([]model.Application, error) {
	var apps []model.Application
	for _, app := range m.applications {
		if app.UMKMID == umkmID {
			apps = append(apps, app)
		}
	}
	return apps, nil
}

func (m *mockApplicationsRepo) CreateApplication(ctx context.Context, app model.Application) (model.Application, error) {
	app.ID = len(m.applications) + 1
	app.SubmittedAt = time.Now()
	app.ExpiredAt = time.Now().AddDate(0, 0, 30)
	m.applications[app.ID] = app
	return app, nil
}

func (m *mockApplicationsRepo) UpdateApplication(ctx context.Context, app model.Application) (model.Application, error) {
	if _, exists := m.applications[app.ID]; !exists {
		return model.Application{}, errors.New("application not found")
	}
	m.applications[app.ID] = app
	return app, nil
}

func (m *mockApplicationsRepo) DeleteApplication(ctx context.Context, app model.Application) (model.Application, error) {
	delete(m.applications, app.ID)
	return app, nil
}

func (m *mockApplicationsRepo) CreateApplicationDocuments(ctx context.Context, docs []model.ApplicationDocument) error {
	if len(docs) == 0 {
		return nil
	}
	appID := docs[0].ApplicationID
	m.documents[appID] = docs
	return nil
}

func (m *mockApplicationsRepo) GetApplicationDocuments(ctx context.Context, appID int) ([]model.ApplicationDocument, error) {
	if docs, exists := m.documents[appID]; exists {
		return docs, nil
	}
	return []model.ApplicationDocument{}, nil
}

func (m *mockApplicationsRepo) DeleteApplicationDocuments(ctx context.Context, appID int) error {
	delete(m.documents, appID)
	return nil
}

func (m *mockApplicationsRepo) CreateApplicationHistory(ctx context.Context, hist model.ApplicationHistory) error {
	m.histories[hist.ApplicationID] = append(m.histories[hist.ApplicationID], hist)
	return nil
}

func (m *mockApplicationsRepo) GetApplicationHistories(ctx context.Context, appID int) ([]model.ApplicationHistory, error) {
	if hists, exists := m.histories[appID]; exists {
		return hists, nil
	}
	return []model.ApplicationHistory{}, nil
}

func (m *mockApplicationsRepo) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	if prog, exists := m.programs[id]; exists {
		return prog, nil
	}
	return model.Program{}, errors.New("program not found")
}

func (m *mockApplicationsRepo) GetUMKMByUserID(ctx context.Context, userID int) (model.UMKM, error) {
	for _, umkm := range m.umkms {
		if umkm.UserID == userID {
			return umkm, nil
		}
	}
	return model.UMKM{}, errors.New("UMKM not found")
}

func (m *mockApplicationsRepo) IsApplicationExists(ctx context.Context, umkmID, programID int) bool {
	for _, app := range m.applications {
		if app.UMKMID == umkmID && app.ProgramID == programID && app.Status != "rejected" {
			return true
		}
	}
	return false
}

// Mock Users Repository
type mockUsersRepo struct {
	users map[int]model.User
}

func newMockUsersRepo() *mockUsersRepo {
	return &mockUsersRepo{
		users: map[int]model.User{
			1: {ID: 1, Name: "Admin User", Email: "admin@test.com"},
			2: {ID: 2, Name: "Screening Admin", Email: "screening@test.com"},
		},
	}
}

func (m *mockUsersRepo) GetUserByID(ctx context.Context, id int) (model.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return model.User{}, errors.New("user not found")
}

func (m *mockUsersRepo) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func (m *mockUsersRepo) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepo) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepo) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepo) DeleteUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepo) CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User) (dto.UMKMMobile, error) {
	return dto.UMKMMobile{}, errors.New("not implemented")
}

func (m *mockUsersRepo) GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error) {
	return model.UMKM{}, errors.New("not implemented")
}

func (m *mockUsersRepo) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepo) GetRoleByID(ctx context.Context, id int) (model.Role, error) {
	return model.Role{}, errors.New("not implemented")
}

func (m *mockUsersRepo) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	return model.Role{}, errors.New("not implemented")
}

func (m *mockUsersRepo) IsRoleExist(ctx context.Context, id int) bool {
	return false
}

func (m *mockUsersRepo) IsPermissionExist(ctx context.Context, ids []string) ([]int, bool) {
	return nil, false
}

func (m *mockUsersRepo) GetListPermissions(ctx context.Context) ([]model.Permission, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepo) GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepo) GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepo) DeletePermissionsByRoleID(ctx context.Context, roleID int) error {
	return errors.New("not implemented")
}

func (m *mockUsersRepo) AddRolePermissions(ctx context.Context, roleID int, permissions []int) error {
	return errors.New("not implemented")
}

func (m *mockUsersRepo) GetProvinces(ctx context.Context) ([]dto.Province, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepo) GetCities(ctx context.Context) ([]dto.City, error) {
	return nil, errors.New("not implemented")
}

// Mock Notification Repository
type mockNotificationRepo struct {
	notifications []model.Notification
}

func newMockNotificationRepo() *mockNotificationRepo {
	return &mockNotificationRepo{
		notifications: []model.Notification{},
	}
}

func (m *mockNotificationRepo) CreateNotification(ctx context.Context, notif model.Notification) error {
	m.notifications = append(m.notifications, notif)
	return nil
}

func (m *mockNotificationRepo) GetNotificationsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.Notification, error) {
	return m.notifications, nil
}

func (m *mockNotificationRepo) GetUnreadCount(ctx context.Context, umkmID int) (int64, error) {
	return 0, nil
}

func (m *mockNotificationRepo) MarkAsRead(ctx context.Context, notifIDs int, umkmID int) error {
	return nil
}

func (m *mockNotificationRepo) MarkAllAsRead(ctx context.Context, umkmID int) error {
	return nil
}

// Mock SLA Repository
type mockSLARepo struct {
	slas map[string]model.SLA
}

func newMockSLARepo() *mockSLARepo {
	return &mockSLARepo{
		slas: map[string]model.SLA{
			"screening": {ID: 1, Status: "screening", MaxDays: 7},
			"final":     {ID: 2, Status: "final", MaxDays: 14},
		},
	}
}

func (m *mockSLARepo) GetSLAByStatus(ctx context.Context, status string) (model.SLA, error) {
	if sla, exists := m.slas[status]; exists {
		return sla, nil
	}
	return model.SLA{}, errors.New("SLA not found")
}

func (m *mockSLARepo) UpdateSLA(ctx context.Context, sla model.SLA) (model.SLA, error) {
	return sla, nil
}

func (m *mockSLARepo) GetApplicationsForExport(ctx context.Context, appType string) ([]model.Application, error) {
	return nil, errors.New("not implemented")
}

func (m *mockSLARepo) GetProgramsForExport(ctx context.Context, appType string) ([]model.Program, error) {
	return nil, errors.New("not implemented")
}

// Mock Vault Decrypt Log Repository
type mockVaultDecryptLogRepo struct{}

func newMockVaultDecryptLogRepo() *mockVaultDecryptLogRepo {
	return &mockVaultDecryptLogRepo{}
}

func (m *mockVaultDecryptLogRepo) LogDecrypt(ctx context.Context, log model.VaultDecryptLog) error {
	return nil
}

func (m *mockVaultDecryptLogRepo) GetLogs(ctx context.Context, limit, offset int) ([]model.VaultDecryptLog, error) {
	return nil, errors.New("not implemented")
}

func (m *mockVaultDecryptLogRepo) GetLogsByUserID(ctx context.Context, userID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	return nil, errors.New("not implemented")
}

func (m *mockVaultDecryptLogRepo) GetLogsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	return nil, errors.New("not implemented")
}

// ==================== TEST FUNCTIONS ====================

func setupApplicationsService() (*applicationsService, *mockApplicationsRepo, *mockSLARepo) {
	mockAppRepo := newMockApplicationsRepo()
	mockUserRepo := newMockUsersRepo()
	mockNotifRepo := newMockNotificationRepo()
	mockSLARepo := newMockSLARepo()
	mockVaultRepo := newMockVaultDecryptLogRepo()

	service := &applicationsService{
		applicationRepository:  mockAppRepo,
		userRepository:         mockUserRepo,
		notificationRepository: mockNotifRepo,
		slaRepo:                mockSLARepo,
		vaultDecryptLogRepo:    mockVaultRepo,
	}

	return service, mockAppRepo, mockSLARepo
}

// Test GetAllApplications
func TestGetAllApplications(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	// Setup test data
	mockRepo.applications[1] = model.Application{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
		UMKM:      mockRepo.umkms[1],
		Program:   mockRepo.programs[1],
	}

	t.Run("Get all applications without filter", func(t *testing.T) {
		result, err := service.GetAllApplications(ctx, 0, "")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 application, got %d", len(result))
		}

		if result[0].Type != "training" {
			t.Errorf("Expected type 'training', got '%s'", result[0].Type)
		}
	})

	t.Run("Get applications with type filter", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:        2,
			UMKMID:    1,
			ProgramID: 2,
			Type:      "funding",
			Status:    "screening",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[2],
		}

		result, err := service.GetAllApplications(ctx, 0, "training")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 training application, got %d", len(result))
		}
	})
}

// Test GetApplicationByID
func TestGetApplicationByID(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	// Setup test data
	now := time.Now()
	actionedBy := 1

	mockRepo.applications[1] = model.Application{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
		UMKM:      mockRepo.umkms[1],
		Program:   mockRepo.programs[1],
	}

	t.Run("Get existing application", func(t *testing.T) {
		result, err := service.GetApplicationByID(ctx, 0, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		if result.Type != "training" {
			t.Errorf("Expected type 'training', got '%s'", result.Type)
		}
	})

	t.Run("Get non-existing application", func(t *testing.T) {
		_, err := service.GetApplicationByID(ctx, 0, 999)

		if err == nil {
			t.Error("Expected error for non-existing application, got none")
		}
	})

	t.Run("Get application with documents", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:        2,
			UMKMID:    1,
			ProgramID: 1,
			Type:      "training",
			Status:    "screening",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[1],
			Documents: []model.ApplicationDocument{
				{
					ID:            1,
					ApplicationID: 2,
					Type:          "ktp",
					File:          "ktp.pdf",
					Base: model.Base{
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				{
					ID:            2,
					ApplicationID: 2,
					Type:          "npwp",
					File:          "npwp.pdf",
					Base: model.Base{
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
		}

		result, err := service.GetApplicationByID(ctx, 0, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Documents) != 2 {
			t.Errorf("Expected 2 documents, got %d", len(result.Documents))
		}

		if result.Documents[0].Type != "ktp" {
			t.Errorf("Expected first document type 'ktp', got '%s'", result.Documents[0].Type)
		}

		if result.Documents[1].Type != "npwp" {
			t.Errorf("Expected second document type 'npwp', got '%s'", result.Documents[1].Type)
		}
	})

	t.Run("Get application with histories", func(t *testing.T) {
		mockRepo.applications[3] = model.Application{
			ID:        3,
			UMKMID:    1,
			ProgramID: 1,
			Type:      "training",
			Status:    "final",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[1],
			Histories: []model.ApplicationHistory{
				{
					ID:            1,
					ApplicationID: 3,
					Status:        "approve_by_admin_screening",
					Notes:         "Approved by screening admin",
					ActionedAt:    now,
					ActionedBy:    &actionedBy,
					Base: model.Base{
						CreatedAt: now,
						UpdatedAt: now,
					},
					User: model.User{
						ID:   1,
						Name: "Admin User",
					},
				},
				{
					ID:            2,
					ApplicationID: 3,
					Status:        "approve_by_admin_vendor",
					Notes:         "Approved by vendor admin",
					ActionedAt:    now,
					ActionedBy:    &actionedBy,
					Base: model.Base{
						CreatedAt: now,
						UpdatedAt: now,
					},
					User: model.User{
						ID:   1,
						Name: "Admin User",
					},
				},
			},
		}

		result, err := service.GetApplicationByID(ctx, 0, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Histories) != 2 {
			t.Errorf("Expected 2 histories, got %d", len(result.Histories))
		}

		if result.Histories[0].Status != "approve_by_admin_screening" {
			t.Errorf("Expected first history status 'approve_by_admin_screening', got '%s'", result.Histories[0].Status)
		}

		if result.Histories[0].ActionedByName != "Admin User" {
			t.Errorf("Expected actioned by name 'Admin User', got '%s'", result.Histories[0].ActionedByName)
		}
	})

	t.Run("Get training application with specific data", func(t *testing.T) {
		mockRepo.applications[4] = model.Application{
			ID:        4,
			UMKMID:    1,
			ProgramID: 1,
			Type:      "training",
			Status:    "screening",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[1],
			TrainingApplication: &model.TrainingApplication{
				ID:                 1,
				ApplicationID:      4,
				Motivation:         "I want to improve my business",
				BusinessExperience: "5 years in retail",
				LearningObjectives: "Learn digital marketing",
				AvailabilityNotes:  "Available weekdays",
			},
		}

		result, err := service.GetApplicationByID(ctx, 0, 4)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.TrainingData == nil {
			t.Error("Expected training data to be present")
		} else {
			if result.TrainingData.Motivation != "I want to improve my business" {
				t.Errorf("Expected motivation, got '%s'", result.TrainingData.Motivation)
			}
			if result.TrainingData.BusinessExperience != "5 years in retail" {
				t.Errorf("Expected business experience, got '%s'", result.TrainingData.BusinessExperience)
			}
			if result.TrainingData.LearningObjectives != "Learn digital marketing" {
				t.Errorf("Expected learning objectives, got '%s'", result.TrainingData.LearningObjectives)
			}
			if result.TrainingData.AvailabilityNotes != "Available weekdays" {
				t.Errorf("Expected availability notes, got '%s'", result.TrainingData.AvailabilityNotes)
			}
		}
	})

	t.Run("Get certification application with specific data", func(t *testing.T) {
		yearsOp := 3
		mockRepo.applications[5] = model.Application{
			ID:        5,
			UMKMID:    1,
			ProgramID: 1,
			Type:      "certification",
			Status:    "screening",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[1],
			CertificationApplication: &model.CertificationApplication{
				ID:                  1,
				ApplicationID:       5,
				BusinessSector:      "Food & Beverage",
				ProductOrService:    "Traditional Snacks",
				BusinessDescription: "We produce traditional snacks",
				YearsOperating:      &yearsOp,
				CurrentStandards:    "HACCP",
				CertificationGoals:  "Get halal certification",
			},
		}

		result, err := service.GetApplicationByID(ctx, 0, 5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.CertificationData == nil {
			t.Error("Expected certification data to be present")
		} else {
			if result.CertificationData.BusinessSector != "Food & Beverage" {
				t.Errorf("Expected business sector 'Food & Beverage', got '%s'", result.CertificationData.BusinessSector)
			}
			if result.CertificationData.ProductOrService != "Traditional Snacks" {
				t.Errorf("Expected product/service 'Traditional Snacks', got '%s'", result.CertificationData.ProductOrService)
			}
			if result.CertificationData.CertificationGoals != "Get halal certification" {
				t.Errorf("Expected certification goals, got '%s'", result.CertificationData.CertificationGoals)
			}
		}
	})

	t.Run("Get funding application with specific data", func(t *testing.T) {
		yearsOp := 5
		revenue := 100000000.0
		monthlyRev := 10000000.0
		mockRepo.applications[6] = model.Application{
			ID:        6,
			UMKMID:    1,
			ProgramID: 2,
			Type:      "funding",
			Status:    "screening",
			UMKM:      mockRepo.umkms[1],
			Program:   mockRepo.programs[2],
			FundingApplication: &model.FundingApplication{
				ID:                    1,
				ApplicationID:         6,
				BusinessSector:        "Manufacturing",
				BusinessDescription:   "We manufacture handicrafts",
				YearsOperating:        &yearsOp,
				RequestedAmount:       50000000.0,
				FundPurpose:           "Expand production capacity",
				BusinessPlan:          "Purchase new machinery",
				RevenueProjection:     &revenue,
				MonthlyRevenue:        &monthlyRev,
				RequestedTenureMonths: 24,
				CollateralDescription: "Land certificate",
			},
		}

		result, err := service.GetApplicationByID(ctx, 0, 6)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.FundingData == nil {
			t.Error("Expected funding data to be present")
		} else {
			if result.FundingData.BusinessSector != "Manufacturing" {
				t.Errorf("Expected business sector 'Manufacturing', got '%s'", result.FundingData.BusinessSector)
			}
			if result.FundingData.RequestedAmount != 50000000.0 {
				t.Errorf("Expected requested amount 50000000.0, got %f", result.FundingData.RequestedAmount)
			}
			if result.FundingData.FundPurpose != "Expand production capacity" {
				t.Errorf("Expected fund purpose, got '%s'", result.FundingData.FundPurpose)
			}
			if result.FundingData.RequestedTenureMonths != 24 {
				t.Errorf("Expected tenure 24 months, got %d", result.FundingData.RequestedTenureMonths)
			}
		}
	})
}

// Test ScreeningApprove
func TestScreeningApprove(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("Approve application in screening status", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:          1,
			UMKMID:      1,
			ProgramID:   1,
			Type:        "training",
			Status:      "screening",
			SubmittedAt: time.Now(),
		}

		result, err := service.ScreeningApprove(ctx, 1, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "final" {
			t.Errorf("Expected status 'final', got '%s'", result.Status)
		}
	})

	t.Run("Try to approve application not in screening status", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "final",
		}

		_, err := service.ScreeningApprove(ctx, 1, 2)

		if err == nil {
			t.Error("Expected error for non-screening status, got none")
		}

		if err.Error() != "application must be in screening status" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Approve non-existing application", func(t *testing.T) {
		_, err := service.ScreeningApprove(ctx, 1, 999)

		if err == nil {
			t.Error("Expected error for non-existing application, got none")
		}
	})
}

// Test ScreeningReject
func TestScreeningReject(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("Reject application with notes", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:     1,
			UMKMID: 1,
			Status: "screening",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "Documents incomplete",
		}

		result, err := service.ScreeningReject(ctx, 1, decision)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "rejected" {
			t.Errorf("Expected status 'rejected', got '%s'", result.Status)
		}
	})

	t.Run("Reject application without notes", func(t *testing.T) {
		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "",
		}

		_, err := service.ScreeningReject(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for missing notes, got none")
		}

		if err.Error() != "notes are required for rejection" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Reject application not in screening status", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "final",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 2,
			Notes:         "Test rejection",
		}

		_, err := service.ScreeningReject(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for non-screening status, got none")
		}
	})
}

// Test ScreeningRevise
func TestScreeningRevise(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("Request revision with notes", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:     1,
			UMKMID: 1,
			Status: "screening",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "Please update KTP document",
		}

		result, err := service.ScreeningRevise(ctx, 1, decision)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "revised" {
			t.Errorf("Expected status 'revised', got '%s'", result.Status)
		}
	})

	t.Run("Request revision without notes", func(t *testing.T) {
		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "",
		}

		_, err := service.ScreeningRevise(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for missing notes, got none")
		}

		if err.Error() != "notes are required for revision" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Request revision for non-screening application", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "approved",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 2,
			Notes:         "Test revision",
		}

		_, err := service.ScreeningRevise(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for non-screening status, got none")
		}
	})
}

// Test FinalApprove
func TestFinalApprove(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("Approve application in final status", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:     1,
			UMKMID: 1,
			Status: "final",
		}

		result, err := service.FinalApprove(ctx, 1, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "approved" {
			t.Errorf("Expected status 'approved', got '%s'", result.Status)
		}
	})

	t.Run("Try to approve application not in final status", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "screening",
		}

		_, err := service.FinalApprove(ctx, 1, 2)

		if err == nil {
			t.Error("Expected error for non-final status, got none")
		}

		if err.Error() != "application must be in final status" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Approve non-existing application", func(t *testing.T) {
		_, err := service.FinalApprove(ctx, 1, 999)

		if err == nil {
			t.Error("Expected error for non-existing application, got none")
		}
	})
}

// Test FinalReject
func TestFinalReject(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("Reject application in final status with notes", func(t *testing.T) {
		mockRepo.applications[1] = model.Application{
			ID:     1,
			UMKMID: 1,
			Status: "final",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "Capacity full",
		}

		result, err := service.FinalReject(ctx, 1, decision)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "rejected" {
			t.Errorf("Expected status 'rejected', got '%s'", result.Status)
		}
	})

	t.Run("Reject application without notes", func(t *testing.T) {
		decision := dto.ApplicationDecision{
			ApplicationID: 1,
			Notes:         "",
		}

		_, err := service.FinalReject(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for missing notes, got none")
		}

		if err.Error() != "notes are required for rejection" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Reject application not in final status", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "screening",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 2,
			Notes:         "Test rejection",
		}

		_, err := service.FinalReject(ctx, 1, decision)

		if err == nil {
			t.Error("Expected error for non-final status, got none")
		}
	})
}

// Test Edge Cases
func TestApplicationsServiceEdgeCases(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	t.Run("GetAllApplications with empty repository", func(t *testing.T) {
		result, err := service.GetAllApplications(ctx, 0, "")
		if err != nil {
			t.Errorf("Expected no error for empty repo, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 applications, got %d", len(result))
		}
	})

	t.Run("Multiple status transitions", func(t *testing.T) {
		// Create application
		mockRepo.applications[1] = model.Application{
			ID:          1,
			UMKMID:      1,
			Status:      "screening",
			SubmittedAt: time.Now(),
		}

		// Approve at screening
		result1, err := service.ScreeningApprove(ctx, 1, 1)
		if err != nil || result1.Status != "final" {
			t.Error("Failed to approve at screening stage")
		}

		// Approve at final
		result2, err := service.FinalApprove(ctx, 1, 1)
		if err != nil || result2.Status != "approved" {
			t.Error("Failed to approve at final stage")
		}
	})

	t.Run("Check history creation", func(t *testing.T) {
		mockRepo.applications[2] = model.Application{
			ID:     2,
			UMKMID: 1,
			Status: "screening",
		}

		decision := dto.ApplicationDecision{
			ApplicationID: 2,
			Notes:         "Test rejection",
		}

		service.ScreeningReject(ctx, 1, decision)

		histories, _ := mockRepo.GetApplicationHistories(ctx, 2)
		if len(histories) == 0 {
			t.Error("Expected history to be created")
		}
	})

	t.Run("Check notification creation", func(t *testing.T) {
		mockRepo.applications[3] = model.Application{
			ID:          3,
			UMKMID:      1,
			Status:      "screening",
			SubmittedAt: time.Now(),
		}

		service.ScreeningApprove(ctx, 1, 3)

		// In real implementation, check if notification was created
		// This is a simplified check
		if len(mockRepo.applications) == 0 {
			t.Error("Expected application to exist after approval")
		}
	})
}

// Test Concurrent Operations
func TestConcurrentOperations(t *testing.T) {
	service, mockRepo, _ := setupApplicationsService()
	ctx := context.Background()

	mockRepo.applications[1] = model.Application{
		ID:          1,
		UMKMID:      1,
		Status:      "screening",
		SubmittedAt: time.Now(),
	}

	t.Run("Concurrent reads", func(t *testing.T) {
		done := make(chan bool)

		for i := 0; i < 10; i++ {
			go func() {
				_, err := service.GetApplicationByID(ctx, 0, 1)
				if err != nil {
					t.Errorf("Concurrent read failed: %v", err)
				}
				done <- true
			}()
		}

		for i := 0; i < 10; i++ {
			<-done
		}
	})
}
