package repositories

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"gorm.io/gorm"
)

type FetchRepo struct {
	db *gorm.DB
}

func NewFetchInterface(db *gorm.DB) interfaces.FetchInterface {
	return &FetchRepo{
		db: db,
	}
}

func (r *FetchRepo) GetUserByUuid(uuid string) (*models.User, error) {
	var data models.User
	if err := r.db.Preload("Role").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetHealthCenters() (*[]models.HealthCenter, error) {
	var data []models.HealthCenter
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetHealthCenterByUuid(uuid string) (*models.HealthCenter, error) {
	var data models.HealthCenter
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCases(health_center_uuid string) (*[]models.Case, error) {
	var data []models.Case
	if health_center_uuid == "" {
		if err := r.db.Preload("HealthCenter").Find(&data).Error; err != nil {
			return nil, err
		}
	} else {
		var healthCenter models.HealthCenter
		if err := r.db.First(&healthCenter, "uuid = ?", health_center_uuid).Error; err != nil {
			return nil, err
		}

		if err := r.db.Preload("HealthCenter").Find(&data, "health_center_id = ?", healthCenter.ID).Error; err != nil {
			return nil, err
		}

	}
	return &data, nil
}

func (r *FetchRepo) GetCaseByUuid(uuid string) (*models.Case, error) {
	var data models.Case
	if err := r.db.Preload("HealthCenter").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCaseYears() (*[]uint, error) {
	var years []uint
	if err := r.db.Model(&models.Case{}).Select("DISTINCT year").Pluck("year", &years).Error; err != nil {
		return nil, err
	}
	return &years, nil
}

func (r *FetchRepo) GetResultByYear(year string) (*[]models.Result, error) {
	var data []models.Result
	if err := r.db.Preload("Case.HealthCenter").Joins("Case").Find(&data, "Case.year = ?", year).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetResultByType(year, types string) (*[]models.Result, error) {
	var data []models.Result
	if err := r.db.Preload("Case.HealthCenter").Joins("Case").Find(&data, "Case.year = ? AND type = ?", year, types).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetResultByUuid(uuid string) (*models.Result, error) {
	var resp models.Result
	if err := r.db.Preload("Case.HealthCenter.Cases").First(&resp, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *FetchRepo) GetYearCases() (*[]string, error) {
	var years []string

	// Mengambil tahun distinct dan mengurutkannya dari tahun terlama
	if err := r.db.Model(&models.Case{}).
		Distinct("year").
		Order("year ASC").
		Pluck("year", &years).Error; err != nil {
		return nil, err
	}

	return &years, nil
}

func (r *FetchRepo) GetCasesByYear(year string) (*[]models.Case, error) {
	var data []models.Case
	if err := r.db.Find(&data, "year = ?", year).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
