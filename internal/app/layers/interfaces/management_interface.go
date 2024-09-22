package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type ManagementInterface interface {
	CreateModel(pointerModel interface{}) error
	UpdateHealthCenter(uuid string, model *models.HealthCenter) error
}
