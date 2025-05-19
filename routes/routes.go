package routes

import (
	"afryn123/technical-test-go/config"
	"afryn123/technical-test-go/controllers"
	"afryn123/technical-test-go/repositories"
	"afryn123/technical-test-go/services"

	"github.com/gin-gonic/gin"
)

func BranchLabaSebelumPajakPenghasilanTaxRoutes(r *gin.Engine) {
	// Repository
	branchLabaSebelumPajakPenghasilanTaxRepository := repositories.NewBranchLabaSebelumPajakPenghasilanTaxRepository()
	logUploadRepository := repositories.NewLogUploadRepository()

	// Service — inject DB dan dua repo
	branchLabaSebelumPajakPenghasilanTaxService := services.NewBranchLabaSebelumPajakPenghasilanTaxService(
		config.DB,
		branchLabaSebelumPajakPenghasilanTaxRepository,
		logUploadRepository,
	)

	// Controller
	branchLabaSebelumPajakPenghasilanTaxController := controllers.NewBranchLabaSebelumPajakPenghasilanTaxController(branchLabaSebelumPajakPenghasilanTaxService)

	// Routes
	api := r.Group("/api")
	api.GET("/getData", branchLabaSebelumPajakPenghasilanTaxController.GetDistinctPeriodeData)
	api.POST("/uploadData", branchLabaSebelumPajakPenghasilanTaxController.UploadExcel)
}
