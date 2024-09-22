package services

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
)

type FetchService struct {
	Repo interfaces.FetchInterface
}

func NewFetchService(repo interfaces.FetchInterface) *FetchService {
	return &FetchService{
		Repo: repo,
	}
}

func (s *FetchService) GetUserByUuid(uuid string) (*response.User, error) {
	item, err := s.Repo.GetUserByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.User{
		Uuid:     item.Uuid,
		Name:     item.Name,
		Username: item.Username,
		RoleName: item.Role.Name,
	}

	return &resp, nil
}

func (s *FetchService) GetHealthCenters() (*[]response.HealthCenter, error) {
	data, err := s.Repo.GetHealthCenters()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}
	var resp []response.HealthCenter
	for _, item := range *data {
		resp = append(resp, response.HealthCenter{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Longitude: item.Longitude,
			Latitude:  item.Latitude,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *FetchService) GetHealthCenterByUuid(uuid string) (*response.HealthCenter, error) {
	item, err := s.Repo.GetHealthCenterByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.HealthCenter{
		Uuid:      item.Uuid,
		Name:      item.Name,
		Longitude: item.Longitude,
		Latitude:  item.Latitude,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	return &resp, nil
}
