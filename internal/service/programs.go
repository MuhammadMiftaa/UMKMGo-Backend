package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/config/storage"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/types/model"
	"sapaUMKM-backend/internal/utils"
)

type ProgramsService interface {
	GetAllPrograms(ctx context.Context) ([]dto.Programs, error)
	GetProgramByID(ctx context.Context, id int) (dto.Programs, error)
	CreateProgram(ctx context.Context, program dto.Programs) (dto.Programs, error)
	UpdateProgram(ctx context.Context, id int, program dto.Programs) (dto.Programs, error)
	DeleteProgram(ctx context.Context, id int) (dto.Programs, error)
	ActivateProgram(ctx context.Context, id int) (dto.Programs, error)
	DeactivateProgram(ctx context.Context, id int) (dto.Programs, error)
}

type programsService struct {
	programRepository repository.ProgramsRepository
	userRepository    repository.UsersRepository
	redisRepository   redis.RedisRepository
	minio             *storage.MinIOManager
}

func NewProgramsService(programRepo repository.ProgramsRepository, userRepo repository.UsersRepository, redisRepo redis.RedisRepository, minio *storage.MinIOManager) ProgramsService {
	return &programsService{
		programRepository: programRepo,
		userRepository:    userRepo,
		redisRepository:   redisRepo,
		minio:             minio,
	}
}

func (s *programsService) GetAllPrograms(ctx context.Context) ([]dto.Programs, error) {
	programs, err := s.programRepository.GetAllPrograms(ctx)
	if err != nil {
		return nil, err
	}

	var programsDTO []dto.Programs
	for _, program := range programs {
		// Get benefits
		benefits, _ := s.programRepository.GetProgramBenefits(ctx, program.ID)
		var benefitNames []string
		for _, b := range benefits {
			benefitNames = append(benefitNames, b.Name)
		}

		// Get requirements
		requirements, _ := s.programRepository.GetProgramRequirements(ctx, program.ID)
		var requirementNames []string
		for _, r := range requirements {
			requirementNames = append(requirementNames, r.Name)
		}

		programDTO := dto.Programs{
			ID:                  program.ID,
			Title:               program.Title,
			Description:         program.Description,
			Banner:              program.Banner,
			Provider:            program.Provider,
			ProviderLogo:        program.ProviderLogo,
			Type:                program.Type,
			TrainingType:        program.TrainingType,
			Batch:               program.Batch,
			BatchStartDate:      program.BatchStartDate,
			BatchEndDate:        program.BatchEndDate,
			Location:            program.Location,
			MinAmount:           program.MinAmount,
			MaxAmount:           program.MaxAmount,
			InterestRate:        program.InterestRate,
			MaxTenureMonths:     program.MaxTenureMonths,
			ApplicationDeadline: program.ApplicationDeadline,
			IsActive:            program.IsActive,
			CreatedBy:           program.CreatedBy,
			CreatedByName:       program.Users.Name,
			CreatedAt:           program.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:           program.UpdatedAt.Format("2006-01-02 15:04:05"),
			Benefits:            benefitNames,
			Requirements:        requirementNames,
		}
		programsDTO = append(programsDTO, programDTO)
	}

	return programsDTO, nil
}

func (s *programsService) GetProgramByID(ctx context.Context, id int) (dto.Programs, error) {
	program, err := s.programRepository.GetProgramByID(ctx, id)
	if err != nil {
		return dto.Programs{}, err
	}

	// Get benefits
	benefits, _ := s.programRepository.GetProgramBenefits(ctx, program.ID)
	var benefitNames []string
	for _, b := range benefits {
		benefitNames = append(benefitNames, b.Name)
	}

	// Get requirements
	requirements, _ := s.programRepository.GetProgramRequirements(ctx, program.ID)
	var requirementNames []string
	for _, r := range requirements {
		requirementNames = append(requirementNames, r.Name)
	}

	return dto.Programs{
		ID:                  program.ID,
		Title:               program.Title,
		Description:         program.Description,
		Banner:              program.Banner,
		Provider:            program.Provider,
		ProviderLogo:        program.ProviderLogo,
		Type:                program.Type,
		TrainingType:        program.TrainingType,
		Batch:               program.Batch,
		BatchStartDate:      program.BatchStartDate,
		BatchEndDate:        program.BatchEndDate,
		Location:            program.Location,
		MinAmount:           program.MinAmount,
		MaxAmount:           program.MaxAmount,
		InterestRate:        program.InterestRate,
		MaxTenureMonths:     program.MaxTenureMonths,
		ApplicationDeadline: program.ApplicationDeadline,
		IsActive:            program.IsActive,
		CreatedBy:           program.CreatedBy,
		CreatedByName:       program.Users.Name,
		CreatedAt:           program.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:           program.UpdatedAt.Format("2006-01-02 15:04:05"),
		Benefits:            benefitNames,
		Requirements:        requirementNames,
	}, nil
}

func (s *programsService) CreateProgram(ctx context.Context, program dto.Programs) (dto.Programs, error) {
	// Validation
	if program.Title == "" || program.Type == "" || program.ApplicationDeadline == "" {
		return dto.Programs{}, errors.New("title, type, and application deadline are required")
	}

	// Validate type
	if program.Type != "training" && program.Type != "certification" && program.Type != "funding" {
		return dto.Programs{}, errors.New("type must be training, certification, or funding")
	}

	// Validate training type if type is training or certification
	if (program.Type == "training" || program.Type == "certification") && program.TrainingType != nil {
		if *program.TrainingType != "online" && *program.TrainingType != "offline" && *program.TrainingType != "hybrid" {
			return dto.Programs{}, errors.New("training type must be online, offline, or hybrid")
		}
	}

	// Check if user exists
	if program.CreatedBy > 0 {
		_, err := s.userRepository.GetUserByID(ctx, program.CreatedBy)
		if err != nil {
			return dto.Programs{}, errors.New("creator user not found")
		}
	}

	// If banner is provided, upload to MinIO
	if program.Banner != "" {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: program.Banner,
			BucketName: storage.ProgramBucket,
			Prefix:     utils.GenerateFileName(program.Title, "banner_"),
		})
		if err != nil {
			return dto.Programs{}, err
		}
		program.Banner = res.URL
	}

	// If provider logo is provided, upload to MinIO
	if program.ProviderLogo != "" {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: program.ProviderLogo,
			BucketName: storage.ProgramBucket,
			Prefix:     utils.GenerateFileName(program.Title, "provider_logos_"),
		})
		if err != nil {
			return dto.Programs{}, err
		}
		program.ProviderLogo = res.URL
	}

	// Create program
	newProgram := model.Programs{
		Title:               program.Title,
		Description:         program.Description,
		Banner:              program.Banner,
		Provider:            program.Provider,
		ProviderLogo:        program.ProviderLogo,
		Type:                program.Type,
		TrainingType:        program.TrainingType,
		Batch:               program.Batch,
		BatchStartDate:      program.BatchStartDate,
		BatchEndDate:        program.BatchEndDate,
		Location:            program.Location,
		MinAmount:           program.MinAmount,
		MaxAmount:           program.MaxAmount,
		InterestRate:        program.InterestRate,
		MaxTenureMonths:     program.MaxTenureMonths,
		ApplicationDeadline: program.ApplicationDeadline,
		IsActive:            true,
		CreatedBy:           program.CreatedBy,
	}

	createdProgram, err := s.programRepository.CreateProgram(ctx, newProgram)
	if err != nil {
		return dto.Programs{}, err
	}

	// Create benefits
	if len(program.Benefits) > 0 {
		var benefits []model.ProgramBenefits
		for _, b := range program.Benefits {
			benefits = append(benefits, model.ProgramBenefits{
				ProgramID: createdProgram.ID,
				Name:      b,
			})
		}
		if err := s.programRepository.CreateProgramBenefits(ctx, benefits); err != nil {
			return dto.Programs{}, err
		}
	}

	// Create requirements
	if len(program.Requirements) > 0 {
		var requirements []model.ProgramRequirements
		for _, r := range program.Requirements {
			requirements = append(requirements, model.ProgramRequirements{
				ProgramID: createdProgram.ID,
				Name:      r,
			})
		}
		if err := s.programRepository.CreateProgramRequirements(ctx, requirements); err != nil {
			return dto.Programs{}, err
		}
	}

	return dto.Programs{
		ID:                  createdProgram.ID,
		Title:               createdProgram.Title,
		Description:         createdProgram.Description,
		Type:                createdProgram.Type,
		ApplicationDeadline: createdProgram.ApplicationDeadline,
		IsActive:            createdProgram.IsActive,
		Benefits:            program.Benefits,
		Requirements:        program.Requirements,
	}, nil
}

func (s *programsService) UpdateProgram(ctx context.Context, id int, program dto.Programs) (dto.Programs, error) {
	// Get existing program
	existingProgram, err := s.programRepository.GetProgramByID(ctx, id)
	if err != nil {
		return dto.Programs{}, err
	}

	// Validation
	if program.Title == "" || program.Type == "" || program.ApplicationDeadline == "" {
		return dto.Programs{}, errors.New("title, type, and application deadline are required")
	}

	// Validate type
	if program.Type != "training" && program.Type != "certification" && program.Type != "funding" {
		return dto.Programs{}, errors.New("type must be training, certification, or funding")
	}

	// If banner is provided, upload to MinIO
	if !(strings.HasPrefix(program.Banner, "http") || strings.HasPrefix(program.Banner, "https")) {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: program.Banner,
			BucketName: storage.ProgramBucket,
			Prefix:     utils.GenerateFileName(program.Title, "banner_"),
		})
		if err != nil {
			return dto.Programs{}, err
		}

		if existingProgram.Banner != "" {
			objectName := storage.ExtractObjectNameFromURL(existingProgram.Banner)
			if objectName != "" {
				if deleteErr := s.minio.DeleteFile(ctx, storage.ProgramBucket, objectName); deleteErr != nil {
					return dto.Programs{}, fmt.Errorf("failed to delete old banner: %w", deleteErr)
				}
			}
		}

		program.Banner = res.URL
	}

	// If provider logo is provided, upload to MinIO
	if !(strings.HasPrefix(program.ProviderLogo, "http") || strings.HasPrefix(program.ProviderLogo, "https")) {
		res, err := s.minio.UploadFile(ctx, storage.UploadRequest{
			Base64Data: program.ProviderLogo,
			BucketName: storage.ProgramBucket,
			Prefix:     utils.GenerateFileName(program.Title, "provider_logos_"),
		})
		if err != nil {
			return dto.Programs{}, err
		}

		if existingProgram.ProviderLogo != "" {
			objectName := storage.ExtractObjectNameFromURL(existingProgram.ProviderLogo)
			if objectName != "" {
				if deleteErr := s.minio.DeleteFile(ctx, storage.ProgramBucket, objectName); deleteErr != nil {
					return dto.Programs{}, fmt.Errorf("failed to delete old provider logo: %w", deleteErr)
				}
			}
		}

		program.ProviderLogo = res.URL
	}

	// Update fields
	existingProgram.Title = program.Title
	existingProgram.Description = program.Description
	existingProgram.Banner = program.Banner
	existingProgram.Provider = program.Provider
	existingProgram.ProviderLogo = program.ProviderLogo
	existingProgram.Type = program.Type
	existingProgram.TrainingType = program.TrainingType
	existingProgram.Batch = program.Batch
	existingProgram.BatchStartDate = program.BatchStartDate
	existingProgram.BatchEndDate = program.BatchEndDate
	existingProgram.Location = program.Location
	existingProgram.MinAmount = program.MinAmount
	existingProgram.MaxAmount = program.MaxAmount
	existingProgram.InterestRate = program.InterestRate
	existingProgram.MaxTenureMonths = program.MaxTenureMonths
	existingProgram.ApplicationDeadline = program.ApplicationDeadline

	updatedProgram, err := s.programRepository.UpdateProgram(ctx, existingProgram)
	if err != nil {
		return dto.Programs{}, err
	}

	// Update benefits
	if len(program.Benefits) > 0 {
		// Delete old benefits
		_ = s.programRepository.DeleteProgramBenefits(ctx, id)

		// Create new benefits
		var benefits []model.ProgramBenefits
		for _, b := range program.Benefits {
			benefits = append(benefits, model.ProgramBenefits{
				ProgramID: id,
				Name:      b,
			})
		}
		if err := s.programRepository.CreateProgramBenefits(ctx, benefits); err != nil {
			return dto.Programs{}, err
		}
	}

	// Update requirements
	if len(program.Requirements) > 0 {
		// Delete old requirements
		_ = s.programRepository.DeleteProgramRequirements(ctx, id)

		// Create new requirements
		var requirements []model.ProgramRequirements
		for _, r := range program.Requirements {
			requirements = append(requirements, model.ProgramRequirements{
				ProgramID: id,
				Name:      r,
			})
		}
		if err := s.programRepository.CreateProgramRequirements(ctx, requirements); err != nil {
			return dto.Programs{}, err
		}
	}

	return dto.Programs{
		ID:                  updatedProgram.ID,
		Title:               updatedProgram.Title,
		Description:         updatedProgram.Description,
		Type:                updatedProgram.Type,
		ApplicationDeadline: updatedProgram.ApplicationDeadline,
		IsActive:            updatedProgram.IsActive,
		Benefits:            program.Benefits,
		Requirements:        program.Requirements,
	}, nil
}

func (s *programsService) DeleteProgram(ctx context.Context, id int) (dto.Programs, error) {
	program, err := s.programRepository.GetProgramByID(ctx, id)
	if err != nil {
		return dto.Programs{}, err
	}

	deletedProgram, err := s.programRepository.DeleteProgram(ctx, program)
	if err != nil {
		return dto.Programs{}, err
	}

	return dto.Programs{
		ID:    deletedProgram.ID,
		Title: deletedProgram.Title,
		Type:  deletedProgram.Type,
	}, nil
}

func (s *programsService) ActivateProgram(ctx context.Context, id int) (dto.Programs, error) {
	program, err := s.programRepository.GetProgramByID(ctx, id)
	if err != nil {
		return dto.Programs{}, err
	}

	program.IsActive = true
	updatedProgram, err := s.programRepository.UpdateProgram(ctx, program)
	if err != nil {
		return dto.Programs{}, err
	}

	return dto.Programs{
		ID:       updatedProgram.ID,
		Title:    updatedProgram.Title,
		Type:     updatedProgram.Type,
		IsActive: updatedProgram.IsActive,
	}, nil
}

func (s *programsService) DeactivateProgram(ctx context.Context, id int) (dto.Programs, error) {
	program, err := s.programRepository.GetProgramByID(ctx, id)
	if err != nil {
		return dto.Programs{}, err
	}

	program.IsActive = false
	updatedProgram, err := s.programRepository.UpdateProgram(ctx, program)
	if err != nil {
		return dto.Programs{}, err
	}

	return dto.Programs{
		ID:       updatedProgram.ID,
		Title:    updatedProgram.Title,
		Type:     updatedProgram.Type,
		IsActive: updatedProgram.IsActive,
	}, nil
}
