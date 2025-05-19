package controllers

import (
	"afryn123/technical-test-go/services"
	"afryn123/technical-test-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchLabaSebelumPajakPenghasilanTaxController interface {
	GetDistinctPeriodeData(ctx *gin.Context)
	UploadExcel(c *gin.Context)
}

type BranchLabaSebelumPajakPenghasilanTaxControllerImpl struct {
	BranchLabaSebelumPajakPenghasilanTaxService services.BranchLabaSebelumPajakPenghasilanTaxService
}

func NewBranchLabaSebelumPajakPenghasilanTaxController(service services.BranchLabaSebelumPajakPenghasilanTaxService) BranchLabaSebelumPajakPenghasilanTaxController {
	return &BranchLabaSebelumPajakPenghasilanTaxControllerImpl{
		BranchLabaSebelumPajakPenghasilanTaxService: service,
	}
}

func (c *BranchLabaSebelumPajakPenghasilanTaxControllerImpl) GetDistinctPeriodeData(ctx *gin.Context) {
	lastPeriode := ctx.Query("last_periode")
	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "invalid limit", err)
		return
	}

	data, err := c.BranchLabaSebelumPajakPenghasilanTaxService.GetDistinctPeriodeData(limit, lastPeriode)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "error fetching data", err)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "success", data)
}

func (rc *BranchLabaSebelumPajakPenghasilanTaxControllerImpl) UploadExcel(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file", err)
		return
	}

	if !utils.IsValidExcelFile(fileHeader) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file format", "file must be .xlsx")
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to open file", err)
		return
	}
	defer file.Close()

	data, logErr, err := rc.BranchLabaSebelumPajakPenghasilanTaxService.ImportExcel(file)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to import data", err)
		return
	}

	if len(logErr) > 0 {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to import data", logErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Upload successful", "data": data})
}
