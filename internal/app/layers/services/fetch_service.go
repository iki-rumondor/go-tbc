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
			ImageName: item.ImageName,
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

func (s *FetchService) GetCases(health_center string) (*[]response.Case, error) {
	data, err := s.Repo.GetCases(health_center)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}
	var resp []response.Case
	for _, item := range *data {
		resp = append(resp, response.Case{
			Uuid:        item.Uuid,
			Year:        item.Year,
			ChildCount:  item.ChildCount,
			AdultCount:  item.AdultCount,
			MaleCount:   item.MaleCount,
			FemaleCount: item.FemaleCount,
			Total:       item.FemaleCount + item.MaleCount,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			HealthCenter: &response.HealthCenter{
				Uuid: item.HealthCenter.Uuid,
				Name: item.HealthCenter.Name,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetCaseByUuid(uuid string) (*response.Case, error) {
	item, err := s.Repo.GetCaseByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Case{
		Uuid:        item.Uuid,
		Year:        item.Year,
		ChildCount:  item.ChildCount,
		AdultCount:  item.AdultCount,
		MaleCount:   item.MaleCount,
		FemaleCount: item.FemaleCount,
		Total:       item.FemaleCount + item.MaleCount,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		HealthCenter: &response.HealthCenter{
			Uuid: item.HealthCenter.Uuid,
			Name: item.HealthCenter.Name,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetCaseYears() (*[]uint, error) {
	resp, err := s.Repo.GetCaseYears()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return resp, nil
}

func (s *FetchService) GetResultByYear(year string) (*map[string][]response.Result, error) {
	data, err := s.Repo.GetResultByYear(year)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = make(map[string][]response.Result)

	for _, item := range *data {
		var total = item.Case.AdultCount + item.Case.ChildCount
		switch item.Type {
		case "adult":
			total = item.Case.AdultCount
		case "child":
			total = item.Case.ChildCount
		case "male":
			total = item.Case.MaleCount
		case "female":
			total = item.Case.FemaleCount
		}

		resp[item.Type] = append(resp[item.Type], response.Result{
			Uuid:      item.Uuid,
			Cluster:   item.Cluster,
			Total:     total,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			Case: &response.Case{
				Year: item.Case.Year,
				HealthCenter: &response.HealthCenter{
					Name: item.Case.HealthCenter.Name,
				},
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetResultByType(year, category string) (*[]response.Result, error) {
	data, err := s.Repo.GetResultByType(year, category)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Result

	for _, item := range *data {
		resp = append(resp, response.Result{
			Uuid:    item.Uuid,
			Cluster: item.Cluster,
			Case: &response.Case{
				HealthCenter: &response.HealthCenter{
					Name:      item.Case.HealthCenter.Name,
					Latitude:  item.Case.HealthCenter.Latitude,
					Longitude: item.Case.HealthCenter.Longitude,
				},
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetResultByUuid(uuid string) (*response.Result, error) {
	item, err := s.Repo.GetResultByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}
	var cases []response.Case

	for _, i := range *item.Case.HealthCenter.Cases {
		cases = append(cases, response.Case{
			Year:        i.Year,
			ChildCount:  i.ChildCount,
			AdultCount:  i.AdultCount,
			MaleCount:   i.MaleCount,
			FemaleCount: i.FemaleCount,
			Total:       i.MaleCount + i.FemaleCount,
		})
	}

	var resp = response.Result{
		Uuid:    item.Uuid,
		Cluster: item.Cluster,
		Type:    item.Type,
		Case: &response.Case{
			HealthCenter: &response.HealthCenter{
				Name:      item.Case.HealthCenter.Name,
				ImageName: item.Case.HealthCenter.ImageName,
				Cases:     &cases,
			},
		},
	}

	return &resp, nil
}

func (s *FetchService) GetDashboardInformation() (*response.DashboardInformation, error) {
	years, err := s.Repo.GetYearCases()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var totalCases []int64

	for _, item := range *years {
		results, err := s.Repo.GetCasesByYear(item)
		if err != nil {
			return nil, response.SERVICE_INTERR
		}

		var total int64

		for _, result := range *results {
			total += result.AdultCount + result.ChildCount
		}

		totalCases = append(totalCases, total)
	}

	var resp = response.DashboardInformation{
		YearCases:  *years,
		TotalCases: totalCases,
	}

	return &resp, nil
}
