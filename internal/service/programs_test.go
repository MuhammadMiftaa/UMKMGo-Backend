package service

import (
	"errors"
	"testing"
	"time"

	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/types/model"
)

// Mock programs repository
type mockProgramsRepository struct {
	programs     map[int]model.Programs
	benefits     map[int][]model.ProgramBenefits
	requirements map[int][]model.ProgramRequirements
}

func newMockProgramsRepository() *mockProgramsRepository {
	return &mockProgramsRepository{
		programs:     make(map[int]model.Programs),
		benefits:     make(map[int][]model.ProgramBenefits),
		requirements: make(map[int][]model.ProgramRequirements),
	}
}

func (m *mockProgramsRepository) GetAllPrograms() ([]model.Programs, error) {
	programs := []model.Programs{}
	for _, p := range m.programs {
		programs = append(programs, p)
	}
	return programs, nil
}

func (m *mockProgramsRepository) GetProgramByID(id int) (model.Programs, error) {
	if program, exists := m.programs[id]; exists {
		return program, nil
	}
	return model.Programs{}, errors.New("program not found")
}

func (m *mockProgramsRepository) CreateProgram(program model.Programs) (model.Programs, error) {
	program.ID = len(m.programs) + 1
	program.CreatedAt = time.Now()
	program.UpdatedAt = time.Now()
	m.programs[program.ID] = program
	return program, nil
}

func (m *mockProgramsRepository) UpdateProgram(program model.Programs) (model.Programs, error) {
	if _, exists := m.programs[program.ID]; !exists {
		return model.Programs{}, errors.New("program not found")
	}
	program.UpdatedAt = time.Now()
	m.programs[program.ID] = program
	return program, nil
}

func (m *mockProgramsRepository) DeleteProgram(program model.Programs) (model.Programs, error) {
	delete(m.programs, program.ID)
	return program, nil
}

func (m *mockProgramsRepository) CreateProgramBenefits(benefits []model.ProgramBenefits) error {
	if len(benefits) == 0 {
		return nil
	}
	programID := benefits[0].ProgramID
	m.benefits[programID] = benefits
	return nil
}

func (m *mockProgramsRepository) CreateProgramRequirements(requirements []model.ProgramRequirements) error {
	if len(requirements) == 0 {
		return nil
	}
	programID := requirements[0].ProgramID
	m.requirements[programID] = requirements
	return nil
}

func (m *mockProgramsRepository) GetProgramBenefits(programID int) ([]model.ProgramBenefits, error) {
	if benefits, exists := m.benefits[programID]; exists {
		return benefits, nil
	}
	return []model.ProgramBenefits{}, nil
}

func (m *mockProgramsRepository) GetProgramRequirements(programID int) ([]model.ProgramRequirements, error) {
	if requirements, exists := m.requirements[programID]; exists {
		return requirements, nil
	}
	return []model.ProgramRequirements{}, nil
}

func (m *mockProgramsRepository) DeleteProgramBenefits(programID int) error {
	delete(m.benefits, programID)
	return nil
}

func (m *mockProgramsRepository) DeleteProgramRequirements(programID int) error {
	delete(m.requirements, programID)
	return nil
}

// Mock user repository for programs service
type mockUserRepositoryForPrograms struct {
	users map[int]model.Users
}

func newMockUserRepositoryForPrograms() *mockUserRepositoryForPrograms {
	return &mockUserRepositoryForPrograms{
		users: map[int]model.Users{
			1: {ID: 1, Name: "Admin User", Email: "admin@example.com"},
		},
	}
}

func (m *mockUserRepositoryForPrograms) GetUserByID(id int) (model.Users, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return model.Users{}, errors.New("user not found")
}

func (m *mockUserRepositoryForPrograms) GetAllUsers() ([]model.Users, error) {
	return nil, nil
}

func (m *mockUserRepositoryForPrograms) GetUserByEmail(email string) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForPrograms) CreateUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForPrograms) UpdateUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForPrograms) DeleteUser(user model.Users) (model.Users, error) {
	return model.Users{}, nil
}

func (m *mockUserRepositoryForPrograms) GetAllRoles() ([]model.Roles, error) {
	return nil, nil
}

func (m *mockUserRepositoryForPrograms) GetRoleByID(id int) (model.Roles, error) {
	return model.Roles{}, nil
}

func (m *mockUserRepositoryForPrograms) IsRoleExist(id int) bool {
	return false
}

func (m *mockUserRepositoryForPrograms) IsPermissionExist(ids []int) bool {
	return false
}

func (m *mockUserRepositoryForPrograms) GetListPermissions() ([]model.Permissions, error) {
	return nil, nil
}

func (m *mockUserRepositoryForPrograms) GetListRolePermissions() ([]model.RolePermissionsResponse, error) {
	return nil, nil
}

func (m *mockUserRepositoryForPrograms) DeletePermissionsByRoleID(roleID int) error {
	return nil
}

func (m *mockUserRepositoryForPrograms) AddRolePermissions(roleID int, permissions []int) error {
	return nil
}

// Test CreateProgram
func TestCreateProgram(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	trainingType := "online"
	tests := []struct {
		name        string
		input       dto.Programs
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid training program",
			input: dto.Programs{
				Title:               "Digital Marketing Training",
				Description:         "Learn digital marketing basics",
				Type:                "training",
				TrainingType:        &trainingType,
				ApplicationDeadline: "2025-12-31",
				CreatedBy:           1,
				Benefits:            []string{"Certificate", "Mentoring"},
				Requirements:        []string{"Basic computer skills"},
			},
			expectError: false,
		},
		{
			name: "Missing required fields",
			input: dto.Programs{
				Description: "Some description",
			},
			expectError: true,
			errorMsg:    "title, type, and application deadline are required",
		},
		{
			name: "Invalid type",
			input: dto.Programs{
				Title:               "Test Program",
				Type:                "invalid_type",
				ApplicationDeadline: "2025-12-31",
			},
			expectError: true,
			errorMsg:    "type must be training, certification, or funding",
		},
		{
			name: "Invalid training type",
			input: dto.Programs{
				Title:               "Test Training",
				Type:                "training",
				TrainingType:        strPtr("invalid"),
				ApplicationDeadline: "2025-12-31",
			},
			expectError: true,
			errorMsg:    "training type must be online, offline, or hybrid",
		},
		{
			name: "Invalid creator user",
			input: dto.Programs{
				Title:               "Test Program",
				Type:                "training",
				ApplicationDeadline: "2025-12-31",
				CreatedBy:           999,
			},
			expectError: true,
			errorMsg:    "creator user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.CreateProgram(tt.input)

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
				if result.Title != tt.input.Title {
					t.Errorf("Expected title %s, got %s", tt.input.Title, result.Title)
				}
				if result.Type != tt.input.Type {
					t.Errorf("Expected type %s, got %s", tt.input.Type, result.Type)
				}
			}
		})
	}
}

// Test GetAllPrograms
func TestGetAllPrograms(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test programs
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Program 1",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
		CreatedBy:           1,
		Users:               model.Users{ID: 1, Name: "Admin User"},
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result, err := service.GetAllPrograms()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 program, got %d", len(result))
	}

	if result[0].Title != "Program 1" {
		t.Errorf("Expected title 'Program 1', got '%s'", result[0].Title)
	}
}

// Test GetProgramByID
func TestGetProgramByID(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test program
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
		Users:               model.Users{ID: 1, Name: "Admin User"},
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	tests := []struct {
		name        string
		programID   int
		expectError bool
	}{
		{
			name:        "Valid program ID",
			programID:   1,
			expectError: false,
		},
		{
			name:        "Invalid program ID",
			programID:   999,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetProgramByID(tt.programID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.ID != tt.programID {
					t.Errorf("Expected program ID %d, got %d", tt.programID, result.ID)
				}
			}
		})
	}
}

// Test UpdateProgram
func TestUpdateProgram(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test program
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Original Title",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	tests := []struct {
		name        string
		programID   int
		input       dto.Programs
		expectError bool
		errorMsg    string
	}{
		{
			name:      "Valid update",
			programID: 1,
			input: dto.Programs{
				Title:               "Updated Title",
				Type:                "certification",
				ApplicationDeadline: "2026-01-31",
			},
			expectError: false,
		},
		{
			name:      "Invalid program ID",
			programID: 999,
			input: dto.Programs{
				Title:               "Updated Title",
				Type:                "training",
				ApplicationDeadline: "2026-01-31",
			},
			expectError: true,
		},
		{
			name:      "Missing required fields",
			programID: 1,
			input: dto.Programs{
				Description: "Only description",
			},
			expectError: true,
			errorMsg:    "title, type, and application deadline are required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.UpdateProgram(tt.programID, tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Title != tt.input.Title {
					t.Errorf("Expected title %s, got %s", tt.input.Title, result.Title)
				}
			}
		})
	}
}

// Test DeleteProgram
func TestDeleteProgram(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test program
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
	}

	tests := []struct {
		name        string
		programID   int
		expectError bool
	}{
		{
			name:        "Valid deletion",
			programID:   1,
			expectError: false,
		},
		{
			name:        "Invalid program ID",
			programID:   999,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.DeleteProgram(tt.programID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.ID != tt.programID {
					t.Errorf("Expected deleted program ID %d, got %d", tt.programID, result.ID)
				}
			}
		})
	}
}

// Test ActivateProgram
func TestActivateProgram(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test program (inactive)
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            false,
	}

	result, err := service.ActivateProgram(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !result.IsActive {
		t.Errorf("Expected program to be active, but it's not")
	}

	// Test with invalid ID
	_, err = service.ActivateProgram(999)
	if err == nil {
		t.Errorf("Expected error for invalid program ID, but got none")
	}
}

// Test DeactivateProgram
func TestDeactivateProgram(t *testing.T) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUserRepositoryForPrograms()
	service := NewProgramsService(mockProgramRepo, mockUserRepo)

	// Add test program (active)
	mockProgramRepo.programs[1] = model.Programs{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
	}

	result, err := service.DeactivateProgram(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result.IsActive {
		t.Errorf("Expected program to be inactive, but it's active")
	}

	// Test with invalid ID
	_, err = service.DeactivateProgram(999)
	if err == nil {
		t.Errorf("Expected error for invalid program ID, but got none")
	}
}

// Helper function
func strPtr(s string) *string {
	return &s
}
