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

type ProcessingHandler struct {
	Service *services.ProcessingService
}

func NewProcessingHandler(service *services.ProcessingService) *ProcessingHandler {
	return &ProcessingHandler{
		Service: service,
	}
}

func (h *ProcessingHandler) KmeansClustering(c *gin.Context) {
	var body request.Clustering
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.KmeansClustering(body.Year); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Proses Clustering Berhasil Dilakukan"))
}
