package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
		utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
		return
	}

	if moreThan := utils.CheckFileSize(file, 1); moreThan {
		utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
		return
	}

	puskesmasFolder := "internal/files/puskesmas"
	filename := utils.RandomFileName(file)
	pathFile := filepath.Join(puskesmasFolder, filename)

	if err := utils.SaveUploadedFile(file, pathFile); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.CreateHealthCenter(filename, &body); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Puskesmas Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateHealthCenter(c *gin.Context) {
	var body request.HealthCenter
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	var filename string

	file, _ := c.FormFile("file")
	if file != nil {
		if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
			utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
			return
		}

		if moreThan := utils.CheckFileSize(file, 1); moreThan {
			utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
			return
		}

		puskesmasFolder := "internal/files/puskesmas"
		filename = utils.RandomFileName(file)
		pathFile := filepath.Join(puskesmasFolder, filename)

		if err := utils.SaveUploadedFile(file, pathFile); err != nil {
			if err := os.Remove(pathFile); err != nil {
				log.Println(err.Error())
			}
			utils.HandleError(c, response.HANDLER_INTERR)
			return
		}
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateHealthCenter(uuid, filename, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Puskesmas Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateCase(c *gin.Context) {
	var body request.Case
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateCase(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kasus Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateCase(c *gin.Context) {
	var body request.Case
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateCase(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kasus Berhasil Diperbarui"))
}

func (h *ManagementHandler) DeleteCase(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteCase(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kasus Berhasil Dihapus"))
}
