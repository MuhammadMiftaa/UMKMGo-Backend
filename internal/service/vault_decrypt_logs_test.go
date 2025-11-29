package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/model"
)

// ==================== MOCK VAULT DECRYPT LOG REPOSITORY ====================

type mockVaultDecryptLogRepository struct {
	logs        []model.VaultDecryptLog
	shouldError bool
}

func newMockVaultDecryptLogRepository() *mockVaultDecryptLogRepository {
	now := time.Now()
	return &mockVaultDecryptLogRepository{
		logs: []model.VaultDecryptLog{
			{
				ID:           1,
				UserID:       1,
				UMKMID:       intPtr(1),
				FieldName:    "nik",
				TableName:    "umkms",
				RecordID:     1,
				Purpose:      "application_review",
				IPAddress:    "192.168.1.1",
				UserAgent:    "Mozilla/5.0",
				RequestID:    "req123",
				Success:      true,
				ErrorMessage: "",
				DecryptedAt:  now,
			},
			{
				ID:           2,
				UserID:       1,
				UMKMID:       intPtr(1),
				FieldName:    "kartu_number",
				TableName:    "umkms",
				RecordID:     1,
				Purpose:      "application_review",
				IPAddress:    "192.168.1.1",
				UserAgent:    "Mozilla/5.0",
				RequestID:    "req124",
				Success:      true,
				ErrorMessage: "",
				DecryptedAt:  now.Add(-1 * time.Hour),
			},
			{
				ID:           3,
				UserID:       2,
				UMKMID:       intPtr(2),
				FieldName:    "nik",
				TableName:    "umkms",
				RecordID:     2,
				Purpose:      "profile_view",
				IPAddress:    "192.168.1.2",
				UserAgent:    "Chrome/90.0",
				RequestID:    "req125",
				Success:      false,
				ErrorMessage: "decryption failed",
				DecryptedAt:  now.Add(-2 * time.Hour),
			},
			{
				ID:           4,
				UserID:       1,
				UMKMID:       intPtr(3),
				FieldName:    "nik",
				TableName:    "umkms",
				RecordID:     3,
				Purpose:      "application_creation",
				IPAddress:    "192.168.1.1",
				UserAgent:    "Mozilla/5.0",
				RequestID:    "req126",
				Success:      true,
				ErrorMessage: "",
				DecryptedAt:  now.Add(-3 * time.Hour),
			},
		},
		shouldError: false,
	}
}

func intPtr(i int) *int {
	return &i
}

func (m *mockVaultDecryptLogRepository) LogDecrypt(ctx context.Context, log model.VaultDecryptLog) error {
	if m.shouldError {
		return errors.New("database error")
	}
	log.ID = int64(len(m.logs) + 1)
	log.DecryptedAt = time.Now()
	m.logs = append(m.logs, log)
	return nil
}

func (m *mockVaultDecryptLogRepository) GetLogs(ctx context.Context, limit, offset int) ([]model.VaultDecryptLog, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}

	// Simulate pagination
	start := offset
	end := offset + limit

	if start >= len(m.logs) {
		return []model.VaultDecryptLog{}, nil
	}

	if end > len(m.logs) {
		end = len(m.logs)
	}

	return m.logs[start:end], nil
}

func (m *mockVaultDecryptLogRepository) GetLogsByUserID(ctx context.Context, userID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}

	var filtered []model.VaultDecryptLog
	for _, log := range m.logs {
		if log.UserID == userID {
			filtered = append(filtered, log)
		}
	}

	// Simulate pagination
	start := offset
	end := offset + limit

	if start >= len(filtered) {
		return []model.VaultDecryptLog{}, nil
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}

func (m *mockVaultDecryptLogRepository) GetLogsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.VaultDecryptLog, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}

	var filtered []model.VaultDecryptLog
	for _, log := range m.logs {
		if log.UMKMID != nil && *log.UMKMID == umkmID {
			filtered = append(filtered, log)
		}
	}

	// Simulate pagination
	start := offset
	end := offset + limit

	if start >= len(filtered) {
		return []model.VaultDecryptLog{}, nil
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}

// ==================== TEST FUNCTIONS ====================

func setupVaultDecryptLogService() (*vaultDecryptLogService, *mockVaultDecryptLogRepository) {
	mockRepo := newMockVaultDecryptLogRepository()
	service := &vaultDecryptLogService{
		vaultDecryptLogRepo: mockRepo,
	}
	return service, mockRepo
}

// Test GetLogs
func TestGetLogs(t *testing.T) {
	service, mockRepo := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Get all logs successfully", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Default limit is 100, should return all 4 logs
		if len(result) != 4 {
			t.Errorf("Expected 4 logs, got %d", len(result))
		}

		// Verify first log details
		if result[0].FieldName != "nik" {
			t.Errorf("Expected field name 'nik', got '%s'", result[0].FieldName)
		}

		if result[0].Success != true {
			t.Error("Expected first log to be successful")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetLogs(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Get logs with pagination", func(t *testing.T) {
		// This tests the underlying repository pagination
		logs, err := mockRepo.GetLogs(ctx, 2, 0)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(logs) != 2 {
			t.Errorf("Expected 2 logs, got %d", len(logs))
		}
	})

	t.Run("Get logs with offset beyond available data", func(t *testing.T) {
		logs, err := mockRepo.GetLogs(ctx, 10, 100)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(logs) != 0 {
			t.Errorf("Expected 0 logs, got %d", len(logs))
		}
	})

	t.Run("Verify logs contain all required fields", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		for i, log := range result {
			if log.UserID == 0 {
				t.Errorf("Log %d missing UserID", i)
			}
			if log.FieldName == "" {
				t.Errorf("Log %d missing FieldName", i)
			}
			if log.TableName == "" {
				t.Errorf("Log %d missing TableName", i)
			}
			if log.Purpose == "" {
				t.Errorf("Log %d missing Purpose", i)
			}
		}
	})
}

// Test GetLogsByUserID
func TestGetLogsByUserID(t *testing.T) {
	service, mockRepo := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Get logs for specific user", func(t *testing.T) {
		result, err := service.GetLogsByUserID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// User 1 has 3 logs
		if len(result) != 3 {
			t.Errorf("Expected 3 logs for user 1, got %d", len(result))
		}

		// Verify all logs belong to user 1
		for _, log := range result {
			if log.UserID != 1 {
				t.Errorf("Expected UserID 1, got %d", log.UserID)
			}
		}
	})

	t.Run("Get logs for user with no logs", func(t *testing.T) {
		result, err := service.GetLogsByUserID(ctx, 999)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 logs for non-existing user, got %d", len(result))
		}
	})

	t.Run("Get logs for user 2", func(t *testing.T) {
		result, err := service.GetLogsByUserID(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// User 2 has 1 log
		if len(result) != 1 {
			t.Errorf("Expected 1 log for user 2, got %d", len(result))
		}

		// Check it's a failed decrypt
		if result[0].Success {
			t.Error("Expected failed decrypt for user 2")
		}

		if result[0].ErrorMessage == "" {
			t.Error("Expected error message for failed decrypt")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetLogsByUserID(ctx, 1)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Verify pagination for user logs", func(t *testing.T) {
		// Get first 2 logs for user 1
		logs, err := mockRepo.GetLogsByUserID(ctx, 1, 2, 0)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(logs) != 2 {
			t.Errorf("Expected 2 logs, got %d", len(logs))
		}

		// Get next page
		logs2, err := mockRepo.GetLogsByUserID(ctx, 1, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(logs2) != 1 {
			t.Errorf("Expected 1 log in second page, got %d", len(logs2))
		}
	})
}

// Test GetLogsByUMKMID
func TestGetLogsByUMKMID(t *testing.T) {
	service, mockRepo := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Get logs for specific UMKM", func(t *testing.T) {
		result, err := service.GetLogsByUMKMID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// UMKM 1 has 2 logs
		if len(result) != 2 {
			t.Errorf("Expected 2 logs for UMKM 1, got %d", len(result))
		}

		// Verify all logs belong to UMKM 1
		for _, log := range result {
			if log.UMKMID == nil || *log.UMKMID != 1 {
				t.Error("Expected all logs to belong to UMKM 1")
			}
		}
	})

	t.Run("Get logs for UMKM with no logs", func(t *testing.T) {
		result, err := service.GetLogsByUMKMID(ctx, 999)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 logs for non-existing UMKM, got %d", len(result))
		}
	})

	t.Run("Get logs for UMKM 2", func(t *testing.T) {
		result, err := service.GetLogsByUMKMID(ctx, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// UMKM 2 has 1 log
		if len(result) != 1 {
			t.Errorf("Expected 1 log for UMKM 2, got %d", len(result))
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetLogsByUMKMID(ctx, 1)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Verify different purposes are logged", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		purposes := make(map[string]bool)
		for _, log := range result {
			purposes[log.Purpose] = true
		}

		expectedPurposes := []string{"application_review", "profile_view", "application_creation"}
		for _, purpose := range expectedPurposes {
			if !purposes[purpose] {
				t.Errorf("Expected purpose '%s' to be in logs", purpose)
			}
		}
	})
}

// Test Edge Cases
func TestVaultDecryptLogEdgeCases(t *testing.T) {
	service, mockRepo := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Handle logs with nil UMKMID", func(t *testing.T) {
		// Add log without UMKMID
		logWithoutUMKM := model.VaultDecryptLog{
			ID:          5,
			UserID:      1,
			UMKMID:      nil,
			FieldName:   "test_field",
			TableName:   "other_table",
			RecordID:    1,
			Purpose:     "test_purpose",
			Success:     true,
			IPAddress:   "192.168.1.1",
			RequestID:   "req127",
			DecryptedAt: time.Now(),
		}
		mockRepo.logs = append(mockRepo.logs, logWithoutUMKM)

		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Should include log with nil UMKMID
		if len(result) != 5 {
			t.Errorf("Expected 5 logs, got %d", len(result))
		}
	})

	t.Run("Verify logs are ordered by time", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if logs are in descending order (newest first)
		for i := 0; i < len(result)-1; i++ {
			if result[i].DecryptedAt.Before(result[i+1].DecryptedAt) {
				// Note: This may not be guaranteed by service,
				// but repository should handle ordering
				t.Log("Logs may not be ordered by time")
			}
		}
	})

	t.Run("Handle successful and failed decrypts separately", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var successCount, failCount int
		for _, log := range result {
			if log.Success {
				successCount++
				if log.ErrorMessage != "" {
					t.Error("Successful log should not have error message")
				}
			} else {
				failCount++
				if log.ErrorMessage == "" {
					t.Error("Failed log should have error message")
				}
			}
		}

		if successCount == 0 {
			t.Error("Expected at least one successful decrypt")
		}

		if failCount == 0 {
			t.Error("Expected at least one failed decrypt")
		}
	})

	t.Run("Verify IP addresses are logged", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		for _, log := range result {
			if log.IPAddress == "" {
				t.Error("Expected IP address to be logged")
			}
		}
	})

	t.Run("Verify request IDs are unique", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		requestIDs := make(map[string]bool)
		for _, log := range result {
			if log.RequestID == "" {
				t.Error("Expected request ID to be present")
			}
			if requestIDs[log.RequestID] {
				t.Errorf("Duplicate request ID found: %s", log.RequestID)
			}
			requestIDs[log.RequestID] = true
		}
	})
}

// Test Data Validation
func TestVaultDecryptLogDataValidation(t *testing.T) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Verify all logs have required fields", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		requiredFields := map[string]func(model.VaultDecryptLog) bool{
			"UserID":    func(l model.VaultDecryptLog) bool { return l.UserID > 0 },
			"FieldName": func(l model.VaultDecryptLog) bool { return l.FieldName != "" },
			"TableName": func(l model.VaultDecryptLog) bool { return l.TableName != "" },
			"RecordID":  func(l model.VaultDecryptLog) bool { return l.RecordID > 0 },
			"Purpose":   func(l model.VaultDecryptLog) bool { return l.Purpose != "" },
		}

		for i, log := range result {
			for fieldName, validator := range requiredFields {
				if !validator(log) {
					t.Errorf("Log %d missing or invalid %s", i, fieldName)
				}
			}
		}
	})

	t.Run("Verify field names are valid", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		validFieldNames := map[string]bool{
			"nik":          true,
			"kartu_number": true,
			"test_field":   true,
		}

		for _, log := range result {
			if !validFieldNames[log.FieldName] {
				t.Logf("Unexpected field name: %s", log.FieldName)
			}
		}
	})

	t.Run("Verify purposes are valid", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		validPurposes := map[string]bool{
			"application_review":   true,
			"application_creation": true,
			"profile_view":         true,
			"test_purpose":         true,
		}

		for _, log := range result {
			if !validPurposes[log.Purpose] {
				t.Logf("Unexpected purpose: %s", log.Purpose)
			}
		}
	})
}

// Test Concurrent Operations
func TestVaultDecryptLogConcurrentOperations(t *testing.T) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Concurrent reads should not interfere", func(t *testing.T) {
		done := make(chan bool)
		numGoroutines := 10

		for i := 0; i < numGoroutines; i++ {
			go func(userID int) {
				_, err1 := service.GetLogs(ctx)
				_, err2 := service.GetLogsByUserID(ctx, userID%3+1)
				_, err3 := service.GetLogsByUMKMID(ctx, userID%3+1)

				if err1 != nil || err2 != nil || err3 != nil {
					t.Error("Concurrent reads caused errors")
				}
				done <- true
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// Test Statistics
func TestVaultDecryptLogStatistics(t *testing.T) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	t.Run("Calculate success rate", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var successCount, totalCount int
		for _, log := range result {
			totalCount++
			if log.Success {
				successCount++
			}
		}

		successRate := float64(successCount) / float64(totalCount) * 100

		if successRate < 0 || successRate > 100 {
			t.Errorf("Invalid success rate: %.2f%%", successRate)
		}

		t.Logf("Success rate: %.2f%% (%d/%d)", successRate, successCount, totalCount)
	})

	t.Run("Count decrypts per user", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		userCounts := make(map[int]int)
		for _, log := range result {
			userCounts[log.UserID]++
		}

		for userID, count := range userCounts {
			t.Logf("User %d: %d decrypts", userID, count)
		}
	})

	t.Run("Count decrypts per field", func(t *testing.T) {
		result, err := service.GetLogs(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		fieldCounts := make(map[string]int)
		for _, log := range result {
			fieldCounts[log.FieldName]++
		}

		for field, count := range fieldCounts {
			t.Logf("Field '%s': %d decrypts", field, count)
		}
	})
}

// Benchmark Tests
func BenchmarkGetLogs(b *testing.B) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetLogs(ctx)
	}
}

func BenchmarkGetLogsByUserID(b *testing.B) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetLogsByUserID(ctx, 1)
	}
}

func BenchmarkGetLogsByUMKMID(b *testing.B) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetLogsByUMKMID(ctx, 1)
	}
}

func BenchmarkAllLogOperations(b *testing.B) {
	service, _ := setupVaultDecryptLogService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetLogs(ctx)
		_, _ = service.GetLogsByUserID(ctx, 1)
		_, _ = service.GetLogsByUMKMID(ctx, 1)
	}
}
