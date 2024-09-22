package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type FetchInterface interface {
	GetUserByUuid(uuid string) (*models.User, error)

	GetHealthCenters() (*[]models.HealthCenter, error)
	GetHealthCenterByUuid(uuid string) (*models.HealthCenter, error)
}
