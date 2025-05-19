package routes

import (
	"afryn123/technical-test-go/controllers"
	"afryn123/technical-test-go/repositories"
	"afryn123/technical-test-go/services"

	"github.com/gin-gonic/gin"
)

func BranchLabaSebelumPajakPenghasilanTaxRoutes(r *gin.Engine) {
	branchLabaSebelumPajakPenghasilanTaxRepository := repositories.NewBranchLabaSebelumPajakPenghasilanTaxRepository()
	branchLabaSebelumPajakPenghasilanTaxService := services.NewBranchLabaSebelumPajakPenghasilanTaxService(branchLabaSebelumPajakPenghasilanTaxRepository)
	branchLabaSebelumPajakPenghasilanTaxController := controllers.NewBranchLabaSebelumPajakPenghasilanTaxController(branchLabaSebelumPajakPenghasilanTaxService)
	api := r.Group("/api")

	api.GET("/getData", branchLabaSebelumPajakPenghasilanTaxController.GetDistinctPeriodeData)
	api.GET("/uploadData", branchLabaSebelumPajakPenghasilanTaxController.UploadExcel)
}
