package controllers

import (
	"afryn123/technical-test-go/services"
	"afryn123/technical-test-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogUploadController interface {
	GetAllLog(ctx *gin.Context)
}

type LogUploadControllerImpl struct {
	LogUploadService services.LogUploadService
}

func NewLogUploadController(logUploadService services.LogUploadService) LogUploadController {
	return &LogUploadControllerImpl{
		LogUploadService: logUploadService,
	}
}

func (c *LogUploadControllerImpl) GetAllLog(ctx *gin.Context) {
	logs, err := c.LogUploadService.GetAllLog()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "error fetching data", err)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "success", logs)
}
