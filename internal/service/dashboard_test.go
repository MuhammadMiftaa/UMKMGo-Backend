package service

import (
	"context"
	"errors"
	"testing"
)

// ==================== MOCK DASHBOARD REPOSITORY ====================

type mockDashboardRepository struct {
	umkmByCardType           []map[string]interface{}
	applicationStatusSummary map[string]int64
	applicationStatusDetail  map[string]int64
	applicationByType        map[string]int64
	shouldError              bool
}

func newMockDashboardRepository() *mockDashboardRepository {
	return &mockDashboardRepository{
		umkmByCardType: []map[string]interface{}{
			{"name": "Kartu Produktif", "count": int64(150)},
			{"name": "Kartu Afirmatif", "count": int64(75)},
		},
		applicationStatusSummary: map[string]int64{
			"total_applications": 500,
			"in_process":         200,
			"approved":           250,
			"rejected":           50,
		},
		applicationStatusDetail: map[string]int64{
			"screening": 100,
			"revised":   50,
			"final":     50,
			"approved":  250,
			"rejected":  50,
		},
		applicationByType: map[string]int64{
			"funding":       200,
			"certification": 150,
			"training":      150,
		},
		shouldError: false,
	}
}

func (m *mockDashboardRepository) GetUMKMByCardType(ctx context.Context) ([]map[string]interface{}, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}
	return m.umkmByCardType, nil
}

func (m *mockDashboardRepository) GetApplicationStatusSummary(ctx context.Context) (map[string]int64, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}
	return m.applicationStatusSummary, nil
}

func (m *mockDashboardRepository) GetApplicationStatusDetail(ctx context.Context) (map[string]int64, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}
	return m.applicationStatusDetail, nil
}

func (m *mockDashboardRepository) GetApplicationByType(ctx context.Context) (map[string]int64, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}
	return m.applicationByType, nil
}

// ==================== TEST FUNCTIONS ====================

func setupDashboardService() (*dashboardService, *mockDashboardRepository) {
	mockRepo := newMockDashboardRepository()
	service := &dashboardService{
		dashboardRepository: mockRepo,
	}
	return service, mockRepo
}

// Test GetUMKMByCardType
func TestGetUMKMByCardType(t *testing.T) {
	service, mockRepo := setupDashboardService()
	ctx := context.Background()

	t.Run("Get UMKM by card type successfully", func(t *testing.T) {
		result, err := service.GetUMKMByCardType(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("Expected 2 card types, got %d", len(result))
		}

		if result[0].Name != "Kartu Produktif" {
			t.Errorf("Expected 'Kartu Produktif', got '%s'", result[0].Name)
		}

		if result[0].Count != 150 {
			t.Errorf("Expected count 150, got %d", result[0].Count)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetUMKMByCardType(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Handle empty results", func(t *testing.T) {
		mockRepo.umkmByCardType = []map[string]interface{}{}
		result, err := service.GetUMKMByCardType(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected 0 results, got %d", len(result))
		}
	})
}

// Test GetApplicationStatusSummary
func TestGetApplicationStatusSummary(t *testing.T) {
	service, mockRepo := setupDashboardService()
	ctx := context.Background()

	t.Run("Get application status summary successfully", func(t *testing.T) {
		result, err := service.GetApplicationStatusSummary(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 4 {
			t.Errorf("Expected 4 summary items, got %d", len(result))
		}

		if result[0].TotalApplications != 500 {
			t.Errorf("Expected total 500, got %d", result[0].TotalApplications)
		}

		if result[1].InProcess != 200 {
			t.Errorf("Expected in_process 200, got %d", result[1].InProcess)
		}

		if result[2].Approved != 250 {
			t.Errorf("Expected approved 250, got %d", result[2].Approved)
		}

		if result[3].Rejected != 50 {
			t.Errorf("Expected rejected 50, got %d", result[3].Rejected)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetApplicationStatusSummary(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Handle zero values", func(t *testing.T) {
		mockRepo.applicationStatusSummary = map[string]int64{
			"total_applications": 0,
			"in_process":         0,
			"approved":           0,
			"rejected":           0,
		}

		result, err := service.GetApplicationStatusSummary(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result[0].TotalApplications != 0 {
			t.Errorf("Expected total 0, got %d", result[0].TotalApplications)
		}
	})
}

// Test GetApplicationStatusDetail
func TestGetApplicationStatusDetail(t *testing.T) {
	service, mockRepo := setupDashboardService()
	ctx := context.Background()

	t.Run("Get application status detail successfully", func(t *testing.T) {
		result, err := service.GetApplicationStatusDetail(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 5 {
			t.Errorf("Expected 5 detail items, got %d", len(result))
		}

		if result[0].Screening != 100 {
			t.Errorf("Expected screening 100, got %d", result[0].Screening)
		}

		if result[1].Revised != 50 {
			t.Errorf("Expected revised 50, got %d", result[1].Revised)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetApplicationStatusDetail(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Handle missing status", func(t *testing.T) {
		mockRepo.applicationStatusDetail = map[string]int64{
			"screening": 100,
		}

		result, err := service.GetApplicationStatusDetail(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Should still return 5 items with zero values
		if len(result) != 5 {
			t.Errorf("Expected 5 items, got %d", len(result))
		}
	})
}

// Test GetApplicationByType
func TestGetApplicationByType(t *testing.T) {
	service, mockRepo := setupDashboardService()
	ctx := context.Background()

	t.Run("Get application by type successfully", func(t *testing.T) {
		result, err := service.GetApplicationByType(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 3 {
			t.Errorf("Expected 3 type items, got %d", len(result))
		}

		if result[0].Funding != 200 {
			t.Errorf("Expected funding 200, got %d", result[0].Funding)
		}

		if result[1].Certification != 150 {
			t.Errorf("Expected certification 150, got %d", result[1].Certification)
		}

		if result[2].Training != 150 {
			t.Errorf("Expected training 150, got %d", result[2].Training)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetApplicationByType(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Handle partial data", func(t *testing.T) {
		mockRepo.applicationByType = map[string]int64{
			"funding": 200,
		}

		result, err := service.GetApplicationByType(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 3 {
			t.Errorf("Expected 3 items, got %d", len(result))
		}
	})
}

// Test Edge Cases
func TestDashboardServiceEdgeCases(t *testing.T) {
	service, mockRepo := setupDashboardService()
	ctx := context.Background()

	t.Run("Handle large numbers", func(t *testing.T) {
		mockRepo.applicationStatusSummary = map[string]int64{
			"total_applications": 1000000,
			"in_process":         500000,
			"approved":           400000,
			"rejected":           100000,
		}

		result, err := service.GetApplicationStatusSummary(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result[0].TotalApplications != 1000000 {
			t.Errorf("Expected 1000000, got %d", result[0].TotalApplications)
		}
	})

	t.Run("Handle type conversion errors gracefully", func(t *testing.T) {
		mockRepo.umkmByCardType = []map[string]interface{}{
			{"name": "Kartu Produktif", "count": "invalid"}, // Invalid type
		}

		// Should panic or handle gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered from panic:", r)
			}
		}()

		_, err := service.GetUMKMByCardType(ctx)
		if err != nil {
			t.Log("Handled type conversion error:", err)
		}
	})

	t.Run("All services return consistent data", func(t *testing.T) {
		// Reset to known state
		mockRepo.shouldError = false
		mockRepo.applicationStatusSummary = map[string]int64{
			"total_applications": 500,
			"in_process":         200,
			"approved":           250,
			"rejected":           50,
		}

		summary, err1 := service.GetApplicationStatusSummary(ctx)
		detail, err2 := service.GetApplicationStatusDetail(ctx)
		byType, err3 := service.GetApplicationByType(ctx)

		if err1 != nil || err2 != nil || err3 != nil {
			t.Error("Expected no errors from all services")
		}

		// Check data consistency
		if summary[0].TotalApplications != 500 {
			t.Error("Inconsistent total applications")
		}

		if len(detail) != 5 {
			t.Error("Expected 5 detail statuses")
		}

		if len(byType) != 3 {
			t.Error("Expected 3 application types")
		}
	})
}

// Test Concurrent Access
func TestDashboardConcurrentAccess(t *testing.T) {
	service, _ := setupDashboardService()
	ctx := context.Background()

	t.Run("Concurrent reads should not cause issues", func(t *testing.T) {
		done := make(chan bool)
		numGoroutines := 10

		for i := 0; i < numGoroutines; i++ {
			go func() {
				_, err1 := service.GetUMKMByCardType(ctx)
				_, err2 := service.GetApplicationStatusSummary(ctx)
				_, err3 := service.GetApplicationStatusDetail(ctx)
				_, err4 := service.GetApplicationByType(ctx)

				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					t.Error("Concurrent access caused errors")
				}
				done <- true
			}()
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// Test Data Validation
func TestDashboardDataValidation(t *testing.T) {
	service, _ := setupDashboardService()
	ctx := context.Background()

	t.Run("Validate UMKM card type names", func(t *testing.T) {
		result, err := service.GetUMKMByCardType(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedNames := map[string]bool{
			"Kartu Produktif": false,
			"Kartu Afirmatif": false,
		}

		for _, item := range result {
			if _, exists := expectedNames[item.Name]; exists {
				expectedNames[item.Name] = true
			}
		}

		for name, found := range expectedNames {
			if !found {
				t.Errorf("Expected card type '%s' not found", name)
			}
		}
	})

	t.Run("Validate summary totals match details", func(t *testing.T) {
		summary, _ := service.GetApplicationStatusSummary(ctx)
		detail, _ := service.GetApplicationStatusDetail(ctx)

		// Sum of detailed statuses should equal total
		detailSum := detail[0].Screening + detail[1].Revised +
			detail[2].Final + detail[3].Approved + detail[4].Rejected

		if detailSum != summary[0].TotalApplications {
			t.Errorf("Detail sum (%d) doesn't match total (%d)",
				detailSum, summary[0].TotalApplications)
		}
	})

	t.Run("Validate application type totals", func(t *testing.T) {
		byType, _ := service.GetApplicationByType(ctx)

		totalByType := byType[0].Funding + byType[1].Certification + byType[2].Training

		if totalByType <= 0 {
			t.Error("Total applications by type should be positive")
		}
	})
}

// Benchmark Tests
func BenchmarkGetUMKMByCardType(b *testing.B) {
	service, _ := setupDashboardService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUMKMByCardType(ctx)
	}
}

func BenchmarkGetApplicationStatusSummary(b *testing.B) {
	service, _ := setupDashboardService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetApplicationStatusSummary(ctx)
	}
}

func BenchmarkAllDashboardServices(b *testing.B) {
	service, _ := setupDashboardService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUMKMByCardType(ctx)
		_, _ = service.GetApplicationStatusSummary(ctx)
		_, _ = service.GetApplicationStatusDetail(ctx)
		_, _ = service.GetApplicationByType(ctx)
	}
}
