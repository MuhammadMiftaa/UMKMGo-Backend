package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
)

type SLAService interface {
	GetSLAScreening(ctx context.Context) (dto.SLA, error)
	GetSLAFinal(ctx context.Context) (dto.SLA, error)
	UpdateSLAScreening(ctx context.Context, slaDTO dto.SLA) (dto.SLA, error)
	UpdateSLAFinal(ctx context.Context, slaDTO dto.SLA) (dto.SLA, error)
	ExportApplications(ctx context.Context, request dto.ExportRequest) ([]byte, string, error)
	ExportPrograms(ctx context.Context, request dto.ExportRequest) ([]byte, string, error)
}

type slaService struct {
	slaRepository repository.SLARepository
}

func NewSLAService(slaRepo repository.SLARepository) SLAService {
	return &slaService{
		slaRepository: slaRepo,
	}
}

func (s *slaService) GetSLAScreening(ctx context.Context) (dto.SLA, error) {
	sla, err := s.slaRepository.GetSLAByStatus(ctx, "screening")
	if err != nil {
		return dto.SLA{}, err
	}

	return dto.SLA{
		ID:          sla.ID,
		Status:      sla.Status,
		MaxDays:     sla.MaxDays,
		Description: sla.Description,
		UpdatedAt:   sla.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *slaService) GetSLAFinal(ctx context.Context) (dto.SLA, error) {
	sla, err := s.slaRepository.GetSLAByStatus(ctx, "final")
	if err != nil {
		return dto.SLA{}, err
	}

	return dto.SLA{
		ID:          sla.ID,
		Status:      sla.Status,
		MaxDays:     sla.MaxDays,
		Description: sla.Description,
		UpdatedAt:   sla.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *slaService) UpdateSLAScreening(ctx context.Context, slaDTO dto.SLA) (dto.SLA, error) {
	if slaDTO.MaxDays <= 0 {
		return dto.SLA{}, errors.New("max_days must be greater than 0")
	}

	existingSLA, err := s.slaRepository.GetSLAByStatus(ctx, "screening")
	if err != nil {
		return dto.SLA{}, err
	}

	existingSLA.MaxDays = slaDTO.MaxDays
	if slaDTO.Description != "" {
		existingSLA.Description = slaDTO.Description
	}

	updatedSLA, err := s.slaRepository.UpdateSLA(ctx, existingSLA)
	if err != nil {
		return dto.SLA{}, err
	}

	return dto.SLA{
		ID:          updatedSLA.ID,
		Status:      updatedSLA.Status,
		MaxDays:     updatedSLA.MaxDays,
		Description: updatedSLA.Description,
		UpdatedAt:   updatedSLA.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *slaService) UpdateSLAFinal(ctx context.Context, slaDTO dto.SLA) (dto.SLA, error) {
	if slaDTO.MaxDays <= 0 {
		return dto.SLA{}, errors.New("max_days must be greater than 0")
	}

	existingSLA, err := s.slaRepository.GetSLAByStatus(ctx, "final")
	if err != nil {
		return dto.SLA{}, err
	}

	existingSLA.MaxDays = slaDTO.MaxDays
	if slaDTO.Description != "" {
		existingSLA.Description = slaDTO.Description
	}

	updatedSLA, err := s.slaRepository.UpdateSLA(ctx, existingSLA)
	if err != nil {
		return dto.SLA{}, err
	}

	return dto.SLA{
		ID:          updatedSLA.ID,
		Status:      updatedSLA.Status,
		MaxDays:     updatedSLA.MaxDays,
		Description: updatedSLA.Description,
		UpdatedAt:   updatedSLA.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *slaService) ExportApplications(ctx context.Context, request dto.ExportRequest) ([]byte, string, error) {
	applications, err := s.slaRepository.GetApplicationsForExport(ctx, request.ApplicationType)
	if err != nil {
		return nil, "", err
	}

	if request.FileType == "pdf" {
		return s.generateApplicationsPDF(applications, request.ApplicationType)
	}
	return s.generateApplicationsExcel(applications, request.ApplicationType)
}

func (s *slaService) ExportPrograms(ctx context.Context, request dto.ExportRequest) ([]byte, string, error) {
	programs, err := s.slaRepository.GetProgramsForExport(ctx, request.ApplicationType)
	if err != nil {
		return nil, "", err
	}

	if request.FileType == "pdf" {
		return s.generateProgramsPDF(programs, request.ApplicationType)
	}
	return s.generateProgramsExcel(programs, request.ApplicationType)
}

// Helper functions for PDF generation
func (s *slaService) generateApplicationsPDF(applications []model.Application, appType string) ([]byte, string, error) {
	// Simple text-based PDF content
	content := fmt.Sprintf("LAPORAN PENGAJUAN UMKM\n")
	content += fmt.Sprintf("Tanggal: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
	content += fmt.Sprintf("Tipe: %s\n\n", appType)
	content += fmt.Sprintf("Total Pengajuan: %d\n\n", len(applications))
	content += "=====================================\n\n"

	for i, app := range applications {
		content += fmt.Sprintf("%d. UMKM: %s\n", i+1, app.UMKM.BusinessName)
		content += fmt.Sprintf("   Program: %s\n", app.Program.Title)
		content += fmt.Sprintf("   Status: %s\n", app.Status)
		content += fmt.Sprintf("   Tanggal Pengajuan: %s\n\n", app.SubmittedAt.Format("2006-01-02"))
	}

	filename := fmt.Sprintf("applications_%s_%s.txt", appType, time.Now().Format("20060102_150405"))
	return []byte(content), filename, nil
}

func (s *slaService) generateProgramsPDF(programs []model.Program, appType string) ([]byte, string, error) {
	content := fmt.Sprintf("LAPORAN PROGRAM UMKM\n")
	content += fmt.Sprintf("Tanggal: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
	content += fmt.Sprintf("Tipe: %s\n\n", appType)
	content += fmt.Sprintf("Total Program: %d\n\n", len(programs))
	content += "=====================================\n\n"

	for i, prog := range programs {
		content += fmt.Sprintf("%d. Program: %s\n", i+1, prog.Title)
		content += fmt.Sprintf("   Provider: %s\n", prog.Provider)
		content += fmt.Sprintf("   Tipe: %s\n", prog.Type)
		content += fmt.Sprintf("   Status: %v\n\n", prog.IsActive)
	}

	filename := fmt.Sprintf("programs_%s_%s.txt", appType, time.Now().Format("20060102_150405"))
	return []byte(content), filename, nil
}

// Helper functions for Excel generation (CSV format)
func (s *slaService) generateApplicationsExcel(applications []model.Application, appType string) ([]byte, string, error) {
	content := "No,UMKM,Program,Status,Tanggal Pengajuan\n"

	for i, app := range applications {
		content += fmt.Sprintf("%d,%s,%s,%s,%s\n",
			i+1,
			app.UMKM.BusinessName,
			app.Program.Title,
			app.Status,
			app.SubmittedAt.Format("2006-01-02"))
	}

	filename := fmt.Sprintf("applications_%s_%s.csv", appType, time.Now().Format("20060102_150405"))
	return []byte(content), filename, nil
}

func (s *slaService) generateProgramsExcel(programs []model.Program, appType string) ([]byte, string, error) {
	content := "No,Program,Provider,Tipe,Status Aktif\n"

	for i, prog := range programs {
		status := "Tidak Aktif"
		if prog.IsActive {
			status = "Aktif"
		}
		content += fmt.Sprintf("%d,%s,%s,%s,%s\n",
			i+1,
			prog.Title,
			prog.Provider,
			prog.Type,
			status)
	}

	filename := fmt.Sprintf("programs_%s_%s.csv", appType, time.Now().Format("20060102_150405"))
	return []byte(content), filename, nil
}
