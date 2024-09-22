package handlers

import (
	"net/http"

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
