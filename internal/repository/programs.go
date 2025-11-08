package repository

import (
	"errors"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type ProgramsRepository interface {
	GetAllPrograms() ([]model.Programs, error)
	GetProgramByID(id int) (model.Programs, error)
	CreateProgram(program model.Programs) (model.Programs, error)
	UpdateProgram(program model.Programs) (model.Programs, error)
	DeleteProgram(program model.Programs) (model.Programs, error)

	// Benefits and Requirements
	CreateProgramBenefits(benefits []model.ProgramBenefits) error
	CreateProgramRequirements(requirements []model.ProgramRequirements) error
	GetProgramBenefits(programID int) ([]model.ProgramBenefits, error)
	GetProgramRequirements(programID int) ([]model.ProgramRequirements, error)
	DeleteProgramBenefits(programID int) error
	DeleteProgramRequirements(programID int) error
}

type programsRepository struct {
	db *gorm.DB
}

func NewProgramsRepository(db *gorm.DB) ProgramsRepository {
	return &programsRepository{db}
}

func (repo *programsRepository) GetAllPrograms() ([]model.Programs, error) {
	var programs []model.Programs
	err := repo.db.Preload("Users").Where("deleted_at IS NULL").Find(&programs).Error
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (repo *programsRepository) GetProgramByID(id int) (model.Programs, error) {
	var program model.Programs
	err := repo.db.Preload("Users").Where("id = ? AND deleted_at IS NULL", id).First(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("program not found")
	}
	return program, nil
}

func (repo *programsRepository) CreateProgram(program model.Programs) (model.Programs, error) {
	err := repo.db.Create(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to create program")
	}
	return program, nil
}

func (repo *programsRepository) UpdateProgram(program model.Programs) (model.Programs, error) {
	err := repo.db.Save(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to update program")
	}
	return program, nil
}

func (repo *programsRepository) DeleteProgram(program model.Programs) (model.Programs, error) {
	err := repo.db.Delete(&program).Error
	if err != nil {
		return model.Programs{}, errors.New("failed to delete program")
	}
	return program, nil
}

func (repo *programsRepository) CreateProgramBenefits(benefits []model.ProgramBenefits) error {
	if len(benefits) == 0 {
		return nil
	}
	err := repo.db.Create(&benefits).Error
	if err != nil {
		return errors.New("failed to create program benefits")
	}
	return nil
}

func (repo *programsRepository) CreateProgramRequirements(requirements []model.ProgramRequirements) error {
	if len(requirements) == 0 {
		return nil
	}
	err := repo.db.Create(&requirements).Error
	if err != nil {
		return errors.New("failed to create program requirements")
	}
	return nil
}

func (repo *programsRepository) GetProgramBenefits(programID int) ([]model.ProgramBenefits, error) {
	var benefits []model.ProgramBenefits
	err := repo.db.Where("program_id = ? AND deleted_at IS NULL", programID).Find(&benefits).Error
	if err != nil {
		return nil, err
	}
	return benefits, nil
}

func (repo *programsRepository) GetProgramRequirements(programID int) ([]model.ProgramRequirements, error) {
	var requirements []model.ProgramRequirements
	err := repo.db.Where("program_id = ? AND deleted_at IS NULL", programID).Find(&requirements).Error
	if err != nil {
		return nil, err
	}
	return requirements, nil
}

func (repo *programsRepository) DeleteProgramBenefits(programID int) error {
	err := repo.db.Where("program_id = ?", programID).Unscoped().Delete(&model.ProgramBenefits{}).Error
	if err != nil {
		return errors.New("failed to delete program benefits")
	}
	return nil
}

func (repo *programsRepository) DeleteProgramRequirements(programID int) error {
	err := repo.db.Where("program_id = ?", programID).Unscoped().Delete(&model.ProgramRequirements{}).Error
	if err != nil {
		return errors.New("failed to delete program requirements")
	}
	return nil
}
