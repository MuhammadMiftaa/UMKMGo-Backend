package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

// ==================== MOCK SLA REPOSITORY ====================

type mockSLARepository struct {
	slas         map[string]model.SLA
	applications []model.Application
	programs     []model.Program
	shouldError  bool
}

func newMockSLARepository() *mockSLARepository {
	return &mockSLARepository{
		slas: map[string]model.SLA{
			"screening": {
				ID:          1,
				Status:      "screening",
				MaxDays:     7,
				Description: "Screening phase - 7 days",
				Base: model.Base{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			"final": {
				ID:          2,
				Status:      "final",
				MaxDays:     14,
				Description: "Final phase - 14 days",
				Base: model.Base{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
		applications: []model.Application{
			{
				ID:          1,
				UMKMID:      1,
				ProgramID:   1,
				Type:        "training",
				Status:      "screening",
				SubmittedAt: time.Now(),
				UMKM: model.UMKM{
					ID:           1,
					BusinessName: "Test Business",
					User:         model.User{Name: "Test User"},
					City: model.City{
						Province: model.Province{Name: "DKI Jakarta"},
					},
				},
				Program: model.Program{
					ID:    1,
					Title: "Training Program",
				},
			},
			{
				ID:          2,
				UMKMID:      2,
				ProgramID:   2,
				Type:        "funding",
				Status:      "approved",
				SubmittedAt: time.Now(),
				UMKM: model.UMKM{
					ID:           2,
					BusinessName: "Another Business",
					User:         model.User{Name: "Another User"},
					City: model.City{
						Province: model.Province{Name: "West Java"},
					},
				},
				Program: model.Program{
					ID:    2,
					Title: "Funding Program",
				},
			},
		},
		programs: []model.Program{
			{
				ID:       1,
				Title:    "Training Program",
				Type:     "training",
				Provider: "Provider A",
				IsActive: true,
			},
			{
				ID:       2,
				Title:    "Funding Program",
				Type:     "funding",
				Provider: "Provider B",
				IsActive: false,
			},
		},
		shouldError: false,
	}
}

func (m *mockSLARepository) GetSLAByStatus(ctx context.Context, status string) (model.SLA, error) {
	if m.shouldError {
		return model.SLA{}, errors.New("database error")
	}
	if sla, exists := m.slas[status]; exists {
		return sla, nil
	}
	return model.SLA{}, errors.New("SLA not found")
}

func (m *mockSLARepository) UpdateSLA(ctx context.Context, sla model.SLA) (model.SLA, error) {
	if m.shouldError {
		return model.SLA{}, errors.New("database error")
	}
	sla.UpdatedAt = time.Now()
	m.slas[sla.Status] = sla
	return sla, nil
}

func (m *mockSLARepository) GetApplicationsForExport(ctx context.Context, applicationType string) ([]model.Application, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}

	if applicationType == "all" {
		return m.applications, nil
	}

	var filtered []model.Application
	for _, app := range m.applications {
		if app.Type == applicationType {
			filtered = append(filtered, app)
		}
	}
	return filtered, nil
}

func (m *mockSLARepository) GetProgramsForExport(ctx context.Context, applicationType string) ([]model.Program, error) {
	if m.shouldError {
		return nil, errors.New("database error")
	}

	if applicationType == "all" {
		return m.programs, nil
	}

	var filtered []model.Program
	for _, prog := range m.programs {
		if prog.Type == applicationType {
			filtered = append(filtered, prog)
		}
	}
	return filtered, nil
}

// ==================== TEST FUNCTIONS ====================

func setupSLAService() (*slaService, *mockSLARepository) {
	mockRepo := newMockSLARepository()
	service := &slaService{
		slaRepository: mockRepo,
	}
	return service, mockRepo
}

// Test GetSLAScreening
func TestGetSLAScreening(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Get screening SLA successfully", func(t *testing.T) {
		result, err := service.GetSLAScreening(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "screening" {
			t.Errorf("Expected status 'screening', got '%s'", result.Status)
		}

		if result.MaxDays != 7 {
			t.Errorf("Expected max days 7, got %d", result.MaxDays)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetSLAScreening(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Handle missing SLA", func(t *testing.T) {
		delete(mockRepo.slas, "screening")
		_, err := service.GetSLAScreening(ctx)

		if err == nil {
			t.Error("Expected error for missing SLA, got none")
		}

		// Restore
		mockRepo.slas["screening"] = model.SLA{
			ID:      1,
			Status:  "screening",
			MaxDays: 7,
		}
	})
}

// Test GetSLAFinal
func TestGetSLAFinal(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Get final SLA successfully", func(t *testing.T) {
		result, err := service.GetSLAFinal(ctx)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Status != "final" {
			t.Errorf("Expected status 'final', got '%s'", result.Status)
		}

		if result.MaxDays != 14 {
			t.Errorf("Expected max days 14, got %d", result.MaxDays)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		_, err := service.GetSLAFinal(ctx)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})
}

// Test UpdateSLAScreening
func TestUpdateSLAScreening(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Update screening SLA successfully", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays:     10,
			Description: "Updated screening phase - 10 days",
		}

		result, err := service.UpdateSLAScreening(ctx, slaDTO)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.MaxDays != 10 {
			t.Errorf("Expected max days 10, got %d", result.MaxDays)
		}

		if result.Description != "Updated screening phase - 10 days" {
			t.Errorf("Expected description to be updated")
		}
	})

	t.Run("Reject zero or negative max days", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays: 0,
		}

		_, err := service.UpdateSLAScreening(ctx, slaDTO)

		if err == nil {
			t.Error("Expected error for zero max days, got none")
		}

		if err.Error() != "max_days must be greater than 0" {
			t.Errorf("Expected specific error message, got '%s'", err.Error())
		}
	})

	t.Run("Reject negative max days", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays: -5,
		}

		_, err := service.UpdateSLAScreening(ctx, slaDTO)

		if err == nil {
			t.Error("Expected error for negative max days, got none")
		}
	})

	t.Run("Update without changing description", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays: 8,
		}

		result, err := service.UpdateSLAScreening(ctx, slaDTO)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.MaxDays != 8 {
			t.Errorf("Expected max days 8, got %d", result.MaxDays)
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		slaDTO := dto.SLA{MaxDays: 10}

		_, err := service.UpdateSLAScreening(ctx, slaDTO)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})
}

// Test UpdateSLAFinal
func TestUpdateSLAFinal(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Update final SLA successfully", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays:     20,
			Description: "Updated final phase - 20 days",
		}

		result, err := service.UpdateSLAFinal(ctx, slaDTO)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.MaxDays != 20 {
			t.Errorf("Expected max days 20, got %d", result.MaxDays)
		}
	})

	t.Run("Reject zero or negative max days", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays: 0,
		}

		_, err := service.UpdateSLAFinal(ctx, slaDTO)

		if err == nil {
			t.Error("Expected error for zero max days, got none")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		slaDTO := dto.SLA{MaxDays: 20}

		_, err := service.UpdateSLAFinal(ctx, slaDTO)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})
}

// Test ExportApplications
func TestExportApplications(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Export all applications as PDF", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		data, filename, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(data) == 0 {
			t.Error("Expected data to be generated")
		}

		if !strings.HasSuffix(filename, ".txt") {
			t.Errorf("Expected .txt extension, got %s", filename)
		}

		if !strings.Contains(filename, "applications_all") {
			t.Errorf("Expected filename to contain 'applications_all', got %s", filename)
		}

		// Check content
		content := string(data)
		if !strings.Contains(content, "LAPORAN PENGAJUAN UMKM") {
			t.Error("Expected report header in content")
		}

		if !strings.Contains(content, "Test Business") {
			t.Error("Expected business name in content")
		}
	})

	t.Run("Export filtered applications as Excel/CSV", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "excel",
			ApplicationType: "training",
		}

		data, filename, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.HasSuffix(filename, ".csv") {
			t.Errorf("Expected .csv extension, got %s", filename)
		}

		if !strings.Contains(filename, "applications_training") {
			t.Errorf("Expected filename to contain 'applications_training', got %s", filename)
		}

		// Check CSV header
		content := string(data)
		if !strings.Contains(content, "No,UMKM,Program,Status,Tanggal Pengajuan") {
			t.Error("Expected CSV header in content")
		}
	})

	t.Run("Export funding applications", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "funding",
		}

		data, filename, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.Contains(filename, "funding") {
			t.Error("Expected filename to contain 'funding'")
		}

		content := string(data)
		if !strings.Contains(content, "Another Business") {
			t.Error("Expected funding application business name")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		_, _, err := service.ExportApplications(ctx, request)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Export empty results", func(t *testing.T) {
		mockRepo.applications = []model.Application{}
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "certification",
		}

		data, filename, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		if !strings.Contains(content, "Total Pengajuan: 0") {
			t.Error("Expected total count of 0")
		}

		if filename == "" {
			t.Error("Expected filename to be generated")
		}
	})
}

// Test ExportPrograms
func TestExportPrograms(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Export all programs as PDF", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		data, filename, err := service.ExportPrograms(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(data) == 0 {
			t.Error("Expected data to be generated")
		}

		if !strings.HasSuffix(filename, ".txt") {
			t.Errorf("Expected .txt extension, got %s", filename)
		}

		// Check content
		content := string(data)
		if !strings.Contains(content, "LAPORAN PROGRAM UMKM") {
			t.Error("Expected report header in content")
		}

		if !strings.Contains(content, "Training Program") {
			t.Error("Expected program title in content")
		}
	})

	t.Run("Export programs as Excel/CSV", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "excel",
			ApplicationType: "all",
		}

		data, filename, err := service.ExportPrograms(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.HasSuffix(filename, ".csv") {
			t.Errorf("Expected .csv extension, got %s", filename)
		}

		// Check CSV header
		content := string(data)
		if !strings.Contains(content, "No,Program,Provider,Tipe,Status Aktif") {
			t.Error("Expected CSV header in content")
		}
	})

	t.Run("Export filtered programs by type", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "training",
		}

		data, _, err := service.ExportPrograms(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		if !strings.Contains(content, "Tipe: training") {
			t.Error("Expected type filter in content")
		}
	})

	t.Run("Handle database error", func(t *testing.T) {
		mockRepo.shouldError = true
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		_, _, err := service.ExportPrograms(ctx, request)

		if err == nil {
			t.Error("Expected error, got none")
		}
		mockRepo.shouldError = false
	})

	t.Run("Export with active/inactive status", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "excel",
			ApplicationType: "all",
		}

		data, _, err := service.ExportPrograms(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		if !strings.Contains(content, "Aktif") || !strings.Contains(content, "Tidak Aktif") {
			t.Error("Expected both active and inactive status in content")
		}
	})
}

// Test Edge Cases
func TestSLAServiceEdgeCases(t *testing.T) {
	service, mockRepo := setupSLAService()
	ctx := context.Background()

	t.Run("Update SLA with very large max days", func(t *testing.T) {
		slaDTO := dto.SLA{
			MaxDays:     365,
			Description: "One year SLA",
		}

		result, err := service.UpdateSLAScreening(ctx, slaDTO)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.MaxDays != 365 {
			t.Errorf("Expected max days 365, got %d", result.MaxDays)
		}
	})

	t.Run("Export with special characters in business names", func(t *testing.T) {
		mockRepo.applications[0].UMKM.BusinessName = "Test & Company, Ltd."

		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		data, _, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		if !strings.Contains(content, "Test & Company") {
			t.Error("Expected special characters to be preserved")
		}
	})

	t.Run("Both SLA statuses can be updated independently", func(t *testing.T) {
		screeningSLA := dto.SLA{MaxDays: 5}
		finalSLA := dto.SLA{MaxDays: 15}

		result1, err1 := service.UpdateSLAScreening(ctx, screeningSLA)
		result2, err2 := service.UpdateSLAFinal(ctx, finalSLA)

		if err1 != nil || err2 != nil {
			t.Error("Expected no errors")
		}

		if result1.MaxDays != 5 || result2.MaxDays != 15 {
			t.Error("Expected both SLAs to be updated independently")
		}
	})

	t.Run("Filename contains timestamp", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		_, filename, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if filename contains date/time pattern
		if !strings.Contains(filename, "_") {
			t.Error("Expected filename to contain timestamp separator")
		}
	})
}

// Test Export Content Validation
func TestExportContentValidation(t *testing.T) {
	service, _ := setupSLAService()
	ctx := context.Background()

	t.Run("PDF export contains all required fields", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		data, _, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		requiredFields := []string{
			"LAPORAN PENGAJUAN UMKM",
			"Tanggal:",
			"Total Pengajuan:",
			"UMKM:",
			"Program:",
			"Status:",
			"Tanggal Pengajuan:",
		}

		for _, field := range requiredFields {
			if !strings.Contains(content, field) {
				t.Errorf("Expected field '%s' in PDF content", field)
			}
		}
	})

	t.Run("CSV export has proper structure", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "excel",
			ApplicationType: "all",
		}

		data, _, err := service.ExportApplications(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		lines := strings.Split(content, "\n")

		if len(lines) < 2 {
			t.Error("Expected at least header and one data row")
		}

		// Check header
		header := lines[0]
		if !strings.Contains(header, ",") {
			t.Error("Expected CSV separator in header")
		}
	})

	t.Run("Program export contains provider information", func(t *testing.T) {
		request := dto.ExportRequest{
			FileType:        "pdf",
			ApplicationType: "all",
		}

		data, _, err := service.ExportPrograms(ctx, request)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		content := string(data)
		if !strings.Contains(content, "Provider:") {
			t.Error("Expected provider field in program export")
		}

		if !strings.Contains(content, "Provider A") {
			t.Error("Expected actual provider name in content")
		}
	})
}

// Test Concurrent Operations
func TestSLAConcurrentOperations(t *testing.T) {
	service, _ := setupSLAService()
	ctx := context.Background()

	t.Run("Concurrent reads should not interfere", func(t *testing.T) {
		done := make(chan bool)
		numGoroutines := 10

		for i := 0; i < numGoroutines; i++ {
			go func() {
				_, err1 := service.GetSLAScreening(ctx)
				_, err2 := service.GetSLAFinal(ctx)

				if err1 != nil || err2 != nil {
					t.Error("Concurrent reads caused errors")
				}
				done <- true
			}()
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})

	t.Run("Sequential updates should work correctly", func(t *testing.T) {
		updates := []int{5, 10, 15, 20}

		for _, days := range updates {
			slaDTO := dto.SLA{MaxDays: days}
			result, err := service.UpdateSLAScreening(ctx, slaDTO)
			if err != nil {
				t.Errorf("Expected no error for update to %d days, got %v", days, err)
			}

			if result.MaxDays != days {
				t.Errorf("Expected max days %d, got %d", days, result.MaxDays)
			}
		}
	})
}

// Benchmark Tests
func BenchmarkGetSLAScreening(b *testing.B) {
	service, _ := setupSLAService()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetSLAScreening(ctx)
	}
}

func BenchmarkUpdateSLA(b *testing.B) {
	service, _ := setupSLAService()
	ctx := context.Background()
	slaDTO := dto.SLA{MaxDays: 10}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.UpdateSLAScreening(ctx, slaDTO)
	}
}

func BenchmarkExportApplications(b *testing.B) {
	service, _ := setupSLAService()
	ctx := context.Background()
	request := dto.ExportRequest{
		FileType:        "pdf",
		ApplicationType: "all",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = service.ExportApplications(ctx, request)
	}
}

func BenchmarkExportPrograms(b *testing.B) {
	service, _ := setupSLAService()
	ctx := context.Background()
	request := dto.ExportRequest{
		FileType:        "excel",
		ApplicationType: "all",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = service.ExportPrograms(ctx, request)
	}
}
