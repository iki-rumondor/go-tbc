package repositories

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"gorm.io/gorm"
)

type ManagementRepo struct {
	db *gorm.DB
}

func NewManagementInterface(db *gorm.DB) interfaces.ManagementInterface {
	return &ManagementRepo{
		db: db,
	}
}

func (r *ManagementRepo) GetHealthCenterByUuid(healthCenterUuid string) (*models.HealthCenter, error) {
	var data models.HealthCenter
	if err := r.db.First(&data, "uuid = ?", healthCenterUuid).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ManagementRepo) CheckCaseUnique(healthCenterID uint, year, caseUuid string) bool {
	rows := r.db.First(&models.Case{}, "uuid != ? AND health_center_id = ? AND year = ?", caseUuid, healthCenterID, year).RowsAffected
	return rows == 0
}

func (r *ManagementRepo) CreateModel(modelPointer interface{}) error {
	return r.db.Create(modelPointer).Error
}

func (r *ManagementRepo) UpdateHealthCenter(uuid string, model *models.HealthCenter) error {
	var data models.HealthCenter
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	model.ID = data.ID
	return r.db.Updates(model).Error
}

func (r *ManagementRepo) UpdateCase(uuid string, model *models.Case) error {
	var data models.Case
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	model.ID = data.ID
	return r.db.Updates(model).Error
}

func (r *ManagementRepo) DeleteCase(uuid string) error {
	var data models.Case
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Where("1=1").Delete(&models.Result{}).Error; err != nil {
			return err
		}
		if err := r.db.Delete(&data).Error; err != nil {
			return err
		}
		return nil
	})

}

func (r *ManagementRepo) DeleteHealthCenter(uuid string) error {
	var data models.HealthCenter
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Where("1=1").Delete(&models.Result{}).Error; err != nil {
			return err
		}
		if err := r.db.Delete(&data).Error; err != nil {
			return err
		}
		return nil
	})

}
