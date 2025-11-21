package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/config/vault"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"
	"UMKMGo-backend/internal/utils/constant"
)

type MobileService interface {
	// Programs
	GetTrainingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetCertificationPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetFundingPrograms(ctx context.Context) ([]dto.ProgramListMobile, error)
	GetProgramDetail(ctx context.Context, id int) (dto.ProgramDetailMobile, error)

	// UMKM Profile
	GetUMKMProfile(ctx context.Context, userID int) (dto.UMKMProfile, error)
	UpdateUMKMProfile(ctx context.Context, userID int, request dto.UpdateUMKMProfile) (dto.UMKMProfile, error)

	// Documents
	UploadNIB(ctx context.Context, userID int, document string) error
	UploadNPWP(ctx context.Context, userID int, document string) error
	UploadRevenueRecord(ctx context.Context, userID int, document string) error
	UploadBusinessPermit(ctx context.Context, userID int, document string) error

	// Applications
	CreateTrainingApplication(ctx context.Context, userID int, request dto.CreateApplicationTraining) (dto.ApplicationDetailMobile, error)
	CreateCertificationApplication(ctx context.Context, userID int, request dto.CreateApplicationCertification) (dto.ApplicationDetailMobile, error)
	CreateFundingApplication(ctx context.Context, userID int, request dto.CreateApplicationFunding) (dto.ApplicationDetailMobile, error)
	GetApplicationList(ctx context.Context, userID int) ([]dto.ApplicationListMobile, error)
	GetApplicationDetail(ctx context.Context, id int) (dto.ApplicationDetailMobile, error)
	GetUMKMProfileWithDecryption(ctx context.Context, userID int, purpose string) (dto.UMKMProfile, error)

	// Notifications
	GetNotificationsByUMKMID(ctx context.Context, umkmID int) ([]dto.NotificationResponse, error)
	GetUnreadCount(ctx context.Context, umkmID int) (int64, error)
	MarkNotificationsAsRead(ctx context.Context, umkmID int, notificationIDs []int) error
	MarkAllNotificationsAsRead(ctx context.Context, umkmID int) error
}

type mobileService struct {
	mobileRepo       repository.MobileRepository
	programRepo      repository.ProgramsRepository
	notificationRepo repository.NotificationRepository
	vaultLogRepo     repository.VaultDecryptLogRepository
	minio            *storage.MinIOManager
}

func NewMobileService(mobileRepo repository.MobileRepository, programRepo repository.ProgramsRepository, notificationRepo repository.NotificationRepository, vaultLogRepo repository.VaultDecryptLogRepository, minio *storage.MinIOManager) MobileService {
	return &mobileService{
		mobileRepo:       mobileRepo,
		programRepo:      programRepo,
		notificationRepo: notificationRepo,
		vaultLogRepo:     vaultLogRepo,
		minio:            minio,
	}
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
	benefits, _ := s.programRepo.GetProgramBenefits(ctx, program.ID)
	requirements, _ := s.programRepo.GetProgramRequirements(ctx, program.ID)

	var benefitNames []string
	for _, b := range benefits {
		benefitNames = append(benefitNames, b.Name)
	}

	var requirementNames []string
	for _, r := range requirements {
		requirementNames = append(requirementNames, r.Name)
	}

	return dto.ProgramDetailMobile{
		ProgramListMobile: s.mapProgramToDTO(program),
		Benefits:          benefitNames,
		Requirements:      requirementNames,
	}, nil
}

// UMKM Profile
func (s *mobileService) GetUMKMProfile(ctx context.Context, userID int) (dto.UMKMProfile, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
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
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Validate phone
	validPhone, err := utils.NormalizePhone(request.Phone)
	if err != nil {
		return dto.UMKMProfile{}, errors.New("invalid phone number")
	}

	// Parse birth date
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		return dto.UMKMProfile{}, errors.New("invalid birth date format, use YYYY-MM-DD")
	}

	// Update fields
	umkm.BusinessName = request.BusinessName
	umkm.Gender = request.Gender
	umkm.BirthDate = birthDate
	umkm.Phone = validPhone
	umkm.Address = request.Address
	umkm.ProvinceID = request.ProvinceID
	umkm.CityID = request.CityID
	umkm.District = request.District
	umkm.Subdistrict = request.Subdistrict
	umkm.PostalCode = request.PostalCode
	umkm.KartuType = request.KartuType

	// Update in database
	_, err = s.mobileRepo.UpdateUMKMProfile(ctx, umkm)
	if err != nil {
		return dto.UMKMProfile{}, err
	}

	// Return updated profile
	return s.GetUMKMProfile(ctx, userID)
}

// Documents
func (s *mobileService) UploadNIB(ctx context.Context, userID int, document string) error {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// Upload to MinIO
	res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
		Base64Data: document,
		BucketName: "umkmgo-documents",
		Prefix:     fmt.Sprintf("nib_%d_", umkm.ID),
	})
	if err != nil {
		return err
	}

	// Delete old document if exists
	if umkm.NIB != "" {
		oldObjectName := storage.ExtractObjectNameFromURL(umkm.NIB)
		if oldObjectName != "" {
			s.minio.DeleteFile(ctx, "umkmgo-documents", oldObjectName)
		}
	}

	return s.mobileRepo.UpdateUMKMDocument(ctx, umkm.ID, "nib", res.URL)
}

func (s *mobileService) UploadNPWP(ctx context.Context, userID int, document string) error {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}

	res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
		Base64Data: document,
		BucketName: "umkmgo-documents",
		Prefix:     fmt.Sprintf("npwp_%d_", umkm.ID),
	})
	if err != nil {
		return err
	}

	if umkm.NPWP != "" {
		oldObjectName := storage.ExtractObjectNameFromURL(umkm.NPWP)
		if oldObjectName != "" {
			s.minio.DeleteFile(ctx, "umkmgo-documents", oldObjectName)
		}
	}

	return s.mobileRepo.UpdateUMKMDocument(ctx, umkm.ID, "npwp", res.URL)
}

func (s *mobileService) UploadRevenueRecord(ctx context.Context, userID int, document string) error {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}

	res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
		Base64Data: document,
		BucketName: "umkmgo-documents",
		Prefix:     fmt.Sprintf("revenue_%d_", umkm.ID),
	})
	if err != nil {
		return err
	}

	if umkm.RevenueRecord != "" {
		oldObjectName := storage.ExtractObjectNameFromURL(umkm.RevenueRecord)
		if oldObjectName != "" {
			s.minio.DeleteFile(ctx, "umkmgo-documents", oldObjectName)
		}
	}

	return s.mobileRepo.UpdateUMKMDocument(ctx, umkm.ID, "revenue_record", res.URL)
}

func (s *mobileService) UploadBusinessPermit(ctx context.Context, userID int, document string) error {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}

	res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
		Base64Data: document,
		BucketName: "umkmgo-documents",
		Prefix:     fmt.Sprintf("permit_%d_", umkm.ID),
	})
	if err != nil {
		return err
	}

	if umkm.BusinessPermit != "" {
		oldObjectName := storage.ExtractObjectNameFromURL(umkm.BusinessPermit)
		if oldObjectName != "" {
			s.minio.DeleteFile(ctx, "umkmgo-documents", oldObjectName)
		}
	}

	return s.mobileRepo.UpdateUMKMDocument(ctx, umkm.ID, "business_permit", res.URL)
}

// Applications
// Training Application
func (s *mobileService) CreateTrainingApplication(ctx context.Context, userID int, request dto.CreateApplicationTraining) (dto.ApplicationDetailMobile, error) {
	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return dto.ApplicationDetailMobile{}, errors.New("UMKM profile not found, please complete your profile first")
	}

	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	if program.Type != "training" {
		return dto.ApplicationDetailMobile{}, errors.New("program type must be training")
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, umkm.ID, request.ProgramID) {
		return dto.ApplicationDetailMobile{}, errors.New("you have already applied for this program")
	}

	// Validate required documents based on program requirements
	requirements, _ := s.mobileRepo.GetProgramRequirements(ctx, request.ProgramID)
	requiredDocs := s.extractRequiredDocuments(requirements)

	// Validate documents
	if err := s.validateDocuments(umkm, requiredDocs, request.Documents, []string{"ktp"}); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "training",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, 30),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
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
		return dto.ApplicationDetailMobile{}, err
	}

	// Process and save documents
	if err := s.processAndSaveDocuments(ctx, createdApp.ID, umkm, request.Documents, requiredDocs); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, userID, "submit", "Training application submitted"); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	return s.GetApplicationDetail(ctx, createdApp.ID)
}

// Certification Application
func (s *mobileService) CreateCertificationApplication(ctx context.Context, userID int, request dto.CreateApplicationCertification) (dto.ApplicationDetailMobile, error) {
	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return dto.ApplicationDetailMobile{}, errors.New("UMKM profile not found, please complete your profile first")
	}

	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	if program.Type != "certification" {
		return dto.ApplicationDetailMobile{}, errors.New("program type must be certification")
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, umkm.ID, request.ProgramID) {
		return dto.ApplicationDetailMobile{}, errors.New("you have already applied for this program")
	}

	// Validate required documents
	requirements, _ := s.mobileRepo.GetProgramRequirements(ctx, request.ProgramID)
	requiredDocs := s.extractRequiredDocuments(requirements)

	// Certification requires: KTP, NIB, NPWP, Portfolio, Business Permit
	mandatoryDocs := []string{"ktp", "nib", "npwp", "portfolio"}
	if err := s.validateDocuments(umkm, requiredDocs, request.Documents, mandatoryDocs); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "certification",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, 30),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
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
		return dto.ApplicationDetailMobile{}, err
	}

	// Process and save documents
	if err := s.processAndSaveDocuments(ctx, createdApp.ID, umkm, request.Documents, requiredDocs); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, userID, "submit", "Certification application submitted"); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	return s.GetApplicationDetail(ctx, createdApp.ID)
}

// Funding Application
func (s *mobileService) CreateFundingApplication(ctx context.Context, userID int, request dto.CreateApplicationFunding) (dto.ApplicationDetailMobile, error) {
	// Get UMKM with decryption
	umkm, err := s.getUMKMWithDecryption(ctx, userID, "application_creation")
	if err != nil {
		return dto.ApplicationDetailMobile{}, errors.New("UMKM profile not found, please complete your profile first")
	}

	// Validate program
	program, err := s.mobileRepo.GetProgramByID(ctx, request.ProgramID)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	if program.Type != "funding" {
		return dto.ApplicationDetailMobile{}, errors.New("program type must be funding")
	}

	// Validate requested amount
	if program.MinAmount != nil && request.RequestedAmount < *program.MinAmount {
		return dto.ApplicationDetailMobile{}, fmt.Errorf("requested amount must be at least %.2f", *program.MinAmount)
	}
	if program.MaxAmount != nil && request.RequestedAmount > *program.MaxAmount {
		return dto.ApplicationDetailMobile{}, fmt.Errorf("requested amount cannot exceed %.2f", *program.MaxAmount)
	}

	// Validate tenure
	if program.MaxTenureMonths != nil && request.RequestedTenureMonths > *program.MaxTenureMonths {
		return dto.ApplicationDetailMobile{}, fmt.Errorf("requested tenure cannot exceed %d months", *program.MaxTenureMonths)
	}

	// Check if already applied
	if s.mobileRepo.IsApplicationExists(ctx, umkm.ID, request.ProgramID) {
		return dto.ApplicationDetailMobile{}, errors.New("you have already applied for this program")
	}

	// Validate required documents
	requirements, _ := s.mobileRepo.GetProgramRequirements(ctx, request.ProgramID)
	requiredDocs := s.extractRequiredDocuments(requirements)

	// Funding requires: KTP, NIB, NPWP, Rekening, Proposal, Financial Records
	mandatoryDocs := []string{"ktp", "nib", "npwp", "rekening", "proposal"}
	if err := s.validateDocuments(umkm, requiredDocs, request.Documents, mandatoryDocs); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create base application
	application := model.Application{
		UMKMID:      umkm.ID,
		ProgramID:   request.ProgramID,
		Type:        "funding",
		Status:      "screening",
		SubmittedAt: time.Now(),
		ExpiredAt:   time.Now().AddDate(0, 0, 30),
	}

	createdApp, err := s.mobileRepo.CreateApplication(ctx, application)
	if err != nil {
		return dto.ApplicationDetailMobile{}, err
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
		return dto.ApplicationDetailMobile{}, err
	}

	// Process and save documents
	if err := s.processAndSaveDocuments(ctx, createdApp.ID, umkm, request.Documents, requiredDocs); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create history
	if err := s.createApplicationHistory(ctx, createdApp.ID, userID, "submit", "Funding application submitted"); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	// Create notification
	if err := s.createNotification(ctx, umkm.ID, createdApp.ID, constant.NotificationSubmitted, constant.NotificationTitleSubmitted, constant.NotificationMessageSubmitted); err != nil {
		return dto.ApplicationDetailMobile{}, err
	}

	return s.GetApplicationDetail(ctx, createdApp.ID)
}

func (s *mobileService) GetApplicationList(ctx context.Context, userID int) ([]dto.ApplicationListMobile, error) {
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
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
			ID:        n.ID,
			Message:   n.Message,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt.Format("2006-01-02 15:04:05"),
			ReadAt:    &readAt,
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

func (s *mobileService) MarkNotificationsAsRead(ctx context.Context, umkmID int, notificationIDs []int) error {
	return s.notificationRepo.MarkAsRead(ctx, notificationIDs, umkmID)
}

func (s *mobileService) MarkAllNotificationsAsRead(ctx context.Context, umkmID int) error {
	return s.notificationRepo.MarkAllAsRead(ctx, umkmID)
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
	umkm, err := s.mobileRepo.GetUMKMProfileByUserID(ctx, userID)
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

func (s *mobileService) extractRequiredDocuments(requirements []model.ProgramRequirement) map[string]bool {
	requiredDocs := make(map[string]bool)
	for _, req := range requirements {
		reqLower := strings.ToLower(req.Name)
		if strings.Contains(reqLower, "nib") {
			requiredDocs["nib"] = true
		}
		if strings.Contains(reqLower, "npwp") {
			requiredDocs["npwp"] = true
		}
		if strings.Contains(reqLower, "catatan pendapatan") || strings.Contains(reqLower, "revenue") {
			requiredDocs["revenue_record"] = true
		}
		if strings.Contains(reqLower, "surat izin") || strings.Contains(reqLower, "business permit") {
			requiredDocs["business_permit"] = true
		}
		if strings.Contains(reqLower, "rekening") || strings.Contains(reqLower, "bank") {
			requiredDocs["rekening"] = true
		}
		if strings.Contains(reqLower, "proposal") {
			requiredDocs["proposal"] = true
		}
		if strings.Contains(reqLower, "portfolio") {
			requiredDocs["portfolio"] = true
		}
	}
	return requiredDocs
}

func (s *mobileService) validateDocuments(umkm model.UMKM, requiredDocs map[string]bool, providedDocs map[string]string, mandatoryDocs []string) error {
	// Check mandatory documents
	for _, docType := range mandatoryDocs {
		if _, exists := providedDocs[docType]; !exists {
			// Check if document already uploaded in profile
			switch docType {
			case "nib":
				if umkm.NIB == "" {
					return fmt.Errorf("NIB document is required")
				}
			case "npwp":
				if umkm.NPWP == "" {
					return fmt.Errorf("NPWP document is required")
				}
			case "revenue_record":
				if umkm.RevenueRecord == "" {
					return fmt.Errorf("Revenue Record document is required")
				}
			case "business_permit":
				if umkm.BusinessPermit == "" {
					return fmt.Errorf("Business Permit document is required")
				}
			case "ktp", "rekening", "proposal", "portfolio":
				return fmt.Errorf("%s document is required", docType)
			}
		}
	}
	return nil
}

func (s *mobileService) processAndSaveDocuments(ctx context.Context, applicationID int, umkm model.UMKM, providedDocs map[string]string, requiredDocs map[string]bool) error {
	var appDocuments []model.ApplicationDocument

	// Add documents from request
	for docType, docData := range providedDocs {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: docData,
			BucketName: "umkmgo-applications",
			Prefix:     fmt.Sprintf("app_%d_%s_", applicationID, docType),
		})
		if err != nil {
			return fmt.Errorf("failed to upload %s document", docType)
		}

		appDocuments = append(appDocuments, model.ApplicationDocument{
			ApplicationID: applicationID,
			Type:          docType,
			File:          res.URL,
		})
	}

	// Add documents from profile if required and not in request
	for docType := range requiredDocs {
		if _, exists := providedDocs[docType]; !exists {
			var fileURL string
			switch docType {
			case "nib":
				fileURL = umkm.NIB
			case "npwp":
				fileURL = umkm.NPWP
			case "revenue_record":
				fileURL = umkm.RevenueRecord
			case "business_permit":
				fileURL = umkm.BusinessPermit
			}

			if fileURL != "" {
				appDocuments = append(appDocuments, model.ApplicationDocument{
					ApplicationID: applicationID,
					Type:          docType,
					File:          fileURL,
				})
			}
		}
	}

	return s.mobileRepo.CreateApplicationDocuments(ctx, appDocuments)
}

func (s *mobileService) createApplicationHistory(ctx context.Context, applicationID, userID int, status, notes string) error {
	history := model.ApplicationHistory{
		ApplicationID: applicationID,
		Status:        status,
		Notes:         notes,
		ActionedBy:    userID,
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
