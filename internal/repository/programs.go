package repository

import (
	"context"
	"errors"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type ProgramsRepository interface {
	GetAllPrograms(ctx context.Context) ([]model.Programs, error)
	GetProgramByID(ctx context.Context, id int) (model.Programs, error)
	CreateProgram(ctx context.Context, program model.Programs) (model.Programs, error)
	UpdateProgram(ctx context.Context, program model.Programs) (model.Programs, error)
	DeleteProgram(ctx context.Context, program model.Programs) (model.Programs, error)

	// Benefits and Requirements
	CreateProgramBenefits(ctx context.Context, benefits []model.ProgramBenefits) error
	CreateProgramRequirements(ctx context.Context, requirements []model.ProgramRequirements) error
	GetProgramBenefits(ctx context.Context, programID int) ([]model.ProgramBenefits, error)
	GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirements, error)
	DeleteProgramBenefits(ctx context.Context, programID int) error
	DeleteProgramRequirements(ctx context.Context, programID int) error
}

type programsRepository struct {
	db *gorm.DB
}

func NewProgramsRepository(db *gorm.DB) ProgramsRepository {
	return &programsRepository{db}
}

func (repo *programsRepository) GetAllPrograms(ctx context.Context) ([]model.Programs, error) {
	var programs []model.Programs
	err := repo.db.WithContext(ctx).Preload("Users").Where("deleted_at IS NULL").Find(&programs).Error
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (repo *programsRepository) GetProgramByID(ctx context.Context, id int) (model.Programs, error) {
	var program model.Programs
	err := repo.db.WithContext(ctx).Preload("Users").Where("id = ? AND deleted_at IS NULL", id).First(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("program not found")
	}
	return program, nil
}

func (repo *programsRepository) CreateProgram(ctx context.Context, program model.Programs) (model.Programs, error) {
	err := repo.db.WithContext(ctx).Create(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to create program")
	}
	return program, nil
}

func (repo *programsRepository) UpdateProgram(ctx context.Context, program model.Programs) (model.Programs, error) {
	err := repo.db.WithContext(ctx).Save(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to update program")
	}
	return program, nil
}

func (repo *programsRepository) DeleteProgram(ctx context.Context, program model.Programs) (model.Programs, error) {
	err := repo.db.WithContext(ctx).Delete(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to delete program")
	}
	return program, nil
}

func (repo *programsRepository) CreateProgramBenefits(ctx context.Context, benefits []model.ProgramBenefits) error {
	if len(benefits) == 0 {
		return nil
	}
	err := repo.db.WithContext(ctx).Create(&benefits).Error
	if err != nil {
		return errors.New("failed to create program benefits")
	}
	return nil
}

func (repo *programsRepository) CreateProgramRequirements(ctx context.Context, requirements []model.ProgramRequirements) error {
	if len(requirements) == 0 {
		return nil
	}
	err := repo.db.WithContext(ctx).Create(&requirements).Error
	if err != nil {
		return errors.New("failed to create program requirements")
	}
	return nil
}

func (repo *programsRepository) GetProgramBenefits(ctx context.Context, programID int) ([]model.ProgramBenefits, error) {
	var benefits []model.ProgramBenefits
	err := repo.db.WithContext(ctx).Where("program_id = ? AND deleted_at IS NULL", programID).Find(&benefits).Error
	if err != nil {
		return nil, err
	}
	return benefits, nil
}

func (repo *programsRepository) GetProgramRequirements(ctx context.Context, programID int) ([]model.ProgramRequirements, error) {
	var requirements []model.ProgramRequirements
	err := repo.db.WithContext(ctx).Where("program_id = ? AND deleted_at IS NULL", programID).Find(&requirements).Error
	if err != nil {
		return nil, err
	}
	return requirements, nil
}

func (repo *programsRepository) DeleteProgramBenefits(ctx context.Context, programID int) error {
	err := repo.db.WithContext(ctx).Where("program_id = ?", programID).Unscoped().Delete(&model.ProgramBenefits{}).Error
	if err != nil {
		return errors.New("failed to delete program benefits")
	}
	return nil
}

func (repo *programsRepository) DeleteProgramRequirements(ctx context.Context, programID int) error {
	err := repo.db.WithContext(ctx).Where("program_id = ?", programID).Unscoped().Delete(&model.ProgramRequirements{}).Error
	if err != nil {
		return errors.New("failed to delete program requirements")
	}
	return nil
}
