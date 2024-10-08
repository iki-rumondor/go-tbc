package services

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/request"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
)

type ManagementService struct {
	Repo interfaces.ManagementInterface
}

func NewManagementService(repo interfaces.ManagementInterface) *ManagementService {
	return &ManagementService{
		Repo: repo,
	}
}

func (s *ManagementService) CreateHealthCenter(req *request.HealthCenter) error {
	model := models.HealthCenter{
		Name: req.Name,
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}
