package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type ManagementInterface interface {
	CheckCaseUnique(healthCenterID uint, year, caseUuid string) bool
	GetHealthCenterByUuid(healthCenterUuid string) (*models.HealthCenter, error)

	CreateModel(pointerModel interface{}) error
	UpdateHealthCenter(uuid string, model *models.HealthCenter) error
	DeleteHealthCenter(uuid string) error

	UpdateCase(uuid string, model *models.Case) error
	DeleteCase(uuid string) error
}
