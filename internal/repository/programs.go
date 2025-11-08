package repository

import "gorm.io/gorm"

type ProgramsRepository interface{}

type programsRepository struct {
	db *gorm.DB
}

func NewProgramsRepository(db *gorm.DB) ProgramsRepository {
	return &programsRepository{db}
}
