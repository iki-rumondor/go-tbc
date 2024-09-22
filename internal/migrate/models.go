package migrate

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.Role{}},
		{Model: models.User{}},
		{Model: models.HealthCenter{}},
		{Model: models.Case{}},
	}
}
