package service

import (
	"context"

	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/types/dto"
)

type DashboardService interface {
	GetUMKMByCardType(ctx context.Context) ([]dto.UMKMByCardType, error)
	GetApplicationStatusSummary(ctx context.Context) ([]dto.ApplicationStatusSummary, error)
	GetApplicationStatusDetail(ctx context.Context) ([]dto.ApplicationStatusDetail, error)
	GetApplicationByType(ctx context.Context) ([]dto.ApplicationByType, error)
}

type dashboardService struct {
	dashboardRepository repository.DashboardRepository
}

func NewDashboardService(dashboardRepo repository.DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepository: dashboardRepo,
	}
}

func (s *dashboardService) GetUMKMByCardType(ctx context.Context) ([]dto.UMKMByCardType, error) {
	results, err := s.dashboardRepository.GetUMKMByCardType(ctx)
	if err != nil {
		return nil, err
	}

	var response []dto.UMKMByCardType
	for _, result := range results {
		response = append(response, dto.UMKMByCardType{
			Name:  result["name"].(string),
			Count: result["count"].(int64),
		})
	}

	return response, nil
}

func (s *dashboardService) GetApplicationStatusSummary(ctx context.Context) ([]dto.ApplicationStatusSummary, error) {
	data, err := s.dashboardRepository.GetApplicationStatusSummary(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.ApplicationStatusSummary{
		{TotalApplications: data["total_applications"]},
		{InProcess: data["in_process"]},
		{Approved: data["approved"]},
		{Rejected: data["rejected"]},
	}

	return response, nil
}

func (s *dashboardService) GetApplicationStatusDetail(ctx context.Context) ([]dto.ApplicationStatusDetail, error) {
	data, err := s.dashboardRepository.GetApplicationStatusDetail(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.ApplicationStatusDetail{
		{Screening: data["screening"]},
		{Revised: data["revised"]},
		{Final: data["final"]},
		{Approved: data["approved"]},
		{Rejected: data["rejected"]},
	}

	return response, nil
}

func (s *dashboardService) GetApplicationByType(ctx context.Context) ([]dto.ApplicationByType, error) {
	data, err := s.dashboardRepository.GetApplicationByType(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.ApplicationByType{
		{Funding: data["funding"]},
		{Certification: data["certification"]},
		{Training: data["training"]},
	}

	return response, nil
}
