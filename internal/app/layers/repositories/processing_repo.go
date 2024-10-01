package repositories

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"gorm.io/gorm"
)

type ProcessingRepo struct {
	db *gorm.DB
}

func NewProcessingInterface(db *gorm.DB) interfaces.ProcessingInterface {
	return &ProcessingRepo{
		db: db,
	}
}

func (r *ProcessingRepo) CreateResult(model *models.Result) error {
	return r.db.Create(model).Error
}

func (r *ProcessingRepo) GetCasesByYear(year string) (*[]models.Case, error) {
	var result []models.Case
	if err := r.db.Find(&result, "year = ?", year).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProcessingRepo) GenerateResult(year string, model *[]models.Result) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		var caseIDs []uint
		if err := tx.Model(&models.Case{}).Where("year = ?", year).Pluck("id", &caseIDs).Error; err != nil {
			return err
		}

		if err := tx.Where("case_id IN (?)", caseIDs).Delete(&models.Result{}).Error; err != nil {
			return err
		}

		if err := tx.Create(model).Error; err != nil {
			return err
		}

		return nil
	})
}
