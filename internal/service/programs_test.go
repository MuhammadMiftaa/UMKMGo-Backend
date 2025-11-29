package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// ==================== MOCK REPOSITORIES FOR PROGRAMS ====================

// Mock Programs Repository
type mockProgramsRepository struct {
	programs     map[int]model.Program
	benefits     map[int][]model.ProgramBenefit
	requirements map[int][]model.ProgramRequirement
}

func newMockProgramsRepository() *mockProgramsRepository {
	return &mockProgramsRepository{
		programs:     make(map[int]model.Program),
		benefits:     make(map[int][]model.ProgramBenefit),
		requirements: make(map[int][]model.ProgramRequirement),
	}
}

func (m *mockProgramsRepository) GetAllPrograms(ctx context.Context) ([]model.Program, error) {
	var programs []model.Program
	for _, p := range m.programs {
		programs = append(programs, p)
	}
	return programs, nil
}

func (m *mockProgramsRepository) GetProgramByID(ctx context.Context, id int) (model.Program, error) {
	if program, exists := m.programs[id]; exists {
		return program, nil
	}
	return model.Program{}, errors.New("program not found")
}

func (m *mockProgramsRepository) CreateProgram(ctx context.Context, program model.Program) (model.Program, error) {
	program.ID = len(m.programs) + 1
	program.CreatedAt = time.Now()
	program.UpdatedAt = time.Now()
	m.programs[program.ID] = program
	return program, nil
}

func (m *mockProgramsRepository) UpdateProgram(ctx context.Context, program model.Program) (model.Program, error) {
	if _, exists := m.programs[program.ID]; !exists {
		return model.Program{}, errors.New("program not found")
	}
	program.UpdatedAt = time.Now()
	m.programs[program.ID] = program
	return program, nil
}

func (m *mockProgramsRepository) DeleteProgram(ctx context.Context, program model.Program) (model.Program, error) {
	if _, exists := m.programs[program.ID]; !exists {
		return model.Program{}, errors.New("program not found")
	}
	delete(m.programs, program.ID)
	return program, nil
}

func (m *mockProgramsRepository) CreateProgramBenefits(ctx context.Context, benefits []model.ProgramBenefit) error {
	if len(benefits) == 0 {
		return nil
	}
	programID := benefits[0].ProgramID
	m.benefits[programID] = benefits
	return nil
}

func (m *mockProgramsRepository) CreateProgramRequirements(ctx context.Context, requirements []model.ProgramRequirement) error {
	if len(requirements) == 0 {
		return nil
	}
	programID := requirements[0].ProgramID
	m.requirements[programID] = requirements
	return nil
}

func (m *mockProgramsRepository) GetProgramBenefits(ctx context.Context, programID int) ([]model.ProgramBenefit, error) {
	if benefits, exists := m.benefits[programID]; exists {
		return benefits, nil
	}
	return []model.ProgramBenefit{}, nil
}

func (m *mockProgramsRepository) GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirement, error) {
	if requirements, exists := m.requirements[programID]; exists {
		return requirements, nil
	}
	return []model.ProgramRequirement{}, nil
}

func (m *mockProgramsRepository) DeleteProgramBenefits(ctx context.Context, programID int) error {
	delete(m.benefits, programID)
	return nil
}

func (m *mockProgramsRepository) DeleteProgramRequirements(ctx context.Context, programID int) error {
	delete(m.requirements, programID)
	return nil
}

// Mock Users Repository for Programs
type mockUsersRepositoryForPrograms struct {
	users map[int]model.User
}

func newMockUsersRepositoryForPrograms() *mockUsersRepositoryForPrograms {
	return &mockUsersRepositoryForPrograms{
		users: map[int]model.User{
			1: {ID: 1, Name: "Admin User", Email: "admin@example.com"},
		},
	}
}

func (m *mockUsersRepositoryForPrograms) GetUserByID(ctx context.Context, id int) (model.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return model.User{}, errors.New("user not found")
}

func (m *mockUsersRepositoryForPrograms) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) DeleteUser(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User) (dto.UMKMMobile, error) {
	return dto.UMKMMobile{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error) {
	return model.UMKM{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetRoleByID(ctx context.Context, id int) (model.Role, error) {
	return model.Role{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	return model.Role{}, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) IsRoleExist(ctx context.Context, id int) bool {
	return false
}

func (m *mockUsersRepositoryForPrograms) IsPermissionExist(ctx context.Context, ids []string) ([]int, bool) {
	return nil, false
}

func (m *mockUsersRepositoryForPrograms) GetListPermissions(ctx context.Context) ([]model.Permission, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) DeletePermissionsByRoleID(ctx context.Context, roleID int) error {
	return errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) AddRolePermissions(ctx context.Context, roleID int, permissions []int) error {
	return errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetProvinces(ctx context.Context) ([]dto.Province, error) {
	return nil, errors.New("not implemented")
}

func (m *mockUsersRepositoryForPrograms) GetCities(ctx context.Context) ([]dto.City, error) {
	return nil, errors.New("not implemented")
}

// ==================== TEST FUNCTIONS FOR PROGRAMS ====================

func setupProgramsService() (*programsService, *mockProgramsRepository) {
	mockProgramRepo := newMockProgramsRepository()
	mockUserRepo := newMockUsersRepositoryForPrograms()

	service := &programsService{
		programRepository: mockProgramRepo,
		userRepository:    mockUserRepo,
		redisRepository:   redis.RedisRepository{}, // Not needed for these tests
		minio:             nil,                     // Not needed for these tests
	}

	return service, mockProgramRepo
}

// Test GetAllPrograms
func TestGetAllPrograms(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	t.Run("Get all programs when empty", func(t *testing.T) {
		result, err := service.GetAllPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 programs, got %d", len(result))
		}
	})

	t.Run("Get all programs with data", func(t *testing.T) {
		mockRepo.programs[1] = model.Program{
			ID:                  1,
			Title:               "Training Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
			IsActive:            true,
			CreatedBy:           1,
			Users:               model.User{ID: 1, Name: "Admin User"},
			Base: model.Base{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		result, err := service.GetAllPrograms(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 program, got %d", len(result))
		}

		if result[0].Title != "Training Program" {
			t.Errorf("Expected title 'Training Program', got '%s'", result[0].Title)
		}
	})
}

// Test GetProgramByID
func TestGetProgramByID(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	mockRepo.programs[1] = model.Program{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
		Users:               model.User{ID: 1, Name: "Admin User"},
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("Get existing program", func(t *testing.T) {
		result, err := service.GetProgramByID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
	})

	t.Run("Get non-existing program", func(t *testing.T) {
		_, err := service.GetProgramByID(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})
}

// Test CreateProgram
func TestCreateProgram(t *testing.T) {
	service, _ := setupProgramsService()
	ctx := context.Background()

	trainingType := "online"

	t.Run("Create valid training program", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Digital Marketing Training",
			Description:         "Learn digital marketing basics",
			Type:                "training",
			TrainingType:        &trainingType,
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
			Benefits:            []string{"Certificate", "Mentoring"},
			Requirements:        []string{"Basic computer skills"},
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Title != input.Title {
			t.Errorf("Expected title '%s', got '%s'", input.Title, result.Title)
		}
	})

	t.Run("Create program with missing required fields", func(t *testing.T) {
		input := dto.Programs{
			Description: "Some description",
		}

		_, err := service.CreateProgram(ctx, input)

		if err == nil {
			t.Error("Expected error for missing fields, got none")
		}

		if err.Error() != "title, type, and application deadline are required" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Create program with invalid type", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Test Program",
			Type:                "invalid_type",
			ApplicationDeadline: "2025-12-31",
		}

		_, err := service.CreateProgram(ctx, input)

		if err == nil {
			t.Error("Expected error for invalid type, got none")
		}
	})

	t.Run("Create training program with invalid training type", func(t *testing.T) {
		invalidTrainingType := "invalid"
		input := dto.Programs{
			Title:               "Test Training",
			Type:                "training",
			TrainingType:        &invalidTrainingType,
			ApplicationDeadline: "2025-12-31",
		}

		_, err := service.CreateProgram(ctx, input)

		if err == nil {
			t.Error("Expected error for invalid training type, got none")
		}
	})

	t.Run("Create program with invalid creator user", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Test Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           999,
		}

		_, err := service.CreateProgram(ctx, input)

		if err == nil {
			t.Error("Expected error for invalid user, got none")
		}

		if err.Error() != "creator user not found" {
			t.Errorf("Expected 'creator user not found', got '%s'", err.Error())
		}
	})

	t.Run("Create certification program", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Halal Certification",
			Type:                "certification",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Type != "certification" {
			t.Errorf("Expected type 'certification', got '%s'", result.Type)
		}
	})

	t.Run("Create funding program", func(t *testing.T) {
		minAmount := 10000000.0
		maxAmount := 50000000.0
		interestRate := 5.5
		maxTenure := 12

		input := dto.Programs{
			Title:               "Business Loan",
			Type:                "funding",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
			MinAmount:           &minAmount,
			MaxAmount:           &maxAmount,
			InterestRate:        &interestRate,
			MaxTenureMonths:     &maxTenure,
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Type != "funding" {
			t.Errorf("Expected type 'funding', got '%s'", result.Type)
		}
	})
}

// Test UpdateProgram
func TestUpdateProgram(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	// Setup existing program
	mockRepo.programs[1] = model.Program{
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

	t.Run("Update existing program", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Updated Title",
			Type:                "certification",
			ApplicationDeadline: "2026-01-31",
			Banner:              "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
			ProviderLogo:        "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
		}

		result, err := service.UpdateProgram(ctx, 1, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Title != "Updated Title" {
			t.Errorf("Expected title 'Updated Title', got '%s'", result.Title)
		}
	})

	t.Run("Update non-existing program", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Updated Title",
			Type:                "training",
			ApplicationDeadline: "2026-01-31",
		}

		_, err := service.UpdateProgram(ctx, 999, input)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})

	t.Run("Update with missing required fields", func(t *testing.T) {
		input := dto.Programs{
			Description: "Only description",
		}

		_, err := service.UpdateProgram(ctx, 1, input)

		if err == nil {
			t.Error("Expected error for missing fields, got none")
		}
	})

	t.Run("Update with benefits and requirements", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Updated Program",
			Type:                "training",
			ApplicationDeadline: "2026-01-31",
			Benefits:            []string{"New Benefit 1", "New Benefit 2"},
			Requirements:        []string{"New Requirement 1"},
			Banner:              "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
			ProviderLogo:        "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
		}

		result, err := service.UpdateProgram(ctx, 1, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Benefits) != 2 {
			t.Errorf("Expected 2 benefits, got %d", len(result.Benefits))
		}
	})
}

// Test DeleteProgram
func TestDeleteProgram(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	mockRepo.programs[1] = model.Program{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
	}

	t.Run("Delete existing program", func(t *testing.T) {
		result, err := service.DeleteProgram(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected deleted program ID 1, got %d", result.ID)
		}

		// Verify deletion
		_, err = mockRepo.GetProgramByID(ctx, 1)
		if err == nil {
			t.Error("Expected program to be deleted")
		}
	})

	t.Run("Delete non-existing program", func(t *testing.T) {
		_, err := service.DeleteProgram(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})
}

// Test ActivateProgram
func TestActivateProgram(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	mockRepo.programs[1] = model.Program{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            false,
	}

	t.Run("Activate inactive program", func(t *testing.T) {
		result, err := service.ActivateProgram(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !result.IsActive {
			t.Error("Expected program to be active")
		}
	})

	t.Run("Activate non-existing program", func(t *testing.T) {
		_, err := service.ActivateProgram(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})
}

// Test DeactivateProgram
func TestDeactivateProgram(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	mockRepo.programs[1] = model.Program{
		ID:                  1,
		Title:               "Test Program",
		Type:                "training",
		ApplicationDeadline: "2025-12-31",
		IsActive:            true,
	}

	t.Run("Deactivate active program", func(t *testing.T) {
		result, err := service.DeactivateProgram(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.IsActive {
			t.Error("Expected program to be inactive")
		}
	})

	t.Run("Deactivate non-existing program", func(t *testing.T) {
		_, err := service.DeactivateProgram(ctx, 999)

		if err == nil {
			t.Error("Expected error for non-existing program, got none")
		}
	})
}

// Test Benefits and Requirements
func TestProgramBenefitsAndRequirements(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	t.Run("Create program with benefits and requirements", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Comprehensive Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
			Benefits: []string{
				"Free certificate",
				"Mentorship session",
				"Networking opportunity",
			},
			Requirements: []string{
				"Business license",
				"Tax registration",
			},
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Benefits) != 3 {
			t.Errorf("Expected 3 benefits, got %d", len(result.Benefits))
		}

		if len(result.Requirements) != 2 {
			t.Errorf("Expected 2 requirements, got %d", len(result.Requirements))
		}
	})

	t.Run("Update program benefits", func(t *testing.T) {
		mockRepo.programs[1] = model.Program{
			ID:                  1,
			Title:               "Test Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
		}

		input := dto.Programs{
			Title:               "Test Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
			Benefits:            []string{"Updated benefit"},
			Banner:              "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
			ProviderLogo:        "https://res.cloudinary.com/dblibr1t2/image/upload/v1762742634/umkmgo_logo.png",
		}

		result, err := service.UpdateProgram(ctx, 1, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Benefits) != 1 {
			t.Errorf("Expected 1 benefit, got %d", len(result.Benefits))
		}
	})
}

// Test Edge Cases
func TestProgramsServiceEdgeCases(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	t.Run("Create program with empty benefits and requirements", func(t *testing.T) {
		input := dto.Programs{
			Title:               "Minimal Program",
			Type:                "training",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
			Benefits:            []string{},
			Requirements:        []string{},
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result.Benefits) != 0 {
			t.Errorf("Expected 0 benefits, got %d", len(result.Benefits))
		}
	})

	t.Run("Toggle program activation multiple times", func(t *testing.T) {
		mockRepo.programs[1] = model.Program{
			ID:       1,
			Title:    "Toggle Program",
			Type:     "training",
			IsActive: true,
		}

		// Deactivate
		result1, _ := service.DeactivateProgram(ctx, 1)
		if result1.IsActive {
			t.Error("Expected program to be inactive")
		}

		// Activate
		result2, _ := service.ActivateProgram(ctx, 1)
		if !result2.IsActive {
			t.Error("Expected program to be active")
		}

		// Deactivate again
		result3, _ := service.DeactivateProgram(ctx, 1)
		if result3.IsActive {
			t.Error("Expected program to be inactive again")
		}
	})

	t.Run("Program with all optional fields", func(t *testing.T) {
		minAmount := 10000000.0
		maxAmount := 50000000.0
		interestRate := 5.5
		maxTenure := 12
		batch := 3
		batchStart := "2025-01-01"
		batchEnd := "2025-03-31"
		location := "Jakarta"
		trainingType := "hybrid"

		input := dto.Programs{
			Title:               "Complete Program",
			Description:         "Full description",
			Type:                "funding",
			ApplicationDeadline: "2025-12-31",
			CreatedBy:           1,
			MinAmount:           &minAmount,
			MaxAmount:           &maxAmount,
			InterestRate:        &interestRate,
			MaxTenureMonths:     &maxTenure,
			Batch:               &batch,
			BatchStartDate:      &batchStart,
			BatchEndDate:        &batchEnd,
			Location:            &location,
			TrainingType:        &trainingType,
		}

		result, err := service.CreateProgram(ctx, input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Title != "Complete Program" {
			t.Errorf("Expected title 'Complete Program', got '%s'", result.Title)
		}
	})
}

// Test Program Types Validation
func TestProgramTypesValidation(t *testing.T) {
	service, _ := setupProgramsService()
	ctx := context.Background()

	validTypes := []string{"training", "certification", "funding"}
	invalidTypes := []string{"invalid", "course", "grant", ""}

	for _, validType := range validTypes {
		t.Run("Valid type: "+validType, func(t *testing.T) {
			input := dto.Programs{
				Title:               "Test Program",
				Type:                validType,
				ApplicationDeadline: "2025-12-31",
				CreatedBy:           1,
			}

			_, err := service.CreateProgram(ctx, input)
			if err != nil {
				t.Errorf("Expected no error for valid type '%s', got %v", validType, err)
			}
		})
	}

	for _, invalidType := range invalidTypes {
		t.Run("Invalid type: "+invalidType, func(t *testing.T) {
			input := dto.Programs{
				Title:               "Test Program",
				Type:                invalidType,
				ApplicationDeadline: "2025-12-31",
				CreatedBy:           1,
			}

			_, err := service.CreateProgram(ctx, input)

			if err == nil {
				t.Errorf("Expected error for invalid type '%s', got none", invalidType)
			}
		})
	}
}

// Test Concurrent Operations
func TestProgramsConcurrentOperations(t *testing.T) {
	service, mockRepo := setupProgramsService()
	ctx := context.Background()

	mockRepo.programs[1] = model.Program{
		ID:    1,
		Title: "Concurrent Test Program",
		Type:  "training",
	}

	t.Run("Concurrent reads", func(t *testing.T) {
		done := make(chan bool)

		for i := 0; i < 10; i++ {
			go func() {
				_, err := service.GetProgramByID(ctx, 1)
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
