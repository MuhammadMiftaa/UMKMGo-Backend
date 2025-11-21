package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils/constant"
)

type ApplicationsService interface {
	GetAllApplications(ctx context.Context, filterType string) ([]dto.Applications, error)
	GetApplicationByID(ctx context.Context, id int) (dto.Applications, error)

	// Screening Decisions
	ScreeningApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error)
	ScreeningReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)
	ScreeningRevise(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)

	// Final Decisions
	FinalApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error)
	FinalReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)
}

type applicationsService struct {
	applicationRepository  repository.ApplicationsRepository
	userRepository         repository.UsersRepository
	notificationRepository repository.NotificationRepository
}

func NewApplicationsService(applicationRepo repository.ApplicationsRepository, userRepo repository.UsersRepository, notificationRepo repository.NotificationRepository) ApplicationsService {
	return &applicationsService{
		applicationRepository:  applicationRepo,
		userRepository:         userRepo,
		notificationRepository: notificationRepo,
	}
}

func (s *applicationsService) GetAllApplications(ctx context.Context, filterType string) ([]dto.Applications, error) {
	applications, err := s.applicationRepository.GetAllApplications(ctx, filterType)
	if err != nil {
		return nil, err
	}

	var applicationsDTO []dto.Applications
	for _, app := range applications {
		// Get documents
		documents, _ := s.applicationRepository.GetApplicationDocuments(ctx, app.ID)
		var documentsDTO []dto.ApplicationDocuments
		for _, doc := range documents {
			documentsDTO = append(documentsDTO, dto.ApplicationDocuments{
				ID:            doc.ID,
				ApplicationID: doc.ApplicationID,
				Type:          doc.Type,
				File:          doc.File,
				CreatedAt:     doc.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:     doc.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		// Get histories
		histories, _ := s.applicationRepository.GetApplicationHistories(ctx, app.ID)
		var historiesDTO []dto.ApplicationHistories
		for _, hist := range histories {
			historiesDTO = append(historiesDTO, dto.ApplicationHistories{
				ID:             hist.ID,
				ApplicationID:  hist.ApplicationID,
				Status:         hist.Status,
				Notes:          hist.Notes,
				ActionedAt:     hist.ActionedAt.Format("2006-01-02 15:04:05"),
				ActionedBy:     hist.ActionedBy,
				ActionedByName: hist.User.Name,
				CreatedAt:      hist.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:      hist.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		applicationDTO := dto.Applications{
			ID:          app.ID,
			UMKMID:      app.UMKMID,
			ProgramID:   app.ProgramID,
			Type:        app.Type,
			Status:      app.Status,
			SubmittedAt: app.SubmittedAt.Format("2006-01-02 15:04:05"),
			ExpiredAt:   app.ExpiredAt.Format("2006-01-02 15:04:05"),
			CreatedAt:   app.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   app.UpdatedAt.Format("2006-01-02 15:04:05"),
			Documents:   documentsDTO,
			Histories:   historiesDTO,
			Program: &dto.Programs{
				ID:                  app.Program.ID,
				Title:               app.Program.Title,
				Type:                app.Program.Type,
				Location:            app.Program.Location,
				ApplicationDeadline: app.Program.ApplicationDeadline,
			},
			UMKM: &dto.UMKMWeb{
				ID:           app.UMKM.ID,
				BusinessName: app.UMKM.BusinessName,
				NIK:          app.UMKM.NIK,
				Address:      app.UMKM.Address,
				District:     app.UMKM.District,
				Subdistrict:  app.UMKM.Subdistrict,
				User: dto.User{
					ID:    app.UMKM.User.ID,
					Name:  app.UMKM.User.Name,
					Email: app.UMKM.User.Email,
				},
				Province: dto.Province{
					ID:   app.UMKM.City.Province.ID,
					Name: app.UMKM.City.Province.Name,
				},
				City: dto.City{
					ID:   app.UMKM.City.ID,
					Name: app.UMKM.City.Name,
				},
			},
		}
		applicationsDTO = append(applicationsDTO, applicationDTO)
	}

	return applicationsDTO, nil
}

func (s *applicationsService) GetApplicationByID(ctx context.Context, id int) (dto.Applications, error) {
	application, err := s.applicationRepository.GetApplicationByID(ctx, id)
	if err != nil {
		return dto.Applications{}, err
	}

	// Get documents
	documents, _ := s.applicationRepository.GetApplicationDocuments(ctx, application.ID)
	var documentsDTO []dto.ApplicationDocuments
	for _, doc := range documents {
		documentsDTO = append(documentsDTO, dto.ApplicationDocuments{
			ID:            doc.ID,
			ApplicationID: doc.ApplicationID,
			Type:          doc.Type,
			File:          doc.File,
			CreatedAt:     doc.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     doc.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// Get histories
	histories, _ := s.applicationRepository.GetApplicationHistories(ctx, application.ID)
	var historiesDTO []dto.ApplicationHistories
	for _, hist := range histories {
		historiesDTO = append(historiesDTO, dto.ApplicationHistories{
			ID:             hist.ID,
			ApplicationID:  hist.ApplicationID,
			Status:         hist.Status,
			Notes:          hist.Notes,
			ActionedAt:     hist.ActionedAt.Format("2006-01-02 15:04:05"),
			ActionedBy:     hist.ActionedBy,
			ActionedByName: hist.User.Name,
			CreatedAt:      hist.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      hist.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return dto.Applications{
		ID:          application.ID,
		UMKMID:      application.UMKMID,
		ProgramID:   application.ProgramID,
		Type:        application.Type,
		Status:      application.Status,
		SubmittedAt: application.SubmittedAt.Format("2006-01-02 15:04:05"),
		ExpiredAt:   application.ExpiredAt.Format("2006-01-02 15:04:05"),
		CreatedAt:   application.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   application.UpdatedAt.Format("2006-01-02 15:04:05"),
		Documents:   documentsDTO,
		Histories:   historiesDTO,
		Program: &dto.Programs{
			ID:                  application.Program.ID,
			Title:               application.Program.Title,
			Type:                application.Program.Type,
			Location:            application.Program.Location,
			ApplicationDeadline: application.Program.ApplicationDeadline,
		},
		UMKM: &dto.UMKMWeb{
			ID:           application.UMKM.ID,
			BusinessName: application.UMKM.BusinessName,
			NIK:          application.UMKM.NIK,
			Address:      application.UMKM.Address,
			District:     application.UMKM.District,
			Subdistrict:  application.UMKM.Subdistrict,
			User: dto.User{
				ID:    application.UMKM.User.ID,
				Name:  application.UMKM.User.Name,
				Email: application.UMKM.User.Email,
			},
			Province: dto.Province{
				ID:   application.UMKM.City.Province.ID,
				Name: application.UMKM.City.Province.Name,
			},
			City: dto.City{
				ID:   application.UMKM.City.ID,
				Name: application.UMKM.City.Name,
			},
		},
	}, nil
}

func (s *applicationsService) ScreeningApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error) {
	// Get application
	application, err := s.applicationRepository.GetApplicationByID(ctx, applicationID)
	if err != nil {
		return dto.Applications{}, err
	}

	// Validate status
	if application.Status != "screening" {
		return dto.Applications{}, errors.New("application must be in screening status")
	}

	// Update status to final
	application.Status = "final"
	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, application)
	if err != nil {
		return dto.Applications{}, err
	}

	// Create history
	history := model.ApplicationHistory{
		ApplicationID: applicationID,
		Status:        "approve_by_admin_screening",
		Notes:         "Approved by admin screening",
		ActionedBy:    userID,
	}
	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
		return dto.Applications{}, err
	}

	// Create notification
	metadata, err := json.Marshal(map[string]any{})
	if err != nil {
		return dto.Applications{}, err
	}

	notification := model.Notification{
		UMKMID:        application.UMKMID,
		Title:         constant.NotificationTitleApproved,
		Message:       constant.NotificationMessageApproved,
		IsRead:        false,
		ApplicationID: &updatedApplication.ID,
		Type:          constant.NotificationApproved,
		Metadata:      string(metadata),
	}
	if err := s.notificationRepository.CreateNotification(ctx, notification); err != nil {
		return dto.Applications{}, err
	}

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}

func (s *applicationsService) ScreeningReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error) {
	// Validate notes
	if decision.Notes == "" {
		return dto.Applications{}, errors.New("notes are required for rejection")
	}

	// Get application
	application, err := s.applicationRepository.GetApplicationByID(ctx, decision.ApplicationID)
	if err != nil {
		return dto.Applications{}, err
	}

	// Validate status
	if application.Status != "screening" {
		return dto.Applications{}, errors.New("application must be in screening status")
	}

	// Update status to rejected
	application.Status = "rejected"
	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, application)
	if err != nil {
		return dto.Applications{}, err
	}

	// Create history
	history := model.ApplicationHistory{
		ApplicationID: decision.ApplicationID,
		Status:        "reject_by_admin_screening",
		Notes:         decision.Notes,
		ActionedBy:    userID,
	}
	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
		return dto.Applications{}, err
	}

	// Create notification
	metadata, err := json.Marshal(map[string]any{})
	if err != nil {
		return dto.Applications{}, err
	}

	notification := model.Notification{
		UMKMID:        application.UMKMID,
		Title:         constant.NotificationTitleRejected,
		Message:       fmt.Sprintf(constant.NotificationMessageRejected, decision.Notes),
		IsRead:        false,
		ApplicationID: &updatedApplication.ID,
		Type:          constant.NotificationRejected,
		Metadata:      string(metadata),
	}
	if err := s.notificationRepository.CreateNotification(ctx, notification); err != nil {
		return dto.Applications{}, err
	}

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}

func (s *applicationsService) ScreeningRevise(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error) {
	// Validate notes
	if decision.Notes == "" {
		return dto.Applications{}, errors.New("notes are required for revision")
	}

	// Get application
	application, err := s.applicationRepository.GetApplicationByID(ctx, decision.ApplicationID)
	if err != nil {
		return dto.Applications{}, err
	}

	// Validate status
	if application.Status != "screening" {
		return dto.Applications{}, errors.New("application must be in screening status")
	}

	// Update status to revised
	application.Status = "revised"
	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, application)
	if err != nil {
		return dto.Applications{}, err
	}

	// Create history
	history := model.ApplicationHistory{
		ApplicationID: decision.ApplicationID,
		Status:        "revise",
		Notes:         decision.Notes,
		ActionedBy:    userID,
	}
	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
		return dto.Applications{}, err
	}

	// Create notification
	metadata, err := json.Marshal(map[string]any{})
	if err != nil {
		return dto.Applications{}, err
	}

	notification := model.Notification{
		UMKMID:        application.UMKMID,
		Title:         constant.NotificationTitleRevised,
		Message:       fmt.Sprintf(constant.NotificationMessageRevised, decision.Notes),
		IsRead:        false,
		ApplicationID: &updatedApplication.ID,
		Type:          constant.NotificationRevised,
		Metadata:      string(metadata),
	}
	if err := s.notificationRepository.CreateNotification(ctx, notification); err != nil {
		return dto.Applications{}, err
	}

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}

func (s *applicationsService) FinalApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error) {
	// Get application
	application, err := s.applicationRepository.GetApplicationByID(ctx, applicationID)
	if err != nil {
		return dto.Applications{}, err
	}

	// Validate status
	if application.Status != "final" {
		return dto.Applications{}, errors.New("application must be in final status")
	}

	// Update status to approved
	application.Status = "approved"
	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, application)
	if err != nil {
		return dto.Applications{}, err
	}

	// Create history
	history := model.ApplicationHistory{
		ApplicationID: applicationID,
		Status:        "approve_by_admin_vendor",
		Notes:         "Approved by admin vendor",
		ActionedBy:    userID,
	}
	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
		return dto.Applications{}, err
	}

	// Create notification
	metadata, err := json.Marshal(map[string]any{})
	if err != nil {
		return dto.Applications{}, err
	}

	notification := model.Notification{
		UMKMID:        application.UMKMID,
		Title:         constant.NotificationTitleApproved,
		Message:       constant.NotificationMessageApproved,
		IsRead:        false,
		ApplicationID: &updatedApplication.ID,
		Type:          constant.NotificationFinalApproved,
		Metadata:      string(metadata),
	}
	if err := s.notificationRepository.CreateNotification(ctx, notification); err != nil {
		return dto.Applications{}, err
	}

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}

func (s *applicationsService) FinalReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error) {
	// Validate notes
	if decision.Notes == "" {
		return dto.Applications{}, errors.New("notes are required for rejection")
	}

	// Get application
	application, err := s.applicationRepository.GetApplicationByID(ctx, decision.ApplicationID)
	if err != nil {
		return dto.Applications{}, err
	}

	// Validate status
	if application.Status != "final" {
		return dto.Applications{}, errors.New("application must be in final status")
	}

	// Update status to rejected
	application.Status = "rejected"
	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, application)
	if err != nil {
		return dto.Applications{}, err
	}

	// Create history
	history := model.ApplicationHistory{
		ApplicationID: decision.ApplicationID,
		Status:        "reject_by_admin_vendor",
		Notes:         decision.Notes,
		ActionedBy:    userID,
	}
	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
		return dto.Applications{}, err
	}

	// Create notification
	metadata, err := json.Marshal(map[string]any{})
	if err != nil {
		return dto.Applications{}, err
	}

	notification := model.Notification{
		UMKMID:        application.UMKMID,
		Title:         constant.NotificationTitleRejected,
		Message:       fmt.Sprintf(constant.NotificationMessageRejected, decision.Notes),
		IsRead:        false,
		ApplicationID: &updatedApplication.ID,
		Type:          constant.NotificationFinalRejected,
		Metadata:      string(metadata),
	}
	if err := s.notificationRepository.CreateNotification(ctx, notification); err != nil {
		return dto.Applications{}, err
	}

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}
