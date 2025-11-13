package repository

import (
	"context"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetUMKMByCardType(ctx context.Context) ([]map[string]interface{}, error)
	GetApplicationStatusSummary(ctx context.Context) (map[string]int64, error)
	GetApplicationStatusDetail(ctx context.Context) (map[string]int64, error)
	GetApplicationByType(ctx context.Context) (map[string]int64, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (repo *dashboardRepository) GetUMKMByCardType(ctx context.Context) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	err := repo.db.WithContext(ctx).
		Raw(`
			SELECT 
				CASE 
					WHEN kartu_type = 'produktif' THEN 'Kartu Produktif'
					WHEN kartu_type = 'afirmatif' THEN 'Kartu Afirmatif'
					ELSE 'Unknown'
				END as name,
				COUNT(*) as count
			FROM umkms
			WHERE deleted_at IS NULL AND kartu_type IS NOT NULL
			GROUP BY kartu_type
		`).
		Scan(&results).Error

	return results, err
}

func (repo *dashboardRepository) GetApplicationStatusSummary(ctx context.Context) (map[string]int64, error) {
	result := make(map[string]int64)

	// Total applications
	var total int64
	repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM applications WHERE deleted_at IS NULL").
		Scan(&total)
	result["total_applications"] = total

	// In process (screening + revised + final)
	var inProcess int64
	repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM applications WHERE deleted_at IS NULL AND status IN ('screening', 'revised', 'final')").
		Scan(&inProcess)
	result["in_process"] = inProcess

	// Approved
	var approved int64
	repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM applications WHERE deleted_at IS NULL AND status = 'approved'").
		Scan(&approved)
	result["approved"] = approved

	// Rejected
	var rejected int64
	repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM applications WHERE deleted_at IS NULL AND status = 'rejected'").
		Scan(&rejected)
	result["rejected"] = rejected

	return result, nil
}

func (repo *dashboardRepository) GetApplicationStatusDetail(ctx context.Context) (map[string]int64, error) {
	result := make(map[string]int64)

	var statuses []struct {
		Status string
		Count  int64
	}

	err := repo.db.WithContext(ctx).
		Raw(`
			SELECT status, COUNT(*) as count
			FROM applications
			WHERE deleted_at IS NULL
			GROUP BY status
		`).
		Scan(&statuses).Error
	if err != nil {
		return nil, err
	}

	for _, s := range statuses {
		result[s.Status] = s.Count
	}

	return result, nil
}

func (repo *dashboardRepository) GetApplicationByType(ctx context.Context) (map[string]int64, error) {
	result := make(map[string]int64)

	var types []struct {
		Type  string
		Count int64
	}

	err := repo.db.WithContext(ctx).
		Raw(`
			SELECT type, COUNT(*) as count
			FROM applications
			WHERE deleted_at IS NULL
			GROUP BY type
		`).
		Scan(&types).Error
	if err != nil {
		return nil, err
	}

	for _, t := range types {
		result[t.Type] = t.Count
	}

	return result, nil
}
