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
