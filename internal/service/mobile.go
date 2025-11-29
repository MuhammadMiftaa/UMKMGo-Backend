package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/log"
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/config/vault"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"
	"UMKMGo-backend/internal/utils/constant"
)

type MobileService interface {
	// Dashboard
	GetDashboard(ctx context.Context, userID int) (dto.DashboardData, error)

	// Programs
	GetTrainingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetCertificationPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetFundingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetProgramDetail(ctx context.Context, id int) (dto.ProgramDetailMobile, error)

	// UMKM Profile
	GetUMKMProfile(ctx context.Context, userID int) (dto.UMKMProfile, error)
	UpdateUMKMProfile(ctx context.Context, userID int, request dto.UpdateUMKMProfile) (dto.UMKMProfile, error)

	// Documents
	GetUMKMDocuments(ctx context.Context, userID int) ([]dto.UMKMDocument, error)
	UploadDocument(ctx context.Context, userID int, doc dto.UploadDocumentRequest) error

	// Applications
	CreateTrainingApplication(ctx context.Context, userID int, request dto.CreateApplicationTraining) error
	CreateCertificationApplication(ctx context.Context, userID int, request dto.CreateApplicationCertification) error
	CreateFundingApplication(ctx context.Context, userID int, request dto.CreateApplicationFunding) error
	GetApplicationList(ctx context.Context, userID int) ([]dto.ApplicationListMobile, error)
	GetApplicationDetail(ctx context.Context, id int) (dto.ApplicationDetailMobile, error)
	GetUMKMProfileWithDecryption(ctx context.Context, userID int, purpose string) (dto.UMKMProfile, error)
	ReviseApplication(ctx context.Context, userID, applicationID int, documents []dto.UploadDocumentRequest) error

	// Notifications
	GetNotificationsByUMKMID(ctx context.Context, umkmID int) ([]dto.NotificationResponse, error)
	GetUnreadCount(ctx context.Context, umkmID int) (int64, error)
	MarkNotificationsAsRead(ctx context.Context, umkmID int, notificationIDs int) error
	MarkAllNotificationsAsRead(ctx context.Context, umkmID int) error

	// News
	GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]dto.NewsListMobile, int64, error)
	GetNewsDetail(ctx context.Context, slug string) (dto.NewsDetailMobile, error)
}

type mobileService struct {
	mobileRepo       repository.MobileRepository
	programRepo      repository.ProgramsRepository
	notificationRepo repository.NotificationRepository
	vaultLogRepo     repository.VaultDecryptLogRepository
	applicationRepo  repository.ApplicationsRepository
	slaRepo          repository.SLARepository
	minio            *storage.MinIOManager
}

func NewMobileService(mobileRepo repository.MobileRepository, programRepo repository.ProgramsRepository, notificationRepo repository.NotificationRepository, vaultLogRepo repository.VaultDecryptLogRepository, applicationRepo repository.ApplicationsRepository, slaRepo repository.SLARepository, minio *storage.MinIOManager) MobileService {
	return &mobileService{
		mobileRepo:       mobileRepo,
		programRepo:      programRepo,
		notificationRepo: notificationRepo,
		vaultLogRepo:     vaultLogRepo,
		applicationRepo:  applicationRepo,
		slaRepo:          slaRepo,
		minio:            minio,
	}
}

// Dashboard
func (s *mobileService) GetDashboard(ctx context.Context, userID int) (dto.DashboardData, error) {
	var res dto.DashboardData
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return dto.DashboardData{}, err
	}

	// Get notifications count
	unreadNotifications, err := s.notificationRepo.GetUnreadCount(ctx, umkm.ID)
	if err != nil {
		return dto.DashboardData{}, err
	}

	// Decrypt Kartu Number
	decryptedKartu, err := vault.DecryptTransit(ctx, env.Cfg.Vault.TransitPath, env.Cfg.Vault.KartuEncryptionKey, umkm.KartuNumber)
	if err != nil {
		return dto.DashboardData{}, errors.New("failed to decrypt Kartu Number")
	}

	// Get approved applications count
	applications, err := s.applicationRepo.GetApplicationsByUMKMID(ctx, userID)
	if err != nil {
		return dto.DashboardData{}, err
	}

	var totalApproved int
	for _, app := range applications {
		if app.Status == constant.ApplicationStatusApproved {
			totalApproved++
		}
	}

	res.Name = umkm.User.Name
	res.NotificationsCount = int(unreadNotifications)
	res.KartuType = umkm.KartuType
	res.KartuNumber = string(decryptedKartu)
	res.QRCode = umkm.QRCode
	res.TotalApplications = len(applications)
	res.ApprovedApplications = totalApproved
	return res, nil
}

// Programs
func (s *mobileService) GetTrainingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error) {
	programs, err := s.mobileRepo.GetProgramsByType(ctx, "training")
	if err != nil {
		return nil, err
	}

	return s.mapProgramsToDTO(programs), nil
}

func (s *mobileService) GetCertificationPrograms(ctx context.Context) ([]dto.ProgramListMobile, error) {
	programs, err := s.mobileRepo.GetProgramsByType(ctx, "certification")
	if err != nil {
		return nil, err
	}

	return s.mapProgramsToDTO(programs), nil
}

func (s *mobileService) GetFundingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error) {
	programs, err := s.mobileRepo.GetProgramsByType(ctx, "funding")
	if err != nil {
		return nil, err
	}

	return s.mapProgramsToDTO(programs), nil
}

func (s *mobileService) GetProgramDetail(ctx context.Context, id int) (dto.ProgramDetailMobile, error) {
	program, err := s.mobileRepo.GetProgramDetailByID(ctx, id)
	if err != nil {
		return dto.ProgramDetailMobile{}, err
	}

	// Get benefits and requirements
	benefits, err := s.programRepo.GetProgramBenefits(ctx, program.ID)
	if err != nil {
		return dto.ProgramDetailMobile{}, err
	}
	requirements, err := s.programRepo.GetProgramRequirements(ctx, program.ID)
	if err != nil {
		return dto.ProgramDetailMobile{}, err
	}

	var benefitNames []string
	for _, b := range benefits {
		benefitNames = append(benefitNames, b.Name)
	}
	if benefitNames == nil {
		benefitNames = []string{}
	}

	var requirementNames []string
	for _, r := range requirements {
		requirementNames = append(requirementNames, r.Name)
	}
	if requirementNames == nil {
		requirementNames = []string{}
	}

	return dto.ProgramDetailMobile{
		ProgramListMobile: s.mapProgramToDTO(program),
		Benefits:          benefitNames,
		Requirements:      requirementNames,
	}, nil
}

// UMKM Profile
func (s *mobileService) GetUMKMProfile(ctx context.Context, userID int) (dto.UMKMProfile, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Decrypt NIK
	decryptedNIK, err := vault.DecryptTransit(ctx, env.Cfg.Vault.TransitPath, env.Cfg.Vault.NIKEncryptionKey, umkm.NIK)
	if err != nil {
		return dto.UMKMProfile{}, errors.New("failed to decrypt NIK")
	}

	// Decrypt Kartu Number
	decryptedKartu, err := vault.DecryptTransit(ctx, env.Cfg.Vault.TransitPath, env.Cfg.Vault.KartuEncryptionKey, umkm.KartuNumber)
	if err != nil {
		return dto.UMKMProfile{}, errors.New("failed to decrypt Kartu Number")
	}

	return dto.UMKMProfile{
		ID:             umkm.ID,
		UserID:         umkm.UserID,
		BusinessName:   umkm.BusinessName,
		NIK:            string(decryptedNIK),
		Gender:         umkm.Gender,
		BirthDate:      umkm.BirthDate.Format("2006-01-02"),
		Phone:          umkm.Phone,
		Address:        umkm.Address,
		ProvinceID:     umkm.ProvinceID,
		CityID:         umkm.CityID,
		District:       umkm.District,
		Subdistrict:    umkm.Subdistrict,
		PostalCode:     umkm.PostalCode,
		NIB:            umkm.NIB,
		NPWP:           umkm.NPWP,
		RevenueRecord:  umkm.RevenueRecord,
		BusinessPermit: umkm.BusinessPermit,
		KartuType:      umkm.KartuType,
		Photo:          umkm.Photo,
		KartuNumber:    string(decryptedKartu),
		Province: dto.Province{
			ID:   umkm.Province.ID,
			Name: umkm.Province.Name,
		},
		City: dto.City{
			ID:   umkm.City.ID,
			Name: umkm.City.Name,
		},
		User: dto.User{
			ID:    umkm.User.ID,
			Name:  umkm.User.Name,
			Email: umkm.User.Email,
		},
	}, nil
}

func (s *mobileService) UpdateUMKMProfile(ctx context.Context, userID int, request dto.UpdateUMKMProfile) (dto.UMKMProfile, error) {
	// Get existing UMKM
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Parse birth date
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		return dto.UMKMProfile{}, errors.New("invalid birth date format, use YYYY-MM-DD")
	}

	// If photo is provided, upload to MinIO
	if request.Photo != "" {
		if !(strings.HasPrefix(request.Photo, "http") || strings.HasPrefix(request.Photo, "https")) {
			res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
				Base64Data: request.Photo,
				BucketName: storage.UMKMBucket,
				Prefix:     utils.GenerateFileName(request.Name, "photo_profile_"),
				Validation: storage.CreateImageValidationConfig(),
			})
			if err != nil {
				return dto.UMKMProfile{}, err
			}

			if umkm.Photo != "" {
				objectName := storage.ExtractObjectNameFromURL(umkm.Photo)
				if objectName != "" {
					if deleteErr := s.minio.DeleteFile(ctx, storage.ProgramBucket, objectName); deleteErr != nil {
						return dto.UMKMProfile{}, fmt.Errorf("failed to delete old provider logo: %w", deleteErr)
					}
				}
			}

			umkm.Photo = res.URL
		}
	}

	// Update fields
	umkm.BusinessName = request.BusinessName
	umkm.Gender = request.Gender
	umkm.BirthDate = birthDate
	umkm.Address = request.Address
	umkm.ProvinceID = request.ProvinceID
	umkm.CityID = request.CityID
	umkm.District = request.District
	umkm.PostalCode = request.PostalCode
	umkm.User.Name = request.Name

	// Update in database
	_, err = s.mobileRepo.UpdateUMKMProfile(ctx, umkm)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Return updated profile
	return s.GetUMKMProfile(ctx, userID)
}

func (s *mobileService) GetUMKMDocuments(ctx context.Context, userID int) ([]dto.UMKMDocument, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var documents []dto.UMKMDocument
	if umkm.NIB != "" {
		documents = append(documents, dto.UMKMDocument{
			DocumentType: constant.DocumentTypeNib,
			DocumentURL:  umkm.NIB,
		})
	}
	if umkm.NPWP != "" {
		documents = append(documents, dto.UMKMDocument{
			DocumentType: constant.DocumentTypeNPWP,
			DocumentURL:  umkm.NPWP,
		})
	}
	if umkm.RevenueRecord != "" {
		documents = append(documents, dto.UMKMDocument{
			DocumentType: constant.DocumentTypeRevenueRecord,
			DocumentURL:  umkm.RevenueRecord,
		})
	}
	if umkm.BusinessPermit != "" {
		documents = append(documents, dto.UMKMDocument{
			DocumentType: constant.DocumentTypeBusinessPermit,
			DocumentURL:  umkm.BusinessPermit,
		})
	}

	return documents, nil
}

func (s *mobileService) UploadDocument(ctx context.Context, userID int, doc dto.UploadDocumentRequest) error {
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return err
	}

	// Map document type to database field and file prefix
	docTypeConfig := map[string]struct {
		field  string
		prefix string
		oldURL string
	}{
		"nib": {
			field:  "nib",
			prefix: fmt.Sprintf("nib_%d_", umkm.ID),
			oldURL: umkm.NIB,
		},
		"npwp": {
			field:  "npwp",
			prefix: fmt.Sprintf("npwp_%d_", umkm.ID),
			oldURL: umkm.NPWP,
		},
		"revenue-record": {
			field:  "revenue_record",
			prefix: fmt.Sprintf("revenue_%d_", umkm.ID),
			oldURL: umkm.RevenueRecord,
		},
		"business-permit": {
			field:  "business_permit",
			prefix: fmt.Sprintf("permit_%d_", umkm.ID),
			oldURL: umkm.BusinessPermit,
		},
	}

	config, exists := docTypeConfig[doc.Type]
	if !exists {
		return errors.New("invalid document type")
	}

	docURL := doc.Document

	if !(strings.HasPrefix(doc.Document, "http") || strings.HasPrefix(doc.Document, "https")) {
		// Upload to MinIO
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: doc.Document,
			BucketName: storage.UMKMBucket,
			Prefix:     config.prefix,
			Validation: storage.CreateImageValidationConfig(),
		})
		if err != nil {
			return err
		}

		// Delete old document if exists
		if config.oldURL != "" {
			oldObjectName := storage.ExtractObjectNameFromURL(config.oldURL)
			if oldObjectName != "" {
				s.minio.DeleteFile(ctx, storage.UMKMBucket, oldObjectName)
			}
		}

		docURL = res.URL
	}

	// Update document in database
	return s.mobileRepo.UpdateUMKMDocument(ctx, umkm.ID, config.field, docURL)
}

// Applications
// Training Application
func (s *mobileService) CreateTrainingApplication(ctx context.Context, userID int, request dto.CreateApplicationTraining) error {
	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return err
	}

	if program.Type != "training" {
		return errors.New("program type must be training")
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, userID, request.ProgramID) {
		return errors.New("you have already applied for this program")
	}

	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return errors.New("UMKM profile not found, please complete your profile first")
	}

	screeningExpiredAt, err := s.slaRepo.GetSLAByStatus(ctx, "screening")
	if err != nil {
		return err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "training",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, screeningExpiredAt.MaxDays),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return err
	}

	// Create training-specific data
	trainingApp := model.TrainingApplication{
		ApplicationID:      createdApp.ID,
		Motivation:         request.Motivation,
		BusinessExperience: request.BusinessExperience,
		LearningObjectives: request.LearningObjectives,
		AvailabilityNotes:  request.AvailabilityNotes,
	}

	if err := s.mobileRepo.CreateTrainingApplication(ctx, trainingApp); err != nil {
		return err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, umkm.UserID, "submit", "Training application submitted"); err != nil {
		return err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return err
	}

	// Process and save documents
	go s.processAndSaveDocuments(ctx, createdApp.ID, request.Documents)

	return nil
}

// Certification Application
func (s *mobileService) CreateCertificationApplication(ctx context.Context, userID int, request dto.CreateApplicationCertification) error {
	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return err
	}

	if program.Type != "certification" {
		return errors.New("program type must be certification")
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, userID, request.ProgramID) {
		return errors.New("you have already applied for this program")
	}

	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return errors.New("UMKM profile not found, please complete your profile first")
	}

	screeningExpiredAt, err := s.slaRepo.GetSLAByStatus(ctx, "screening")
	if err != nil {
		return err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "certification",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, screeningExpiredAt.MaxDays),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return err
	}

	// Create certification-specific data
	certApp := model.CertificationApplication{
		ApplicationID:       createdApp.ID,
		BusinessSector:      request.BusinessSector,
		ProductOrService:    request.ProductOrService,
		BusinessDescription: request.BusinessDescription,
		YearsOperating:      request.YearsOperating,
		CurrentStandards:    request.CurrentStandards,
		CertificationGoals:  request.CertificationGoals,
	}

	if err := s.mobileRepo.CreateCertificationApplication(ctx, certApp); err != nil {
		return err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, umkm.UserID, "submit", "Certification application submitted"); err != nil {
		return err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return err
	}

	// Process and save documents
	go s.processAndSaveDocuments(ctx, createdApp.ID, request.Documents)

	return nil
}

// Funding Application
func (s *mobileService) CreateFundingApplication(ctx context.Context, userID int, request dto.CreateApplicationFunding) error {
	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return err
	}

	// Validate requested amount
	if program.MinAmount != nil && request.RequestedAmount < *program.MinAmount {
		return fmt.Errorf("requested amount must be at least %.2f", *program.MinAmount)
	}
	if program.MaxAmount != nil && request.RequestedAmount > *program.MaxAmount {
		return fmt.Errorf("requested amount cannot exceed %.2f", *program.MaxAmount)
	}

	// Validate tenure
	if program.MaxTenureMonths != nil && request.RequestedTenureMonths > *program.MaxTenureMonths {
		return fmt.Errorf("requested tenure cannot exceed %d months", *program.MaxTenureMonths)
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, userID, request.ProgramID) {
		return errors.New("you have already applied for this program")
	}

	if program.Type != "funding" {
		return errors.New("program type must be funding")
	}

	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return errors.New("UMKM profile not found, please complete your profile first")
	}

	// Get screening SLA
	screeningExpiredAt, err := s.slaRepo.GetSLAByStatus(ctx, "screening")
	if err != nil {
		return err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "funding",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, screeningExpiredAt.MaxDays),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return err
	}

	// Create funding-specific data
	fundingApp := model.FundingApplication{
		ApplicationID:         createdApp.ID,
		BusinessSector:        request.BusinessSector,
		BusinessDescription:   request.BusinessDescription,
		YearsOperating:        request.YearsOperating,
		RequestedAmount:       request.RequestedAmount,
		FundPurpose:           request.FundPurpose,
		BusinessPlan:          request.BusinessPlan,
		RevenueProjection:     request.RevenueProjection,
		MonthlyRevenue:        request.MonthlyRevenue,
		RequestedTenureMonths: request.RequestedTenureMonths,
		CollateralDescription: request.CollateralDescription,
	}

	if err := s.mobileRepo.CreateFundingApplication(ctx, fundingApp); err != nil {
		return err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, umkm.UserID, "submit", "Funding application submitted"); err != nil {
		return err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return err
	}

	// Process and save documents
	go s.processAndSaveDocuments(ctx, createdApp.ID, request.Documents)

	return nil
}

func (s *mobileService) GetApplicationList(ctx context.Context, userID int) ([]dto.ApplicationListMobile, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	applications, err := s.mobileRepo.GetApplicationsByUMKMID(ctx, umkm.ID)
	if err != nil {
		return nil, err
	}

	var result []dto.ApplicationListMobile
	for _, app := range applications {
		result = append(result, dto.ApplicationListMobile{
			ID:          app.ID,
			ProgramID:   app.ProgramID,
			ProgramName: app.Program.Title,
			Type:        app.Type,
			Status:      app.Status,
			SubmittedAt: app.SubmittedAt.Format("2006-01-02 15:04:05"),
			ExpiredAt:   app.ExpiredAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}

func (s *mobileService) GetApplicationDetail(ctx context.Context, id int) (dto.ApplicationDetailMobile, error) {
	application, err := s.mobileRepo.GetApplicationDetailByID(ctx, id)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Get program benefits and requirements
	benefits, _ := s.programRepo.GetProgramBenefits(ctx, application.ProgramID)
	requirements, _ := s.programRepo.GetProgramRequirements(ctx, application.ProgramID)

	var benefitNames []string
	for _, b := range benefits {
		benefitNames = append(benefitNames, b.Name)
	}

	var requirementNames []string
	for _, r := range requirements {
		requirementNames = append(requirementNames, r.Name)
	}

	// Map documents
	var documents []dto.ApplicationDocuments
	for _, doc := range application.Documents {
		documents = append(documents, dto.ApplicationDocuments{
			ID:            doc.ID,
			ApplicationID: doc.ApplicationID,
			Type:          doc.Type,
			File:          doc.File,
			CreatedAt:     doc.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     doc.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// Map histories
	var histories []dto.ApplicationHistories
	for _, hist := range application.Histories {
		histories = append(histories, dto.ApplicationHistories{
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

	// Create detailed response with specific application data
	detail := dto.ApplicationDetailMobile{
		ID:          application.ID,
		UMKMID:      application.UMKMID,
		ProgramID:   application.ProgramID,
		Type:        application.Type,
		Status:      application.Status,
		SubmittedAt: application.SubmittedAt.Format("2006-01-02 15:04:05"),
		ExpiredAt:   application.ExpiredAt.Format("2006-01-02 15:04:05"),
		Documents:   documents,
		Histories:   histories,
		Program: dto.ProgramDetailMobile{
			ProgramListMobile: s.mapProgramToDTO(application.Program),
			Benefits:          benefitNames,
			Requirements:      requirementNames,
		},
	}

	// Add specific application data based on type
	switch application.Type {
	case "training":
		if application.TrainingApplication != nil {
			detail.TrainingData = &dto.TrainingApplicationData{
				Motivation:         application.TrainingApplication.Motivation,
				BusinessExperience: application.TrainingApplication.BusinessExperience,
				LearningObjectives: application.TrainingApplication.LearningObjectives,
				AvailabilityNotes:  application.TrainingApplication.AvailabilityNotes,
			}
		}
	case "certification":
		if application.CertificationApplication != nil {
			detail.CertificationData = &dto.CertificationApplicationData{
				BusinessSector:      application.CertificationApplication.BusinessSector,
				ProductOrService:    application.CertificationApplication.ProductOrService,
				BusinessDescription: application.CertificationApplication.BusinessDescription,
				YearsOperating:      application.CertificationApplication.YearsOperating,
				CurrentStandards:    application.CertificationApplication.CurrentStandards,
				CertificationGoals:  application.CertificationApplication.CertificationGoals,
			}
		}
	case "funding":
		if application.FundingApplication != nil {
			detail.FundingData = &dto.FundingApplicationData{
				BusinessSector:        application.FundingApplication.BusinessSector,
				BusinessDescription:   application.FundingApplication.BusinessDescription,
				YearsOperating:        application.FundingApplication.YearsOperating,
				RequestedAmount:       application.FundingApplication.RequestedAmount,
				FundPurpose:           application.FundingApplication.FundPurpose,
				BusinessPlan:          application.FundingApplication.BusinessPlan,
				RevenueProjection:     application.FundingApplication.RevenueProjection,
				MonthlyRevenue:        application.FundingApplication.MonthlyRevenue,
				RequestedTenureMonths: application.FundingApplication.RequestedTenureMonths,
				CollateralDescription: application.FundingApplication.CollateralDescription,
			}
		}
	}

	return detail, nil
}

func (s *mobileService) ReviseApplication(ctx context.Context, userID, applicationID int, documents []dto.UploadDocumentRequest) error {
	application, err := s.mobileRepo.GetApplicationDetailByID(ctx, applicationID)
	if err != nil {
		return err
	}

	if application.Status != constant.ApplicationStatusRevised {
		return errors.New("application is not in a revisable state")
	}

	documentsMap := make(map[string]string)
	for _, doc := range documents {
		documentsMap[doc.Type] = doc.Document
	}

	// Process and save documents
	go s.processAndSaveDocuments(ctx, applicationID, documentsMap)

	// Update application status to 'resubmitted'
	application.Status = constant.ApplicationStatusScreening
	application.SubmittedAt = time.Now()

	_, err = s.applicationRepo.UpdateApplication(ctx, application)
	if err != nil {
		return err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, application.ID, userID, "submit", "Application resubmitted after revision"); err != nil {
		return err
	}

	// Create notification
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, application.UMKMID)
	if err != nil {
		return err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, application.ID, constant.NotificationSubmitted, constant.NotificationTitleResubmitted, constant.NotificationMessageResubmitted); err != nil {
		return err
	}

	// Delete previous documents
	if err := s.mobileRepo.DeleteApplicationDocumentsByApplicationID(ctx, application.ID); err != nil {
		return err
	}

	return nil
}

func (s *mobileService) GetNotificationsByUMKMID(ctx context.Context, umkmID int) ([]dto.NotificationResponse, error) {
	notifications, err := s.notificationRepo.GetNotificationsByUMKMID(ctx, umkmID, 100, 0)
	if err != nil {
		return nil, err
	}

	var response []dto.NotificationResponse
	for _, n := range notifications {
		var readAt string
		if n.ReadAt != nil {
			readAt = n.ReadAt.Format("2006-01-02 15:04:05")
		}
		response = append(response, dto.NotificationResponse{
			ID:            n.ID,
			Type:          n.Type,
			Title:         n.Title,
			Message:       n.Message,
			IsRead:        n.IsRead,
			CreatedAt:     n.CreatedAt.Format("2006-01-02 15:04:05"),
			ReadAt:        &readAt,
			ApplicationID: n.ApplicationID,
		})
	}

	return response, nil
}

func (s *mobileService) GetUnreadCount(ctx context.Context, umkmID int) (int64, error) {
	count, err := s.notificationRepo.GetUnreadCount(ctx, umkmID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *mobileService) MarkNotificationsAsRead(ctx context.Context, umkmID int, notificationIDs int) error {
	return s.notificationRepo.MarkAsRead(ctx, notificationIDs, umkmID)
}

func (s *mobileService) MarkAllNotificationsAsRead(ctx context.Context, umkmID int) error {
	return s.notificationRepo.MarkAllAsRead(ctx, umkmID)
}

func (s *mobileService) GetPublishedNews(ctx context.Context, params dto.NewsQueryParams) ([]dto.NewsListMobile, int64, error) {
	news, total, err := s.mobileRepo.GetPublishedNews(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	var response []dto.NewsListMobile
	for _, n := range news {
		response = append(response, dto.NewsListMobile{
			ID:         n.ID,
			Title:      n.Title,
			Slug:       n.Slug,
			Excerpt:    n.Excerpt,
			Thumbnail:  n.Thumbnail,
			Category:   n.Category,
			AuthorName: n.Author.Name,
			ViewsCount: n.ViewsCount,
			CreatedAt:  n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response, total, nil
}

func (s *mobileService) GetNewsDetail(ctx context.Context, slug string) (dto.NewsDetailMobile, error) {
	news, err := s.mobileRepo.GetPublishedNewsBySlug(ctx, slug)
	if err != nil {
		return dto.NewsDetailMobile{}, err
	}

	// Increment views
	s.mobileRepo.IncrementViews(ctx, news.ID)

	var tags []string
	for _, tag := range news.Tags {
		tags = append(tags, tag.TagName)
	}

	return dto.NewsDetailMobile{
		ID:         news.ID,
		Title:      news.Title,
		Slug:       news.Slug,
		Content:    news.Content,
		Thumbnail:  news.Thumbnail,
		Category:   news.Category,
		AuthorName: news.Author.Name,
		ViewsCount: news.ViewsCount + 1, // Show incremented value
		CreatedAt:  news.CreatedAt.Format("2006-01-02 15:04:05"),
		Tags:       tags,
	}, nil
}

// Get UMKM Profile with decryption
func (s *mobileService) GetUMKMProfileWithDecryption(ctx context.Context, userID int, purpose string) (dto.UMKMProfile, error) {
	umkm, err := s.getUMKMWithDecryption(ctx, userID, purpose)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Map to DTO
	return dto.UMKMProfile{
		ID:           umkm.ID,
		UserID:       umkm.UserID,
		BusinessName: umkm.BusinessName,
		NIK:          umkm.NIK,
		KartuNumber:  umkm.KartuNumber,
	}, nil
}

// Helper functions
func (s *mobileService) getUMKMWithDecryption(ctx context.Context, userID int, purpose string) (model.UMKM, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByID(ctx, userID)
	if err != nil {
		return model.UMKM{}, err
	}

	// Get context info for logging
	ipAddress, userAgent, requestID := vault.GetContextInfo(ctx)

	// Decrypt NIK with logging
	decryptParams := vault.DecryptParams{
		UserID:    umkm.User.ID,
		UMKMID:    &umkm.ID,
		RecordID:  umkm.ID,
		Purpose:   purpose,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		RequestID: requestID,
	}

	umkm.NIK, err = vault.DecryptNIKWithLog(ctx, umkm.NIK, decryptParams, s.vaultLogRepo)
	if err != nil {
		return model.UMKM{}, errors.New("failed to decrypt NIK")
	}

	// Decrypt Kartu Number with logging
	umkm.KartuNumber, err = vault.DecryptKartuNumberWithLog(ctx, umkm.KartuNumber, decryptParams, s.vaultLogRepo)
	if err != nil {
		return model.UMKM{}, errors.New("failed to decrypt Kartu Number")
	}

	return umkm, nil
}

func (s *mobileService) mapProgramsToDTO(programs []model.Program) []dto.ProgramListMobile {
	var result []dto.ProgramListMobile
	for _, p := range programs {
		result = append(result, s.mapProgramToDTO(p))
	}
	return result
}

func (s *mobileService) mapProgramToDTO(p model.Program) dto.ProgramListMobile {
	return dto.ProgramListMobile{
		ID:                  p.ID,
		Title:               p.Title,
		Description:         p.Description,
		Banner:              p.Banner,
		Provider:            p.Provider,
		ProviderLogo:        p.ProviderLogo,
		Type:                p.Type,
		TrainingType:        p.TrainingType,
		Batch:               p.Batch,
		BatchStartDate:      p.BatchStartDate,
		BatchEndDate:        p.BatchEndDate,
		Location:            p.Location,
		MinAmount:           p.MinAmount,
		MaxAmount:           p.MaxAmount,
		InterestRate:        p.InterestRate,
		MaxTenureMonths:     p.MaxTenureMonths,
		ApplicationDeadline: p.ApplicationDeadline,
		IsActive:            p.IsActive,
	}
}

func (s *mobileService) processAndSaveDocuments(ctx context.Context, applicationID int, providedDocs map[string]string) {
	var appDocuments []model.ApplicationDocument
	var url string

	// Add documents from request
	for docType, docData := range providedDocs {
		if !(strings.HasPrefix(docData, "http") || strings.HasPrefix(docData, "https")) {
			res, err := s.minio.UploadFile(context.Background(), storage.UploadRequest{
				Base64Data: docData,
				BucketName: storage.ApplicationBucket,
				Prefix:     fmt.Sprintf("app_%d_%s_", applicationID, docType),
				Validation: storage.CreateImageValidationConfig(),
			})
			if err != nil {
				log.Log.Errorf("failed to upload %s document, %v, for application ID %d", docType, err, applicationID)
			}

			log.Log.Infof("uploaded %s document for application ID %d: %s", docType, applicationID, res.URL)
			url = res.URL
		} else {
			url = docData
		}

		appDocuments = append(appDocuments, model.ApplicationDocument{
			ApplicationID: applicationID,
			Type:          docType,
			File:          url,
		})
	}

	err := s.mobileRepo.CreateApplicationDocuments(ctx, appDocuments)
	if err != nil {
		log.Log.Errorf("failed to save application documents for application ID %d: %v", applicationID, err)
	}
}

func (s *mobileService) createApplicationHistory(ctx context.Context, applicationID, userID int, status, notes string) error {
	history := model.ApplicationHistory{
		ApplicationID: applicationID,
		Status:        status,
		Notes:         notes,
		ActionedBy:    &userID,
		ActionedAt:    time.Now(),
	}
	return s.mobileRepo.CreateApplicationHistory(ctx, history)
}

func (s *mobileService) createNotification(ctx context.Context, umkmID, applicationID int, notifType, title, message string) error {
	metadata, _ := json.Marshal(map[string]any{})
	notification := model.Notification{
		UMKMID:        umkmID,
		ApplicationID: &applicationID,
		Type:          notifType,
		Title:         title,
		Message:       message,
		IsRead:        false,
		Metadata:      string(metadata),
	}
	return s.notificationRepo.CreateNotification(ctx, notification)
}
