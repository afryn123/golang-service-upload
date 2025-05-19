package controllers

import (
	"afryn123/technical-test-go/models"
	"afryn123/technical-test-go/services"
	"afryn123/technical-test-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Items       []*models.BranchLabaSebelumPajakPenghasilanTax `json:"items"`
	LastPeriode string                                         `json:"last_periode"`
	HasNext     bool                                           `json:"has_next"`
}
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

	items, lastPeriodeResult, hasMore, err := c.BranchLabaSebelumPajakPenghasilanTaxService.GetDistinctPeriodeData(limit, lastPeriode)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "error fetching data", err)
		return
	}

	data := Data{
		Items:       items,
		LastPeriode: lastPeriodeResult,
		HasNext:     hasMore,
	}

	utils.JSONResponse(ctx, http.StatusOK, "success", data)
}

func (rc *BranchLabaSebelumPajakPenghasilanTaxControllerImpl) UploadExcel(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file", err)
		return
	}
	filename := fileHeader.Filename

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

	data, err := rc.BranchLabaSebelumPajakPenghasilanTaxService.ImportExcel(file, filename)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to import data", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Upload successful", "data": data})
}
