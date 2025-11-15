package service

import (
	"context"
	"errors"
	"log"

	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/types/model"
)

type ApplicationsService interface {
	GetAllApplications(ctx context.Context, filterType string) ([]dto.Applications, error)
	GetApplicationByID(ctx context.Context, id int) (dto.Applications, error)
	// CreateApplication(ctx context.Context, userID int, application dto.Applications) (dto.Applications, error)
	// UpdateApplication(ctx context.Context, id int, application dto.Applications) (dto.Applications, error)
	// DeleteApplication(ctx context.Context, id int) (dto.Applications, error)

	// Screening Decisions
	ScreeningApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error)
	ScreeningReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)
	ScreeningRevise(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)

	// Final Decisions
	FinalApprove(ctx context.Context, userID int, applicationID int) (dto.Applications, error)
	FinalReject(ctx context.Context, userID int, decision dto.ApplicationDecision) (dto.Applications, error)
}

type applicationsService struct {
	applicationRepository repository.ApplicationsRepository
	userRepository        repository.UsersRepository
}

func NewApplicationsService(applicationRepo repository.ApplicationsRepository, userRepo repository.UsersRepository) ApplicationsService {
	return &applicationsService{
		applicationRepository: applicationRepo,
		userRepository:        userRepo,
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
	log.Println("Fetching application by ID:", application)

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

// func (s *applicationsService) CreateApplication(ctx context.Context, userID int, application dto.Applications) (dto.Applications, error) {
// 	// Validation
// 	if application.ProgramID == 0 || len(application.Documents) == 0 {
// 		return dto.Applications{}, errors.New("program_id and documents are required")
// 	}

// 	// Check if program exists and is active
// 	program, err := s.applicationRepository.GetProgramByID(ctx, application.ProgramID)
// 	if err != nil {
// 		return dto.Applications{}, errors.New("program not found")
// 	}

// 	if !program.IsActive {
// 		return dto.Applications{}, errors.New("program is not active")
// 	}

// 	// Get UMKM by user ID
// 	umkm, err := s.applicationRepository.GetUMKMByUserID(ctx, userID)
// 	if err != nil {
// 		return dto.Applications{}, errors.New("UMKM data not found, please complete your profile first")
// 	}

// 	// Check if application already exists
// 	if s.applicationRepository.IsApplicationExists(ctx, umkm.ID, application.ProgramID) {
// 		return dto.Applications{}, errors.New("you have already applied for this program")
// 	}

// 	// Create application
// 	newApplication := model.Applications{
// 		UMKMID:    umkm.ID,
// 		ProgramID: application.ProgramID,
// 		Type:      program.Type,
// 		Status:    "screening",
// 	}

// 	createdApplication, err := s.applicationRepository.CreateApplication(ctx, newApplication)
// 	if err != nil {
// 		return dto.Applications{}, err
// 	}

// 	// Create documents
// 	if len(application.Documents) > 0 {
// 		var documents []model.ApplicationDocuments
// 		for _, doc := range application.Documents {
// 			documents = append(documents, model.ApplicationDocuments{
// 				ApplicationID: createdApplication.ID,
// 				Type:          doc.Type,
// 				File:          doc.File,
// 			})
// 		}
// 		if err := s.applicationRepository.CreateApplicationDocuments(ctx, documents); err != nil {
// 			return dto.Applications{}, err
// 		}
// 	}

// 	// Create history - submit
// 	history := model.ApplicationHistory{
// 		ApplicationID: createdApplication.ID,
// 		Status:        "submit",
// 		Notes:         "Application submitted",
// 		ActionedBy:    userID,
// 	}
// 	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
// 		return dto.Applications{}, err
// 	}

// 	return dto.Applications{
// 		ID:        createdApplication.ID,
// 		UMKMID:    createdApplication.UMKMID,
// 		ProgramID: createdApplication.ProgramID,
// 		Type:      createdApplication.Type,
// 		Status:    createdApplication.Status,
// 	}, nil
// }

// func (s *applicationsService) UpdateApplication(ctx context.Context, id int, application dto.Applications) (dto.Applications, error) {
// 	// Get existing application
// 	existingApplication, err := s.applicationRepository.GetApplicationByID(ctx, id)
// 	if err != nil {
// 		return dto.Applications{}, err
// 	}

// 	// Only allow update if status is 'revised'
// 	if existingApplication.Status != "revised" {
// 		return dto.Applications{}, errors.New("only applications with status 'revised' can be updated")
// 	}

// 	// Validation
// 	if len(application.Documents) == 0 {
// 		return dto.Applications{}, errors.New("documents are required")
// 	}

// 	// Update status back to screening
// 	existingApplication.Status = "screening"

// 	updatedApplication, err := s.applicationRepository.UpdateApplication(ctx, existingApplication)
// 	if err != nil {
// 		return dto.Applications{}, err
// 	}

// 	// Update documents
// 	if len(application.Documents) > 0 {
// 		// Delete old documents
// 		_ = s.applicationRepository.DeleteApplicationDocuments(ctx, id)

// 		// Create new documents
// 		var documents []model.ApplicationDocuments
// 		for _, doc := range application.Documents {
// 			documents = append(documents, model.ApplicationDocuments{
// 				ApplicationID: id,
// 				Type:          doc.Type,
// 				File:          doc.File,
// 			})
// 		}
// 		if err := s.applicationRepository.CreateApplicationDocuments(ctx, documents); err != nil {
// 			return dto.Applications{}, err
// 		}
// 	}

// 	// Create history - resubmit
// 	history := model.ApplicationHistory{
// 		ApplicationID: updatedApplication.ID,
// 		Status:        "submit",
// 		Notes:         "Application resubmitted after revision",
// 		ActionedBy:    existingApplication.UMKM.UserID,
// 	}
// 	if err := s.applicationRepository.CreateApplicationHistory(ctx, history); err != nil {
// 		return dto.Applications{}, err
// 	}

// 	return dto.Applications{
// 		ID:        updatedApplication.ID,
// 		UMKMID:    updatedApplication.UMKMID,
// 		ProgramID: updatedApplication.ProgramID,
// 		Type:      updatedApplication.Type,
// 		Status:    updatedApplication.Status,
// 	}, nil
// }

// func (s *applicationsService) DeleteApplication(ctx context.Context, id int) (dto.Applications, error) {
// 	application, err := s.applicationRepository.GetApplicationByID(ctx, id)
// 	if err != nil {
// 		return dto.Applications{}, err
// 	}

// 	deletedApplication, err := s.applicationRepository.DeleteApplication(ctx, application)
// 	if err != nil {
// 		return dto.Applications{}, err
// 	}

// 	return dto.Applications{
// 		ID:     deletedApplication.ID,
// 		Status: deletedApplication.Status,
// 	}, nil
// }

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

	return dto.Applications{
		ID:     updatedApplication.ID,
		Status: updatedApplication.Status,
	}, nil
}
