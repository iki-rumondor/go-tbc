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

func (r *FetchRepo) GetCases() (*[]models.Case, error) {
	var data []models.Case
	if err := r.db.Preload("HealthCenter").Find(&data).Error; err != nil {
		return nil, err
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
