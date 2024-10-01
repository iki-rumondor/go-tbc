package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-tbc/internal/app/layers/services"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
	"github.com/iki-rumondor/go-tbc/internal/utils"
)

type FetchHandler struct {
	Service *services.FetchService
}

func NewFetchHandler(service *services.FetchService) *FetchHandler {
	return &FetchHandler{
		Service: service,
	}
}

func (h *FetchHandler) GetUserByUuid(c *gin.Context) {
	uuid := c.GetString("uuid")
	resp, err := h.Service.GetUserByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetHealthCenters(c *gin.Context) {

	resp, err := h.Service.GetHealthCenters()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetHealthCenterByUuid(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetHealthCenterByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCases(c *gin.Context) {

	resp, err := h.Service.GetCases()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCaseByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.GetCaseByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetCaseYears(c *gin.Context) {
	resp, err := h.Service.GetCaseYears()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetResultByYear(c *gin.Context) {
	year := c.Param("year")
	resp, err := h.Service.GetResultByYear(year)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetResultByType(c *gin.Context) {
	year := c.Param("year")
	category := c.Param("category")
	resp, err := h.Service.GetResultByType(year, category)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetResultByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.GetResultByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetDashboardInformation(c *gin.Context) {
	resp, err := h.Service.GetDashboardInformation()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FetchHandler) GetHealthCenterImage(c *gin.Context) {
	filename := c.Param("filename")
	folder := "internal/files/puskesmas"
	pathFile := filepath.Join(folder, filename)
	c.File(pathFile)
}
