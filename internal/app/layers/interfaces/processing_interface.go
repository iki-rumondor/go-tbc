package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type ProcessingInterface interface {
	GetCasesByYear(year string) (*[]models.Case, error)
	GenerateResult(year string, model *[]models.Result) error
}
