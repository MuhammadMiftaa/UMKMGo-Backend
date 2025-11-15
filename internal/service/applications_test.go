package service

import (
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// Mock applications repository
type mockApplicationsRepository struct {
	applications map[int]model.Applications
	documents    map[int][]model.ApplicationDocuments
	histories    map[int][]model.ApplicationHistories
	programs     map[int]model.Programs
	umkms        map[int]model.UMKMS
}

func newMockApplicationsRepository() *mockApplicationsRepository {
	return &mockApplicationsRepository{
		applications: make(map[int]model.Applications),
		documents:    make(map[int][]model.ApplicationDocuments),
		histories:    make(map[int][]model.ApplicationHistories),
		programs: map[int]model.Programs{
			1: {ID: 1, Title: "Test Program", Type: "training", IsActive: true},
			2: {ID: 2, Title: "Inactive Program", Type: "training", IsActive: false},
		},
		umkms: map[int]model.UMKMS{
			1: {ID: 1, UserID: 1, BusinessName: "Test Business", NIK: "1234567890"},
		},
	}
}

func (m *mockApplicationsRepository) GetAllApplications(filterType string) ([]model.Applications, error) {
	var applications []model.Applications
	for _, app := range m.applications {
		if filterType == "" || app.Type == filterType {
			applications = append(applications, app)
		}
	}
	return applications, nil
}

func (m *mockApplicationsRepository) GetApplicationByID(id int) (model.Applications, error) {
	if app, exists := m.applications[id]; exists {
		return app, nil
	}
	return model.Applications{}, errors.New("application not found")
}

func (m *mockApplicationsRepository) CreateApplication(application model.Applications) (model.Applications, error) {
	application.ID = len(m.applications) + 1
	application.SubmittedAt = time.Now()
	application.ExpiredAt = time.Now().AddDate(0, 0, 30)
	m.applications[application.ID] = application
	return application, nil
}

func (m *mockApplicationsRepository) UpdateApplication(application model.Applications) (model.Applications, error) {
	if _, exists := m.applications[application.ID]; !exists {
		return model.Applications{}, errors.New("application not found")
	}
	m.applications[application.ID] = application
	return application, nil
}

func (m *mockApplicationsRepository) DeleteApplication(application model.Applications) (model.Applications, error) {
	delete(m.applications, application.ID)
	return application, nil
}

func (m *mockApplicationsRepository) CreateApplicationDocuments(documents []model.ApplicationDocuments) error {
	if len(documents) == 0 {
		return nil
	}
	appID := documents[0].ApplicationID
	m.documents[appID] = documents
	return nil
}

func (m *mockApplicationsRepository) GetApplicationDocuments(applicationID int) ([]model.ApplicationDocuments, error) {
	if docs, exists := m.documents[applicationID]; exists {
		return docs, nil
	}
	return []model.ApplicationDocuments{}, nil
}

func (m *mockApplicationsRepository) DeleteApplicationDocuments(applicationID int) error {
	delete(m.documents, applicationID)
	return nil
}

func (m *mockApplicationsRepository) CreateApplicationHistory(history model.ApplicationHistories) error {
	m.histories[history.ApplicationID] = append(m.histories[history.ApplicationID], history)
	return nil
}

func (m *mockApplicationsRepository) GetApplicationHistories(applicationID int) ([]model.ApplicationHistories, error) {
	if histories, exists := m.histories[applicationID]; exists {
		return histories, nil
	}
	return []model.ApplicationHistories{}, nil
}

func (m *mockApplicationsRepository) GetProgramByID(id int) (model.Programs, error) {
	if program, exists := m.programs[id]; exists {
		return program, nil
	}
	return model.Programs{}, errors.New("program not found")
}

func (m *mockApplicationsRepository) GetUMKMByUserID(userID int) (model.UMKMS, error) {
	for _, umkm := range m.umkms {
		if umkm.UserID == userID {
			return umkm, nil
		}
	}
	return model.UMKMS{}, errors.New("UMKM not found")
}

func (m *mockApplicationsRepository) IsApplicationExists(umkmID, programID int) bool {
	for _, app := range m.applications {
		if app.UMKMID == umkmID && app.ProgramID == programID && app.Status != "rejected" {
			return true
		}
	}
	return false
}

// Mock user repository for applications service
type mockUserRepositoryForApplications struct {
	users map[int]model.Users
}

func newMockUserRepositoryForApplications() *mockUserRepositoryForApplications {
	return &mockUserRepositoryForApplications{
		users: map[int]model.Users{
			1: {ID: 1, Name: "Admin User", Email: "admin@example.com"},
			2: {ID: 2, Name: "UMKM User", Email: "umkm@example.com"},
		},
	}
}

func (m *mockUserRepositoryForApplications) GetUserByID(id int) (model.Users, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return model.Users{}, errors.New("user not found")
}

func (m *mockUserRepositoryForApplications) GetAllUsers() ([]model.Users, error) {
	return nil, nil
}

func (m *mockUserRepositoryForApplications) GetUserByEmail(email string) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForApplications) CreateUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForApplications) UpdateUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForApplications) DeleteUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForApplications) GetAllRoles() ([]model.Roles, error) {
	return nil, nil
}

func (m *mockUserRepositoryForApplications) GetRoleByID(id int) (model.Roles, error) {
	return model.Roles{}, nil
}

func (m *mockUserRepositoryForApplications) IsRoleExist(id int) bool {
	return false
}

func (m *mockUserRepositoryForApplications) IsPermissionExist(id []string) ([]int, bool) {
	return nil, false
}

func (m *mockUserRepositoryForApplications) GetListPermissions() ([]model.Permissions, error) {
	return nil, nil
}

func (m *mockUserRepositoryForApplications) GetListPermissionsByRoleID(roleID int) ([]string, error) {
	return nil, nil
}

func (m *mockUserRepositoryForApplications) GetListRolePermissions() ([]model.RolePermissionsResponse, error) {
	return nil, nil
}

func (m *mockUserRepositoryForApplications) DeletePermissionsByRoleID(roleID int) error {
	return nil
}

func (m *mockUserRepositoryForApplications) AddRolePermissions(roleID int, permissions []int) error {
	return nil
}

// Test CreateApplication
// func TestCreateApplication(t *testing.T) {
// 	mockAppRepo := newMockApplicationsRepository()
// 	mockUserRepo := newMockUserRepositoryForApplications()
// 	service := NewApplicationsService(mockAppRepo, mockUserRepo)

// 	tests := []struct {
// 		name        string
// 		userID      int
// 		input       dto.Applications
// 		expectError bool
// 		errorMsg    string
// 	}{
// 		{
// 			name:   "Valid application",
// 			userID: 1,
// 			input: dto.Applications{
// 				ProgramID: 1,
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 					{Type: "nib", File: "nib.pdf"},
// 				},
// 			},
// 			expectError: false,
// 		},
// 		{
// 			name:   "Missing program ID",
// 			userID: 1,
// 			input: dto.Applications{
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 			errorMsg:    "program_id and documents are required",
// 		},
// 		{
// 			name:   "Missing documents",
// 			userID: 1,
// 			input: dto.Applications{
// 				ProgramID: 1,
// 			},
// 			expectError: true,
// 			errorMsg:    "program_id and documents are required",
// 		},
// 		{
// 			name:   "Program not found",
// 			userID: 1,
// 			input: dto.Applications{
// 				ProgramID: 999,
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 			errorMsg:    "program not found",
// 		},
// 		{
// 			name:   "Program not active",
// 			userID: 1,
// 			input: dto.Applications{
// 				ProgramID: 2,
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 			errorMsg:    "program is not active",
// 		},
// 		{
// 			name:   "UMKM not found",
// 			userID: 999,
// 			input: dto.Applications{
// 				ProgramID: 1,
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 			errorMsg:    "UMKM data not found, please complete your profile first",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result, err := service.CreateApplication(tt.userID, tt.input)

// 			if tt.expectError {
// 				if err == nil {
// 					t.Errorf("Expected error but got none")
// 				} else if err.Error() != tt.errorMsg {
// 					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("Unexpected error: %v", err)
// 				}
// 				if result.ProgramID != tt.input.ProgramID {
// 					t.Errorf("Expected program ID %d, got %d", tt.input.ProgramID, result.ProgramID)
// 				}
// 				if result.Status != "screening" {
// 					t.Errorf("Expected status 'screening', got '%s'", result.Status)
// 				}
// 			}
// 		})
// 	}
// }

// Test GetAllApplications
func TestGetAllApplications(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test applications
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Program: mockAppRepo.programs[1],
		UMKM:    mockAppRepo.umkms[1],
	}

	result, err := service.GetAllApplications("")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 application, got %d", len(result))
	}

	if result[0].Type != "training" {
		t.Errorf("Expected type 'training', got '%s'", result[0].Type)
	}
}

// Test GetApplicationByID
func TestGetApplicationByID(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Program: mockAppRepo.programs[1],
		UMKM:    mockAppRepo.umkms[1],
	}

	tests := []struct {
		name          string
		applicationID int
		expectError   bool
	}{
		{
			name:          "Valid application ID",
			applicationID: 1,
			expectError:   false,
		},
		{
			name:          "Invalid application ID",
			applicationID: 999,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetApplicationByID(tt.applicationID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.ID != tt.applicationID {
					t.Errorf("Expected application ID %d, got %d", tt.applicationID, result.ID)
				}
			}
		})
	}
}

// Test UpdateApplication
// func TestUpdateApplication(t *testing.T) {
// 	mockAppRepo := newMockApplicationsRepository()
// 	mockUserRepo := newMockUserRepositoryForApplications()
// 	service := NewApplicationsService(mockAppRepo, mockUserRepo)

// 	// Add test application in revised status
// 	mockAppRepo.applications[1] = model.Applications{
// 		ID:        1,
// 		UMKMID:    1,
// 		ProgramID: 1,
// 		Type:      "training",
// 		Status:    "revised",
// 		Base: model.Base{
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		UMKM: mockAppRepo.umkms[1],
// 	}

// 	// Add test application in screening status
// 	mockAppRepo.applications[2] = model.Applications{
// 		ID:        2,
// 		UMKMID:    1,
// 		ProgramID: 1,
// 		Type:      "training",
// 		Status:    "screening",
// 		Base: model.Base{
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		UMKM: mockAppRepo.umkms[1],
// 	}

// 	tests := []struct {
// 		name          string
// 		applicationID int
// 		input         dto.Applications
// 		expectError   bool
// 		errorMsg      string
// 	}{
// 		{
// 			name:          "Valid update",
// 			applicationID: 1,
// 			input: dto.Applications{
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp-updated.pdf"},
// 				},
// 			},
// 			expectError: false,
// 		},
// 		{
// 			name:          "Missing documents",
// 			applicationID: 1,
// 			input:         dto.Applications{},
// 			expectError:   true,
// 			errorMsg:      "documents are required",
// 		},
// 		{
// 			name:          "Wrong status",
// 			applicationID: 2,
// 			input: dto.Applications{
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 			errorMsg:    "only applications with status 'revised' can be updated",
// 		},
// 		{
// 			name:          "Application not found",
// 			applicationID: 999,
// 			input: dto.Applications{
// 				Documents: []dto.ApplicationDocuments{
// 					{Type: "ktp", File: "ktp.pdf"},
// 				},
// 			},
// 			expectError: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result, err := service.UpdateApplication(tt.applicationID, tt.input)

// 			if tt.expectError {
// 				if err == nil {
// 					t.Errorf("Expected error but got none")
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("Unexpected error: %v", err)
// 				}
// 				if result.Status != "screening" {
// 					t.Errorf("Expected status 'screening', got '%s'", result.Status)
// 				}
// 			}
// 		})
// 	}
// }

// Test ScreeningApprove
func TestScreeningApprove(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application in screening status
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
	}

	// Add test application in final status
	mockAppRepo.applications[2] = model.Applications{
		ID:        2,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "final",
	}

	tests := []struct {
		name          string
		userID        int
		applicationID int
		expectError   bool
		errorMsg      string
	}{
		{
			name:          "Valid approve",
			userID:        1,
			applicationID: 1,
			expectError:   false,
		},
		{
			name:          "Wrong status",
			userID:        1,
			applicationID: 2,
			expectError:   true,
			errorMsg:      "application must be in screening status",
		},
		{
			name:          "Application not found",
			userID:        1,
			applicationID: 999,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ScreeningApprove(tt.userID, tt.applicationID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Status != "final" {
					t.Errorf("Expected status 'final', got '%s'", result.Status)
				}
			}
		})
	}
}

// Test ScreeningReject
func TestScreeningReject(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application in screening status
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
	}

	tests := []struct {
		name        string
		userID      int
		decision    dto.ApplicationDecision
		expectError bool
		errorMsg    string
	}{
		{
			name:   "Valid reject",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
				Notes:         "Documents incomplete",
			},
			expectError: false,
		},
		{
			name:   "Missing notes",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
			},
			expectError: true,
			errorMsg:    "notes are required for rejection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ScreeningReject(tt.userID, tt.decision)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Status != "rejected" {
					t.Errorf("Expected status 'rejected', got '%s'", result.Status)
				}
			}
		})
	}
}

// Test ScreeningRevise
func TestScreeningRevise(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application in screening status
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
	}

	tests := []struct {
		name        string
		userID      int
		decision    dto.ApplicationDecision
		expectError bool
		errorMsg    string
	}{
		{
			name:   "Valid revise",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
				Notes:         "Please update KTP document",
			},
			expectError: false,
		},
		{
			name:   "Missing notes",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
			},
			expectError: true,
			errorMsg:    "notes are required for revision",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ScreeningRevise(tt.userID, tt.decision)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Status != "revised" {
					t.Errorf("Expected status 'revised', got '%s'", result.Status)
				}
			}
		})
	}
}

// Test FinalApprove
func TestFinalApprove(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application in final status
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "final",
	}

	// Add test application in screening status
	mockAppRepo.applications[2] = model.Applications{
		ID:        2,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "screening",
	}

	tests := []struct {
		name          string
		userID        int
		applicationID int
		expectError   bool
		errorMsg      string
	}{
		{
			name:          "Valid approve",
			userID:        1,
			applicationID: 1,
			expectError:   false,
		},
		{
			name:          "Wrong status",
			userID:        1,
			applicationID: 2,
			expectError:   true,
			errorMsg:      "application must be in final status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.FinalApprove(tt.userID, tt.applicationID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Status != "approved" {
					t.Errorf("Expected status 'approved', got '%s'", result.Status)
				}
			}
		})
	}
}

// Test FinalReject
func TestFinalReject(t *testing.T) {
	mockAppRepo := newMockApplicationsRepository()
	mockUserRepo := newMockUserRepositoryForApplications()
	service := NewApplicationsService(mockAppRepo, mockUserRepo)

	// Add test application in final status
	mockAppRepo.applications[1] = model.Applications{
		ID:        1,
		UMKMID:    1,
		ProgramID: 1,
		Type:      "training",
		Status:    "final",
	}

	tests := []struct {
		name        string
		userID      int
		decision    dto.ApplicationDecision
		expectError bool
		errorMsg    string
	}{
		{
			name:   "Valid reject",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
				Notes:         "Program capacity full",
			},
			expectError: false,
		},
		{
			name:   "Missing notes",
			userID: 1,
			decision: dto.ApplicationDecision{
				ApplicationID: 1,
			},
			expectError: true,
			errorMsg:    "notes are required for rejection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.FinalReject(tt.userID, tt.decision)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Status != "rejected" {
					t.Errorf("Expected status 'rejected', got '%s'", result.Status)
				}
			}
		})
	}
}

// Test DeleteApplication
// func TestDeleteApplication(t *testing.T) {
// 	mockAppRepo := newMockApplicationsRepository()
// 	mockUserRepo := newMockUserRepositoryForApplications()
// 	service := NewApplicationsService(mockAppRepo, mockUserRepo)

// 	// Add test application
// 	mockAppRepo.applications[1] = model.Applications{
// 		ID:        1,
// 		UMKMID:    1,
// 		ProgramID: 1,
// 		Type:      "training",
// 		Status:    "screening",
// 	}

// 	tests := []struct {
// 		name          string
// 		applicationID int
// 		expectError   bool
// 	}{
// 		{
// 			name:          "Valid deletion",
// 			applicationID: 1,
// 			expectError:   false,
// 		},
// 		{
// 			name:          "Application not found",
// 			applicationID: 999,
// 			expectError:   true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result, err := service.DeleteApplication(tt.applicationID)

// 			if tt.expectError {
// 				if err == nil {
// 					t.Errorf("Expected error but got none")
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("Unexpected error: %v", err)
// 				}
// 				if result.ID != tt.applicationID {
// 					t.Errorf("Expected deleted application ID %d, got %d", tt.applicationID, result.ID)
// 				}
// 			}
// 		})
// 	}
// }
