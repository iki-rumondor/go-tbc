package services

import (
	"fmt"

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
			Longitude: fmt.Sprintf("%f", item.Longitude),
			Latitude:  fmt.Sprintf("%f", item.Latitude),
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
		Longitude: fmt.Sprintf("%f", item.Longitude),
		Latitude:  fmt.Sprintf("%f", item.Latitude),
	}

	return &resp, nil
}
