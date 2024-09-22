package services

import (
	"fmt"

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
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateHealthCenter(uuid string, req *request.HealthCenter) error {
	model := models.HealthCenter{
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	if err := s.Repo.UpdateHealthCenter(uuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) CreateCase(req *request.Case) error {
	if req.AdultCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus dewasa tidak valid")
	}

	if req.ChildCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus anak-anak tidak valid")
	}

	if req.MaleCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus laki-laki tidak valid")
	}

	if req.FemaleCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus perempuan tidak valid")
	}

	totalByAge := req.ChildCount + req.AdultCount
	totalByGender := req.MaleCount + req.FemaleCount
	if totalByAge != totalByGender {
		return response.BADREQ_ERR("Total kasus berdasarkan umur dan berdasarkan jenis kelamin tidak sama")
	}

	healthCenter, err := s.Repo.GetHealthCenterByUuid(req.HealthCenterUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	if unique := s.Repo.CheckCaseUnique(healthCenter.ID, req.Year, ""); !unique {
		message := fmt.Sprintf("Kasus untuk puskesmas %s pada tahun %s sudah ada", healthCenter.Name, req.Year)
		return response.BADREQ_ERR(message)
	}

	model := models.Case{
		Year:           req.Year,
		ChildCount:     req.ChildCount,
		AdultCount:     req.AdultCount,
		MaleCount:      req.MaleCount,
		FemaleCount:    req.FemaleCount,
		HealthCenterID: healthCenter.ID,
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateCase(uuid string, req *request.Case) error {
	if req.AdultCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus dewasa tidak valid")
	}

	if req.ChildCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus anak-anak tidak valid")
	}

	if req.MaleCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus laki-laki tidak valid")
	}

	if req.FemaleCount < 0 {
		return response.BADREQ_ERR("Jumlah kasus perempuan tidak valid")
	}

	totalByAge := req.ChildCount + req.AdultCount
	totalByGender := req.MaleCount + req.FemaleCount
	if totalByAge != totalByGender {
		return response.BADREQ_ERR("Total kasus berdasarkan umur dan berdasarkan jenis kelamin tidak sama")
	}

	healthCenter, err := s.Repo.GetHealthCenterByUuid(req.HealthCenterUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	if unique := s.Repo.CheckCaseUnique(healthCenter.ID, req.Year, uuid); !unique {
		message := fmt.Sprintf("Kasus untuk puskesmas %s pada tahun %s sudah ada", healthCenter.Name, req.Year)
		return response.BADREQ_ERR(message)
	}

	model := models.Case{
		Year:           req.Year,
		ChildCount:     req.ChildCount,
		AdultCount:     req.AdultCount,
		MaleCount:      req.MaleCount,
		FemaleCount:    req.FemaleCount,
		HealthCenterID: healthCenter.ID,
	}

	if err := s.Repo.UpdateCase(uuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) DeleteCase(uuid string) error {
	if err := s.Repo.DeleteCase(uuid); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}
