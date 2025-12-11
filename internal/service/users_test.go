package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"
	"UMKMGo-backend/internal/utils/constant"

	redisPackage "github.com/go-redis/redis/v8"
)

// ==================== Additional Mock Repositories ====================

type mockRedisRepository struct {
	data  map[string]string
	hdata map[string]map[string]string
}

func newMockRedisRepository() redis.RedisRepository {
	return &mockRedisRepository{
		data:  make(map[string]string),
		hdata: make(map[string]map[string]string),
	}
}

func (m *mockRedisRepository) Set(ctx context.Context, key, value string, exp time.Duration) error {
	m.data[key] = value
	return nil
}

func (m *mockRedisRepository) SetNX(ctx context.Context, key, value string, exp time.Duration) (bool, error) {
	if _, exists := m.data[key]; exists {
		return false, nil
	}
	m.data[key] = value
	return true, nil
}

func (m *mockRedisRepository) HSet(ctx context.Context, key string, value map[string]any, exp time.Duration) error {
	if m.hdata[key] == nil {
		m.hdata[key] = make(map[string]string)
	}
	for k, v := range value {
		m.hdata[key][k] = v.(string)
	}
	return nil
}

func (m *mockRedisRepository) HGet(ctx context.Context, key, field string) (string, error) {
	if hmap, exists := m.hdata[key]; exists {
		if val, ok := hmap[field]; ok {
			return val, nil
		}
	}
	return "", errors.New("field not found")
}

func (m *mockRedisRepository) Publish(ctx context.Context, channel, message string) error {
	return nil
}

func (m *mockRedisRepository) Subscribe(ctx context.Context, channel string) *redisPackage.PubSub {
	return nil
}

func (m *mockRedisRepository) Get(ctx context.Context, key string) (string, error) {
	if val, exists := m.data[key]; exists {
		return val, nil
	}
	return "", errors.New("key not found")
}

func (m *mockRedisRepository) Del(ctx context.Context, keys ...string) (int64, error) {
	var count int64
	for _, key := range keys {
		if _, exists := m.data[key]; exists {
			delete(m.data, key)
			count++
		}
	}
	return count, nil
}

func (m *mockRedisRepository) Exists(ctx context.Context, keys ...string) (int64, error) {
	var count int64
	for _, key := range keys {
		if _, exists := m.data[key]; exists {
			count++
		}
	}
	return count, nil
}

func (m *mockRedisRepository) Incr(ctx context.Context, key string) (int64, error) {
	return 1, nil
}

func (m *mockRedisRepository) Expire(ctx context.Context, key string, exp time.Duration) error {
	return nil
}

func (m *mockRedisRepository) MGet(ctx context.Context, keys []string) ([]any, error) {
	var values []any
	for _, key := range keys {
		if val, exists := m.data[key]; exists {
			values = append(values, val)
		} else {
			values = append(values, nil)
		}
	}
	return values, nil
}

func (m *mockRedisRepository) MSet(ctx context.Context, data map[string]any) error {
	for key, value := range data {
		m.data[key] = value.(string)
	}
	return nil
}

func (m *mockRedisRepository) Scan(ctx context.Context, match string, count int64) ([]string, error) {
	var keys []string
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys, nil
}

func (m *mockRedisRepository) Keys(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys, nil
}

type mockOTPRepository struct {
	otps map[string]*model.OTP
}

func newMockOTPRepository() *mockOTPRepository {
	return &mockOTPRepository{
		otps: make(map[string]*model.OTP),
	}
}

func (m *mockOTPRepository) CreateOTP(ctx context.Context, otp model.OTP) error {
	m.otps[otp.PhoneNumber] = &otp
	return nil
}

func (m *mockOTPRepository) GetOTPByPhone(ctx context.Context, phone string) (*model.OTP, error) {
	if otp, exists := m.otps[phone]; exists {
		return otp, nil
	}
	return nil, errors.New("OTP not found")
}

func (m *mockOTPRepository) GetOTPByTempToken(ctx context.Context, tempToken string) (*model.OTP, error) {
	for _, otp := range m.otps {
		if otp.TempToken != nil && *otp.TempToken == tempToken {
			return otp, nil
		}
	}
	return nil, errors.New("OTP not found")
}

func (m *mockOTPRepository) UpdateOTP(ctx context.Context, otp model.OTP) error {
	m.otps[otp.PhoneNumber] = &otp
	return nil
}

// Update the setup function to include all dependencies
func setupUsersServiceComplete() (*usersService, *mockUsersRepositoryForTests, redis.RedisRepository, *mockOTPRepository) {
	mockUserRepo := &mockUsersRepositoryForTests{
		users: make(map[int]model.User),
		roles: make(map[int]model.Role),
		umkms: make(map[string]model.UMKM),
		permissions: []model.Permission{
			{ID: 1, Code: "VIEW_DASHBOARD", Name: "View Dashboard"},
			{ID: 2, Code: "MANAGE_USERS", Name: "Manage Users"},
		},
		rolePermissions: make(map[int][]string),
		provinces: []dto.Province{
			{ID: 1, Name: "DKI Jakarta"},
			{ID: 2, Name: "Jawa Barat"},
		},
		cities: []dto.City{
			{ID: 1, Name: "Jakarta Pusat", ProvinceID: 1},
			{ID: 2, Name: "Bandung", ProvinceID: 2},
		},
	}

	mockRedisRepo := newMockRedisRepository()
	mockOTPRepo := newMockOTPRepository()

	// Setup default roles
	mockUserRepo.roles[1] = model.Role{ID: 1, Name: "superadmin"}
	mockUserRepo.roles[2] = model.Role{ID: 2, Name: "admin_screening"}
	mockUserRepo.roles[3] = model.Role{ID: 3, Name: "admin_vendor"}
	mockUserRepo.roles[4] = model.Role{ID: 4, Name: constant.RoleUMKM}

	service := &usersService{
		userRepository:  mockUserRepo,
		otpRepository:   mockOTPRepo,
		redisRepository: mockRedisRepo,
		minio:           nil,
	}

	return service, mockUserRepo, mockRedisRepo, mockOTPRepo
}

// ==================== TEST UpdateProfile ====================

func TestUsersServiceUpdateProfile(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "Old Name",
		Email: "old@example.com",
	}

	t.Run("Update profile successfully", func(t *testing.T) {
		request := dto.Users{
			Name:  "New Name",
			Email: "new@example.com",
		}

		result, err := service.UpdateProfile(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Name != "New Name" {
			t.Errorf("Expected name 'New Name', got '%s'", result.Name)
		}

		if result.Email != "new@example.com" {
			t.Errorf("Expected email 'new@example.com', got '%s'", result.Email)
		}
	})

	t.Run("Update profile with empty name", func(t *testing.T) {
		request := dto.Users{
			Name:  "",
			Email: "test@example.com",
		}

		_, err := service.UpdateProfile(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for empty name, got none")
		}
	})

	t.Run("Update profile with invalid email", func(t *testing.T) {
		request := dto.Users{
			Name:  "Test Name",
			Email: "invalid-email",
		}

		_, err := service.UpdateProfile(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	t.Run("Update profile with existing email", func(t *testing.T) {
		mockRepo.users[2] = model.User{
			ID:    2,
			Email: "existing@example.com",
		}

		request := dto.Users{
			Name:  "Test Name",
			Email: "existing@example.com",
		}

		_, err := service.UpdateProfile(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for existing email, got none")
		}
	})

	t.Run("Update profile for non-existent user", func(t *testing.T) {
		request := dto.Users{
			Name:  "Test Name",
			Email: "test@example.com",
		}

		_, err := service.UpdateProfile(ctx, 999, request)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

// ==================== TEST SetOTP ====================

func TestUsersServiceSetOTP(t *testing.T) {
	service, mockRepo, mockRedis, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Email: "test@example.com",
	}

	t.Run("Set OTP successfully", func(t *testing.T) {
		err := service.SetOTP(ctx, "test@example.com", "123456", 5*time.Minute)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify OTP is stored in Redis
		storedOTP, err := mockRedis.Get(ctx, "test@example.com")
		if err != nil {
			t.Errorf("Expected OTP to be stored, got error: %v", err)
		}
		if storedOTP != "123456" {
			t.Errorf("Expected OTP '123456', got '%s'", storedOTP)
		}
	})

	t.Run("Set OTP with empty email", func(t *testing.T) {
		err := service.SetOTP(ctx, "", "123456", 5*time.Minute)
		if err == nil {
			t.Error("Expected error for empty email, got none")
		}
	})

	t.Run("Set OTP with invalid email", func(t *testing.T) {
		err := service.SetOTP(ctx, "invalid-email", "123456", 5*time.Minute)
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	t.Run("Set OTP for non-existent user", func(t *testing.T) {
		err := service.SetOTP(ctx, "nonexistent@example.com", "123456", 5*time.Minute)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

// ==================== TEST ValidateOTP ====================

func TestUsersServiceValidateOTP(t *testing.T) {
	service, mockRepo, mockRedis, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Email: "test@example.com",
	}

	t.Run("Validate OTP successfully", func(t *testing.T) {
		mockRedis.Set(ctx, "test@example.com", "123456", 5*time.Minute)

		valid, err := service.ValidateOTP(ctx, "test@example.com", "123456")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !valid {
			t.Error("Expected OTP to be valid")
		}
	})

	t.Run("Validate OTP with wrong code", func(t *testing.T) {
		mockRedis.Set(ctx, "test@example.com", "123456", 5*time.Minute)

		valid, err := service.ValidateOTP(ctx, "test@example.com", "wrong")
		if err == nil {
			t.Error("Expected error for wrong OTP, got none")
		}
		if valid {
			t.Error("Expected OTP to be invalid")
		}
	})

	t.Run("Validate OTP with empty email", func(t *testing.T) {
		_, err := service.ValidateOTP(ctx, "", "123456")
		if err == nil {
			t.Error("Expected error for empty email, got none")
		}
	})

	t.Run("Validate OTP with invalid email", func(t *testing.T) {
		_, err := service.ValidateOTP(ctx, "invalid-email", "123456")
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})
}

// ==================== TEST VerifyUser ====================

func TestUsersServiceVerifyUser(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:       1,
		Name:     "Test User",
		Email:    "test@example.com",
		IsActive: false,
	}

	t.Run("Verify user successfully", func(t *testing.T) {
		result, err := service.VerifyUser(ctx, "test@example.com")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		// Check if user is now active
		user, _ := mockRepo.GetUserByEmail(ctx, "test@example.com")
		if !user.IsActive {
			t.Error("Expected user to be active")
		}
	})

	t.Run("Verify non-existent user", func(t *testing.T) {
		_, err := service.VerifyUser(ctx, "nonexistent@example.com")
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

// ==================== TEST GetUserByID and GetUserByEmail ====================

func TestUsersServiceGetUserByID(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	t.Run("Get existing user by ID", func(t *testing.T) {
		result, err := service.GetUserByID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
	})

	t.Run("Get non-existent user by ID", func(t *testing.T) {
		_, err := service.GetUserByID(ctx, 999)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

func TestUsersServiceGetUserByEmail(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	t.Run("Get existing user by email", func(t *testing.T) {
		result, err := service.GetUserByEmail(ctx, "test@example.com")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Email != "test@example.com" {
			t.Errorf("Expected email 'test@example.com', got '%s'", result.Email)
		}
	})

	t.Run("Get non-existent user by email", func(t *testing.T) {
		_, err := service.GetUserByEmail(ctx, "nonexistent@example.com")
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

// ==================== TEST Permission Functions ====================

func TestUsersServiceGetListPermissions(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	t.Run("Get list of permissions", func(t *testing.T) {
		result, err := service.GetListPermissions(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != len(mockRepo.permissions) {
			t.Errorf("Expected %d permissions, got %d", len(mockRepo.permissions), len(result))
		}
	})
}

func TestUsersServiceGetListRolePermissions(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.rolePermissions[1] = []string{"VIEW_DASHBOARD", "MANAGE_USERS"}

	t.Run("Get list of role permissions", func(t *testing.T) {
		result, err := service.GetListRolePermissions(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) == 0 {
			t.Error("Expected role permissions, got empty list")
		}
	})
}

func TestUsersServiceUpdateRolePermissions(t *testing.T) {
	service, _, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	t.Run("Update role permissions successfully", func(t *testing.T) {
		request := dto.RolePermissions{
			RoleID:      1,
			Permissions: []string{"VIEW_DASHBOARD", "MANAGE_USERS"},
		}

		err := service.UpdateRolePermissions(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Update with invalid role ID", func(t *testing.T) {
		request := dto.RolePermissions{
			RoleID:      999,
			Permissions: []string{"VIEW_DASHBOARD"},
		}

		err := service.UpdateRolePermissions(ctx, request)
		if err == nil {
			t.Error("Expected error for invalid role ID, got none")
		}
	})

	t.Run("Update with invalid permissions", func(t *testing.T) {
		request := dto.RolePermissions{
			RoleID:      1,
			Permissions: []string{"INVALID_PERMISSION"},
		}

		err := service.UpdateRolePermissions(ctx, request)
		if err == nil {
			t.Error("Expected error for invalid permissions, got none")
		}
	})
}

// ==================== TEST Mobile Functions ====================

func TestUsersServiceMetaCityAndProvince(t *testing.T) {
	service, _, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	t.Run("Get meta city and province", func(t *testing.T) {
		result, err := service.MetaCityAndProvince(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) == 0 {
			t.Error("Expected meta data, got empty list")
		}

		if len(result[0].Provinces) == 0 {
			t.Error("Expected provinces, got empty list")
		}

		if len(result[0].Cities) == 0 {
			t.Error("Expected cities, got empty list")
		}
	})
}

func TestUsersServiceRegisterMobile(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	t.Run("Register mobile with empty email", func(t *testing.T) {
		err := service.RegisterMobile(ctx, "", "81234567890")
		if err == nil {
			t.Error("Expected error for empty email, got none")
		}
	})

	t.Run("Register mobile with invalid email", func(t *testing.T) {
		err := service.RegisterMobile(ctx, "invalid-email", "81234567890")
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	t.Run("Register mobile with invalid phone", func(t *testing.T) {
		err := service.RegisterMobile(ctx, "test@example.com", "123")
		if err == nil {
			t.Error("Expected error for invalid phone, got none")
		}
	})

	t.Run("Register mobile with existing email", func(t *testing.T) {
		mockRepo.users[1] = model.User{
			ID:    1,
			Email: "existing@example.com",
		}

		err := service.RegisterMobile(ctx, "existing@example.com", "81234567890")
		if err == nil {
			t.Error("Expected error for existing email, got none")
		}
	})
}

func TestUsersServiceLoginMobile(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	hashedPass, _ := utils.PasswordHashing("Password123")
	mockRepo.umkms["81234567890"] = model.UMKM{
		ID:           1,
		BusinessName: "Test Business",
		Phone:        "81234567890",
		KartuType:    "produktif",
		User: model.User{
			ID:       1,
			Name:     "Test User",
			Email:    "test@example.com",
			Password: hashedPass,
		},
	}

	t.Run("Login mobile successfully", func(t *testing.T) {
		request := dto.UMKMMobile{
			Phone:    "081234567890",
			Password: "Password123",
		}

		token, err := service.LoginMobile(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if token == nil || *token == "" {
			t.Error("Expected token to be generated")
		}
	})

	t.Run("Login mobile with empty phone", func(t *testing.T) {
		request := dto.UMKMMobile{
			Phone:    "",
			Password: "Password123",
		}

		_, err := service.LoginMobile(ctx, request)
		if err == nil {
			t.Error("Expected error for empty phone, got none")
		}
	})

	t.Run("Login mobile with wrong password", func(t *testing.T) {
		request := dto.UMKMMobile{
			Phone:    "081234567890",
			Password: "WrongPassword",
		}

		_, err := service.LoginMobile(ctx, request)
		if err == nil {
			t.Error("Expected error for wrong password, got none")
		}
	})

	t.Run("Login mobile with non-existent user", func(t *testing.T) {
		request := dto.UMKMMobile{
			Phone:    "089999999999",
			Password: "Password123",
		}

		_, err := service.LoginMobile(ctx, request)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

func TestUsersServiceForgotPassword(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.umkms["81234567890"] = model.UMKM{
		ID:    1,
		Phone: "81234567890",
		User: model.User{
			Email: "test@example.com",
		},
	}

	t.Run("Forgot password with empty phone", func(t *testing.T) {
		err := service.ForgotPassword(ctx, "")
		if err == nil {
			t.Error("Expected error for empty phone, got none")
		}
	})

	t.Run("Forgot password with invalid phone", func(t *testing.T) {
		err := service.ForgotPassword(ctx, "123")
		if err == nil {
			t.Error("Expected error for invalid phone, got none")
		}
	})

	t.Run("Forgot password for non-existent user", func(t *testing.T) {
		err := service.ForgotPassword(ctx, "089999999999")
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

func TestUsersServiceResetPassword(t *testing.T) {
	service, mockRepo, _, mockOTPRepo := setupUsersServiceComplete()
	ctx := context.Background()

	tempToken := "temptoken123"
	mockRepo.users[1] = model.User{
		ID:    1,
		Email: "test@example.com",
	}

	mockOTPRepo.otps["81234567890"] = &model.OTP{
		PhoneNumber: "81234567890",
		Email:       "test@example.com",
		TempToken:   &tempToken,
		Status:      constant.OTPStatusActive,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
	}

	t.Run("Reset password with empty password", func(t *testing.T) {
		request := dto.ResetPasswordMobile{
			Password:        "",
			ConfirmPassword: "",
		}

		err := service.ResetPassword(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for empty password, got none")
		}
	})

	t.Run("Reset password with weak password", func(t *testing.T) {
		request := dto.ResetPasswordMobile{
			Password:        "weak",
			ConfirmPassword: "weak",
		}

		err := service.ResetPassword(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for weak password, got none")
		}
	})

	t.Run("Reset password with mismatched passwords", func(t *testing.T) {
		request := dto.ResetPasswordMobile{
			Password:        "Password123",
			ConfirmPassword: "DifferentPassword123",
		}

		err := service.ResetPassword(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for mismatched passwords, got none")
		}
	})

	t.Run("Reset password successfully", func(t *testing.T) {
		request := dto.ResetPasswordMobile{
			Password:        "NewPassword123",
			ConfirmPassword: "NewPassword123",
		}

		err := service.ResetPassword(ctx, request, tempToken)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify OTP status is updated
		otp, _ := mockOTPRepo.GetOTPByPhone(ctx, "81234567890")
		if otp.Status != constant.OTPStatusUsed {
			t.Error("Expected OTP status to be 'used'")
		}
	})

	t.Run("Reset password with invalid temp token", func(t *testing.T) {
		request := dto.ResetPasswordMobile{
			Password:        "NewPassword123",
			ConfirmPassword: "NewPassword123",
		}

		err := service.ResetPassword(ctx, request, "invalid_token")
		if err == nil {
			t.Error("Expected error for invalid temp token, got none")
		}
	})
}

// ==================== Enhanced Mock Repository ====================

type mockUsersRepositoryForTests struct {
	users           map[int]model.User
	roles           map[int]model.Role
	umkms           map[string]model.UMKM
	permissions     []model.Permission
	rolePermissions map[int][]string
	provinces       []dto.Province
	cities          []dto.City
}

func (m *mockUsersRepositoryForTests) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (m *mockUsersRepositoryForTests) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	user.ID = len(m.users) + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	m.users[user.ID] = user
	return user, nil
}

func (m *mockUsersRepositoryForTests) GetUserByID(ctx context.Context, id int) (model.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return model.User{}, errors.New("user not found")
}

func (m *mockUsersRepositoryForTests) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	user.UpdatedAt = time.Now()
	m.users[user.ID] = user
	return user, nil
}

func (m *mockUsersRepositoryForTests) DeleteUser(ctx context.Context, user model.User) (model.User, error) {
	delete(m.users, user.ID)
	return user, nil
}

func (m *mockUsersRepositoryForTests) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func (m *mockUsersRepositoryForTests) IsRoleExist(ctx context.Context, id int) bool {
	_, exists := m.roles[id]
	return exists
}

func (m *mockUsersRepositoryForTests) GetRoleByID(ctx context.Context, id int) (model.Role, error) {
	if role, exists := m.roles[id]; exists {
		return role, nil
	}
	return model.Role{}, errors.New("role not found")
}

func (m *mockUsersRepositoryForTests) GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error) {
	if perms, exists := m.rolePermissions[roleID]; exists {
		return perms, nil
	}
	return []string{"VIEW_DASHBOARD", "MANAGE_USERS"}, nil
}

func (m *mockUsersRepositoryForTests) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	for _, role := range m.roles {
		if role.Name == name {
			return role, nil
		}
	}
	return model.Role{}, errors.New("role not found")
}

func (m *mockUsersRepositoryForTests) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	for _, r := range m.roles {
		roles = append(roles, r)
	}
	return roles, nil
}

func (m *mockUsersRepositoryForTests) CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User) (dto.UMKMMobile, error) {
	user.ID = len(m.users) + 1
	m.users[user.ID] = user
	umkm.UserID = user.ID
	umkm.ID = len(m.umkms) + 1
	m.umkms[umkm.Phone] = umkm

	return dto.UMKMMobile{
		ID:           umkm.ID,
		UserID:       user.ID,
		Fullname:     user.Name,
		BusinessName: umkm.BusinessName,
		Email:        user.Email,
		Phone:        umkm.Phone,
	}, nil
}

func (m *mockUsersRepositoryForTests) GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error) {
	if umkm, exists := m.umkms[phone]; exists {
		return umkm, nil
	}
	return model.UMKM{}, errors.New("UMKM not found")
}

func (m *mockUsersRepositoryForTests) IsPermissionExist(ctx context.Context, ids []string) ([]int, bool) {
	var permIDs []int
	for _, code := range ids {
		for _, perm := range m.permissions {
			if perm.Code == code {
				permIDs = append(permIDs, perm.ID)
				break
			}
		}
	}
	return permIDs, len(permIDs) == len(ids)
}

func (m *mockUsersRepositoryForTests) GetListPermissions(ctx context.Context) ([]model.Permission, error) {
	return m.permissions, nil
}

func (m *mockUsersRepositoryForTests) GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error) {
	var result []dto.RolePermissionsResponse
	for roleID, perms := range m.rolePermissions {
		role, _ := m.GetRoleByID(ctx, roleID)
		result = append(result, dto.RolePermissionsResponse{
			RoleID:   roleID,
			RoleName: role.Name,
		})
		_ = perms
	}
	return result, nil
}

func (m *mockUsersRepositoryForTests) DeletePermissionsByRoleID(ctx context.Context, roleID int) error {
	delete(m.rolePermissions, roleID)
	return nil
}

func (m *mockUsersRepositoryForTests) AddRolePermissions(ctx context.Context, roleID int, permissions []int) error {
	var codes []string
	for _, permID := range permissions {
		for _, perm := range m.permissions {
			if perm.ID == permID {
				codes = append(codes, perm.Code)
				break
			}
		}
	}
	m.rolePermissions[roleID] = codes
	return nil
}

func (m *mockUsersRepositoryForTests) GetProvinces(ctx context.Context) ([]dto.Province, error) {
	return m.provinces, nil
}

func (m *mockUsersRepositoryForTests) GetCities(ctx context.Context) ([]dto.City, error) {
	return m.cities, nil
}

// ==================== TEST Register ====================

func TestUsersServiceRegister(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	t.Run("Register with valid data", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "New User",
			Email:           "newuser@example.com",
			Password:        "Password123",
			ConfirmPassword: "Password123",
			RoleID:          &roleID,
		}

		result, err := service.Register(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Email != request.Email {
			t.Errorf("Expected email '%s', got '%s'", request.Email, result.Email)
		}

		if result.Name != request.Name {
			t.Errorf("Expected name '%s', got '%s'", request.Name, result.Name)
		}

		if result.ID == 0 {
			t.Error("Expected user ID to be assigned")
		}
	})

	t.Run("Register with missing fields", func(t *testing.T) {
		request := dto.Users{
			Name:  "User",
			Email: "",
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for missing fields, got none")
		}

		expectedMsg := "name, email, and password cannot be blank"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Register with invalid email", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "invalid-email",
			Password:        "Password123",
			ConfirmPassword: "Password123",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	t.Run("Register with existing email", func(t *testing.T) {
		mockRepo.users[1] = model.User{
			ID:    1,
			Email: "existing@example.com",
		}

		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "existing@example.com",
			Password:        "Password123",
			ConfirmPassword: "Password123",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for existing email, got none")
		}
	})

	t.Run("Register with weak password - no letters", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "user2@example.com",
			Password:        "12345678",
			ConfirmPassword: "12345678",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for password without letters, got none")
		}
	})

	t.Run("Register with weak password - no digits", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "user3@example.com",
			Password:        "PasswordOnly",
			ConfirmPassword: "PasswordOnly",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for password without digits, got none")
		}
	})

	t.Run("Register with weak password - too short", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "user4@example.com",
			Password:        "Pass1",
			ConfirmPassword: "Pass1",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for short password, got none")
		}
	})

	t.Run("Register with password mismatch", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:            "User",
			Email:           "user5@example.com",
			Password:        "Password123",
			ConfirmPassword: "DifferentPass123",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for password mismatch, got none")
		}
	})

	t.Run("Register with invalid role ID", func(t *testing.T) {
		roleID := 999
		request := dto.Users{
			Name:            "User",
			Email:           "user6@example.com",
			Password:        "Password123",
			ConfirmPassword: "Password123",
			RoleID:          &roleID,
		}

		_, err := service.Register(ctx, request)
		if err == nil {
			t.Error("Expected error for invalid role ID, got none")
		}
	})
}

// ==================== TEST Login ====================

func TestUsersServiceLogin(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	// Setup user with hashed password
	hashedPass, _ := utils.PasswordHashing("Password123")
	mockRepo.users[1] = model.User{
		ID:       1,
		Email:    "test@example.com",
		Password: hashedPass,
		RoleID:   2,
		IsActive: true,
	}
	mockRepo.roles[2] = model.Role{
		ID:   2,
		Name: "admin",
	}

	t.Run("Login with valid credentials", func(t *testing.T) {
		request := dto.Users{
			Email:    "test@example.com",
			Password: "Password123",
		}

		token, err := service.Login(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if token == nil || *token == "" {
			t.Error("Expected token to be generated")
		}
	})

	t.Run("Login with wrong password", func(t *testing.T) {
		request := dto.Users{
			Email:    "test@example.com",
			Password: "WrongPassword",
		}

		_, err := service.Login(ctx, request)
		if err == nil {
			t.Error("Expected error for wrong password, got none")
		}
	})

	t.Run("Login with non-existent user", func(t *testing.T) {
		request := dto.Users{
			Email:    "nonexistent@example.com",
			Password: "Password123",
		}

		_, err := service.Login(ctx, request)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})

	t.Run("Login with inactive user", func(t *testing.T) {
		mockRepo.users[2] = model.User{
			ID:       2,
			Email:    "inactive@example.com",
			Password: hashedPass,
			IsActive: false,
		}

		request := dto.Users{
			Email:    "inactive@example.com",
			Password: "Password123",
		}

		_, err := service.Login(ctx, request)
		if err == nil {
			t.Error("Expected error for inactive user, got none")
		}
	})

	t.Run("Login with empty email", func(t *testing.T) {
		request := dto.Users{
			Email:    "",
			Password: "Password123",
		}

		_, err := service.Login(ctx, request)
		if err == nil {
			t.Error("Expected error for empty email, got none")
		}
	})

	t.Run("Login with empty password", func(t *testing.T) {
		request := dto.Users{
			Email:    "test@example.com",
			Password: "",
		}

		_, err := service.Login(ctx, request)
		if err == nil {
			t.Error("Expected error for empty password, got none")
		}
	})
}

// ==================== TEST GetAllUsers ====================

func TestUsersServiceGetAllUsers(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "User 1",
		Email: "user1@example.com",
		Roles: model.Role{Name: "admin"},
	}
	mockRepo.users[2] = model.User{
		ID:    2,
		Name:  "User 2",
		Email: "user2@example.com",
		Roles: model.Role{Name: "user"},
	}

	t.Run("Get all users", func(t *testing.T) {
		result, err := service.GetAllUsers(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("Expected 2 users, got %d", len(result))
		}
	})

	t.Run("Get all users when empty", func(t *testing.T) {
		mockRepo.users = make(map[int]model.User)

		result, err := service.GetAllUsers(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 users, got %d", len(result))
		}
	})
}

// ==================== TEST UpdateUser ====================

func TestUsersServiceUpdateUser(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "Old Name",
		Email: "old@example.com",
	}
	mockRepo.roles[2] = model.Role{ID: 2, Name: "admin"}

	t.Run("Update user successfully", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:   "New Name",
			Email:  "new@example.com",
			RoleID: &roleID,
		}

		result, err := service.UpdateUser(ctx, 1, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Name != "New Name" {
			t.Errorf("Expected name 'New Name', got '%s'", result.Name)
		}
	})

	t.Run("Update user with empty fields", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:   "",
			Email:  "test@example.com",
			RoleID: &roleID,
		}

		_, err := service.UpdateUser(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for empty fields, got none")
		}
	})

	t.Run("Update user with invalid email", func(t *testing.T) {
		roleID := 2
		request := dto.Users{
			Name:   "Test",
			Email:  "invalid-email",
			RoleID: &roleID,
		}

		_, err := service.UpdateUser(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	t.Run("Update user with invalid role", func(t *testing.T) {
		roleID := 999
		request := dto.Users{
			Name:   "Test",
			Email:  "test@example.com",
			RoleID: &roleID,
		}

		_, err := service.UpdateUser(ctx, 1, request)
		if err == nil {
			t.Error("Expected error for invalid role, got none")
		}
	})
}

// ==================== TEST DeleteUser ====================

func TestUsersServiceDeleteUser(t *testing.T) {
	service, mockRepo, _, _ := setupUsersServiceComplete()
	ctx := context.Background()

	mockRepo.users[1] = model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	t.Run("Delete existing user", func(t *testing.T) {
		result, err := service.DeleteUser(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
	})

	t.Run("Delete non-existent user", func(t *testing.T) {
		_, err := service.DeleteUser(ctx, 999)
		if err == nil {
			t.Error("Expected error for non-existent user, got none")
		}
	})
}

// ==================== TEST VerifyOTP (Mobile) ====================

func TestUsersServiceVerifyOTPMobile(t *testing.T) {
	service, _, _, mockOTPRepo := setupUsersServiceComplete()
	ctx := context.Background()

	validOTP := &model.OTP{
		PhoneNumber: "81234567890",
		Email:       "test@example.com",
		OTPCode:     "123456",
		Status:      constant.OTPStatusActive,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
	}
	mockOTPRepo.otps["81234567890"] = validOTP

	t.Run("Verify OTP successfully", func(t *testing.T) {
		token, err := service.VerifyOTP(ctx, "081234567890", "123456")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if token == nil || *token == "" {
			t.Error("Expected temp token to be generated")
		}
	})

	t.Run("Verify OTP with wrong code", func(t *testing.T) {
		_, err := service.VerifyOTP(ctx, "081234567890", "wrong")
		if err == nil {
			t.Error("Expected error for wrong OTP, got none")
		}
	})

	t.Run("Verify OTP with expired OTP", func(t *testing.T) {
		expiredOTP := &model.OTP{
			PhoneNumber: "81234567891",
			Email:       "test2@example.com",
			OTPCode:     "123456",
			Status:      constant.OTPStatusActive,
			ExpiresAt:   time.Now().Add(-5 * time.Minute),
		}
		mockOTPRepo.otps["81234567891"] = expiredOTP

		_, err := service.VerifyOTP(ctx, "081234567891", "123456")
		if err == nil {
			t.Error("Expected error for expired OTP, got none")
		}
	})

	t.Run("Verify OTP with invalid phone", func(t *testing.T) {
		_, err := service.VerifyOTP(ctx, "123", "123456")
		if err == nil {
			t.Error("Expected error for invalid phone, got none")
		}
	})
}

// ==================== TEST RegisterMobileProfile ====================

func TestUsersServiceRegisterMobileProfile(t *testing.T) {
	service, _, _, mockOTPRepo := setupUsersServiceComplete()
	ctx := context.Background()

	tempToken := "temptoken123"
	mockOTPRepo.otps["81234567890"] = &model.OTP{
		PhoneNumber: "81234567890",
		Email:       "test@example.com",
		TempToken:   &tempToken,
		Status:      constant.OTPStatusActive,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
	}

	t.Run("Register with missing fullname", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname: "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing fullname, got none")
		}
	})

	t.Run("Register with missing business name", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing business name, got none")
		}
	})

	t.Run("Register with missing NIK", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing NIK, got none")
		}
	})

	t.Run("Register with missing birth date", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing birth date, got none")
		}
	})

	t.Run("Register with missing gender", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing gender, got none")
		}
	})

	t.Run("Register with missing address", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing address, got none")
		}
	})

	t.Run("Register with missing province", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   0,
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing province, got none")
		}
	})

	t.Run("Register with missing city", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       0,
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing city, got none")
		}
	})

	t.Run("Register with missing district", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing district, got none")
		}
	})

	t.Run("Register with missing postal code", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing postal code, got none")
		}
	})

	t.Run("Register with missing kartu type", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "12345",
			KartuType:    "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing kartu type, got none")
		}
	})

	t.Run("Register with missing kartu number", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "12345",
			KartuType:    "produktif",
			KartuNumber:  "",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for missing kartu number, got none")
		}
	})

	t.Run("Register with invalid phone number", func(t *testing.T) {
		invalidPhoneToken := "invalid_phone_token"
		mockOTPRepo.otps["invalid"] = &model.OTP{
			PhoneNumber: "",
			Email:       "test3@example.com",
			TempToken:   &invalidPhoneToken,
			Status:      constant.OTPStatusActive,
			ExpiresAt:   time.Now().Add(5 * time.Minute),
		}

		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "12345",
			KartuType:    "produktif",
			KartuNumber:  "KUR123456",
			Password:     "Password123",
		}

		_, err := service.RegisterMobileProfile(ctx, request, invalidPhoneToken)
		if err == nil {
			t.Error("Expected error for invalid phone number, got none")
		}
	})

	t.Run("Register with invalid temp token", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
		}

		_, err := service.RegisterMobileProfile(ctx, request, "invalid_token")
		if err == nil {
			t.Error("Expected error for invalid temp token, got none")
		}
	})

	t.Run("Register with expired OTP", func(t *testing.T) {
		expiredToken := "expired_token"
		mockOTPRepo.otps["81234567891"] = &model.OTP{
			PhoneNumber: "81234567891",
			Email:       "test2@example.com",
			TempToken:   &expiredToken,
			Status:      constant.OTPStatusUsed,
			ExpiresAt:   time.Now().Add(-5 * time.Minute),
		}

		request := dto.UMKMMobile{
			Fullname: "Test User",
		}

		_, err := service.RegisterMobileProfile(ctx, request, expiredToken)
		if err == nil {
			t.Error("Expected error for expired OTP, got none")
		}
	})

	t.Run("Register with weak password", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "1990-01-01",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "12345",
			KartuType:    "produktif",
			KartuNumber:  "KUR123456",
			Password:     "weak",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for weak password, got none")
		}
	})

	t.Run("Register with invalid birth date format", func(t *testing.T) {
		request := dto.UMKMMobile{
			Fullname:     "Test User",
			BusinessName: "Test Business",
			NIK:          "1234567890123456",
			BirthDate:    "invalid-date",
			Gender:       "male",
			Address:      "Test Address",
			ProvinceID:   1,
			CityID:       1,
			District:     "Test District",
			PostalCode:   "12345",
			KartuType:    "produktif",
			KartuNumber:  "KUR123456",
			Password:     "Password123",
		}

		_, err := service.RegisterMobileProfile(ctx, request, tempToken)
		if err == nil {
			t.Error("Expected error for invalid birth date format, got none")
		}
	})
}
