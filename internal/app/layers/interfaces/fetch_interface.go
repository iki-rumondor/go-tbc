package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type FetchInterface interface {
	GetUserByUuid(uuid string) (*models.User, error)

	GetHealthCenters() (*[]models.HealthCenter, error)
	GetHealthCenterByUuid(uuid string) (*models.HealthCenter, error)

	GetCases() (*[]models.Case, error)
	GetCaseByUuid(uuid string) (*models.Case, error)

	GetCaseYears() (*[]uint, error)
	GetResultByYear(year string) (*[]models.Result, error)
	GetResultByType(year, types string) (*[]models.Result, error)
	GetResultByUuid(uuid string) (*models.Result, error)
}
