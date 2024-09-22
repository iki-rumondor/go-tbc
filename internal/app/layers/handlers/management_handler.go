package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-tbc/internal/app/layers/services"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/request"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
	"github.com/iki-rumondor/go-tbc/internal/utils"
)

type ManagementHandler struct {
	Service *services.ManagementService
}

func NewManagementHandler(service *services.ManagementService) *ManagementHandler {
	return &ManagementHandler{
		Service: service,
	}
}

func (h *ManagementHandler) CreateHealthCenter(c *gin.Context) {
	var body request.HealthCenter
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateHealthCenter(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Puskesmas Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateHealthCenter(c *gin.Context) {
	var body request.HealthCenter
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateHealthCenter(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Puskesmas Berhasil Diperbarui"))
}
