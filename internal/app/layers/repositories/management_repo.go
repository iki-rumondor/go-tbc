package repositories

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
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
