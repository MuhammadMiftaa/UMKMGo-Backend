package service

import (
	"errors"
	"testing"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// Mock repository for testing
type mockUsersRepository struct {
	users           map[string]model.Users
	usersById       map[int]model.Users
	roles           map[int]model.Roles
	permissions     []model.Permissions
	rolePermissions []model.RolePermissionsResponse
}

func newMockUsersRepository() *mockUsersRepository {
	// roleID := 1
	return &mockUsersRepository{
		users:     make(map[string]model.Users),
		usersById: make(map[int]model.Users),
		roles: map[int]model.Roles{
			1: {ID: 1, Name: "admin", Description: "Administrator"},
		},
		permissions: []model.Permissions{
			{ID: 1, Name: "View Users", Code: "VIEW_USERS"},
			{ID: 2, Name: "Manage Users", Code: "MANAGE_USERS"},
		},
		rolePermissions: []model.RolePermissionsResponse{
			{RoleID: 1, RoleName: "admin"},
		},
	}
}

func (m *mockUsersRepository) GetAllUsers() ([]model.Users, error) {
	users := []model.Users{}
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func (m *mockUsersRepository) GetUserByID(id int) (model.Users, error) {
	if user, exists := m.usersById[id]; exists {
		return user, nil
	}
	return model.Users{}, errors.New("user not found")
}

func (m *mockUsersRepository) GetUserByEmail(email string) (model.Users, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return model.Users{}, errors.New("user not found")
}

func (m *mockUsersRepository) CreateUser(user model.Users) (model.Users, error) {
	user.ID = len(m.users) + 1
	m.users[user.Email] = user
	m.usersById[user.ID] = user
	return user, nil
}

func (m *mockUsersRepository) UpdateUser(user model.Users) (model.Users, error) {
	if _, exists := m.usersById[user.ID]; !exists {
		return model.Users{}, errors.New("user not found")
	}
	m.users[user.Email] = user
	m.usersById[user.ID] = user
	return user, nil
}

func (m *mockUsersRepository) DeleteUser(user model.Users) (model.Users, error) {
	delete(m.users, user.Email)
	delete(m.usersById, user.ID)
	return user, nil
}

func (m *mockUsersRepository) GetAllRoles() ([]model.Roles, error) {
	roles := []model.Roles{}
	for _, r := range m.roles {
		roles = append(roles, r)
	}
	return roles, nil
}

func (m *mockUsersRepository) GetRoleByID(id int) (model.Roles, error) {
	if role, exists := m.roles[id]; exists {
		return role, nil
	}
	return model.Roles{}, errors.New("role not found")
}

func (m *mockUsersRepository) IsRoleExist(id int) bool {
	_, exists := m.roles[id]
	return exists
}

func (m *mockUsersRepository) IsPermissionExist(ids []string) ([]int, bool) {
	var existingIDs []int
	for _, perm := range m.permissions {
		for _, id := range ids {
			if perm.Code == id {
				existingIDs = append(existingIDs, perm.ID)
			}
		}
	}
	return existingIDs, len(existingIDs) == len(ids)
}

func (m *mockUsersRepository) GetListPermissions() ([]model.Permissions, error) {
	return m.permissions, nil
}

func (m *mockUsersRepository) GetListPermissionsByRoleID(roleID int) ([]string, error) {
	if _, exists := m.roles[roleID]; exists {
		// Mocking permissions for the role
		return []string{"VIEW_USERS", "MANAGE_USERS"}, nil
	}
	return nil, errors.New("role not found")
}

func (m *mockUsersRepository) GetListRolePermissions() ([]model.RolePermissionsResponse, error) {
	return m.rolePermissions, nil
}

func (m *mockUsersRepository) DeletePermissionsByRoleID(roleID int) error {
	return nil
}

func (m *mockUsersRepository) AddRolePermissions(roleID int, permissions []int) error {
	return nil
}

// Test Register
func TestRegister(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	roleID := 1
	tests := []struct {
		name        string
		input       dto.Users
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid registration",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "john@example.com",
				Password:        "Password123",
				ConfirmPassword: "Password123",
				RoleID:          &roleID,
			},
			expectError: false,
		},
		{
			name: "Missing required fields",
			input: dto.Users{
				Name:  "",
				Email: "john@example.com",
			},
			expectError: true,
			errorMsg:    "name, email, and password cannot be blank",
		},
		{
			name: "Invalid email format",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "invalid-email",
				Password:        "Password123",
				ConfirmPassword: "Password123",
				RoleID:          &roleID,
			},
			expectError: true,
			errorMsg:    "please enter a valid email address",
		},
		{
			name: "Password too short",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "john2@example.com",
				Password:        "Pass1",
				ConfirmPassword: "Pass1",
				RoleID:          &roleID,
			},
			expectError: true,
			errorMsg:    "password must be at least 8 characters long",
		},
		{
			name: "Password without letter",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "john3@example.com",
				Password:        "12345678",
				ConfirmPassword: "12345678",
				RoleID:          &roleID,
			},
			expectError: true,
			errorMsg:    "password must contain at least one letter",
		},
		{
			name: "Password without number",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "john4@example.com",
				Password:        "Password",
				ConfirmPassword: "Password",
				RoleID:          &roleID,
			},
			expectError: true,
			errorMsg:    "password must contain at least one number",
		},
		{
			name: "Password mismatch",
			input: dto.Users{
				Name:            "John Doe",
				Email:           "john5@example.com",
				Password:        "Password123",
				ConfirmPassword: "DifferentPass123",
				RoleID:          &roleID,
			},
			expectError: true,
			errorMsg:    "password and confirm password do not match",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Register(tt.input)

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
				if result.Email != tt.input.Email {
					t.Errorf("Expected email %s, got %s", tt.input.Email, result.Email)
				}
			}
		})
	}
}

// Test Login
// func TestLogin(t *testing.T) {
// 	mockRepo := newMockUsersRepository()
// 	service := NewUsersService(mockRepo)

// 	// Create a test user
// 	roleID := 1
// 	hashedPass := "$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS" // hash of "Password123"
// 	mockRepo.users["test@example.com"] = model.Users{
// 		ID:          1,
// 		Name:        "Test User",
// 		Email:       "test@example.com",
// 		Password:    hashedPass,
// 		RoleID:      roleID,
// 		IsActive:    true,
// 		LastLoginAt: time.Now(),
// 	}
// 	mockRepo.usersById[1] = mockRepo.users["test@example.com"]

// 	tests := []struct {
// 		name        string
// 		input       dto.Users
// 		expectError bool
// 		errorMsg    string
// 	}{
// 		{
// 			name: "Valid login",
// 			input: dto.Users{
// 				Email:    "test@example.com",
// 				Password: "Password123",
// 			},
// 			expectError: false,
// 		},
// 		{
// 			name: "Missing credentials",
// 			input: dto.Users{
// 				Email:    "",
// 				Password: "",
// 			},
// 			expectError: true,
// 			errorMsg:    "email and password cannot be blank",
// 		},
// 		{
// 			name: "User not found",
// 			input: dto.Users{
// 				Email:    "nonexistent@example.com",
// 				Password: "Password123",
// 			},
// 			expectError: true,
// 			errorMsg:    "user not found",
// 		},
// 		{
// 			name: "Wrong password",
// 			input: dto.Users{
// 				Email:    "test@example.com",
// 				Password: "WrongPassword123",
// 			},
// 			expectError: true,
// 			errorMsg:    "password is incorrect",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			token, err := service.Login(tt.input)

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
// 				if token == nil || *token == "" {
// 					t.Errorf("Expected valid token, got empty")
// 				}
// 			}
// 		})
// 	}
// }

// Test GetUserByID
func TestGetUserByID(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	// Create test user
	mockRepo.usersById[1] = model.Users{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	tests := []struct {
		name        string
		userID      int
		expectError bool
	}{
		{
			name:        "Valid user ID",
			userID:      1,
			expectError: false,
		},
		{
			name:        "Invalid user ID",
			userID:      999,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetUserByID(tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.ID != tt.userID {
					t.Errorf("Expected user ID %d, got %d", tt.userID, result.ID)
				}
			}
		})
	}
}

// Test UpdateUser
func TestUpdateUser(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	roleID := 1
	// Create test user
	mockRepo.usersById[1] = model.Users{
		ID:       1,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS",
		RoleID:   roleID,
	}

	tests := []struct {
		name        string
		userID      int
		input       dto.Users
		expectError bool
		errorMsg    string
	}{
		{
			name:   "Valid update",
			userID: 1,
			input: dto.Users{
				Name:            "Updated User",
				Email:           "updated@example.com",
				Password:        "NewPassword123",
				ConfirmPassword: "NewPassword123",
				RoleID:          &roleID,
			},
			expectError: false,
		},
		{
			name:   "Invalid user ID",
			userID: 999,
			input: dto.Users{
				Name:            "Updated User",
				Email:           "updated@example.com",
				Password:        "NewPassword123",
				ConfirmPassword: "NewPassword123",
				RoleID:          &roleID,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.UpdateUser(tt.userID, tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.Name != tt.input.Name {
					t.Errorf("Expected name %s, got %s", tt.input.Name, result.Name)
				}
			}
		})
	}
}

// Test DeleteUser
func TestDeleteUser(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	// Create test user
	mockRepo.usersById[1] = model.Users{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	tests := []struct {
		name        string
		userID      int
		expectError bool
	}{
		{
			name:        "Valid deletion",
			userID:      1,
			expectError: false,
		},
		{
			name:        "Invalid user ID",
			userID:      999,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.DeleteUser(tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result.ID != tt.userID {
					t.Errorf("Expected deleted user ID %d, got %d", tt.userID, result.ID)
				}
			}
		})
	}
}

// Test GetListPermissions
func TestGetListPermissions(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	result, err := service.GetListPermissions()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != len(mockRepo.permissions) {
		t.Errorf("Expected %d permissions, got %d", len(mockRepo.permissions), len(result))
	}
}

// Test GetListRolePermissions
func TestGetListRolePermissions(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	result, err := service.GetListRolePermissions()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != len(mockRepo.rolePermissions) {
		t.Errorf("Expected %d role permissions, got %d", len(mockRepo.rolePermissions), len(result))
	}
}

// Test UpdateRolePermissions
func TestUpdateRolePermissions(t *testing.T) {
	mockRepo := newMockUsersRepository()
	service := NewUsersService(mockRepo)

	tests := []struct {
		name        string
		input       dto.RolePermissions
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid update",
			input: dto.RolePermissions{
				RoleID:      1,
				Permissions: []string{"VIEW_USERS", "MANAGE_USERS"},
			},
			expectError: false,
		},
		{
			name: "Invalid role ID",
			input: dto.RolePermissions{
				RoleID:      999,
				Permissions: []string{"VIEW_USERS", "MANAGE_USERS"},
			},
			expectError: true,
			errorMsg:    "role id not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateRolePermissions(tt.input)

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
			}
		})
	}
}
