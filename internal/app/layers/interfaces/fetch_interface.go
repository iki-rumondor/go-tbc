package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type FetchInterface interface {
	GetHealthCenters() (*[]models.HealthCenter, error)
	GetHealthCenterByUuid(uuid string) (*models.HealthCenter, error)
}
