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

	branchLabaSebelumPajakPenghasilanTaxService := services.NewBranchLabaSebelumPajakPenghasilanTaxService(
		config.DB,
		branchLabaSebelumPajakPenghasilanTaxRepository,
		logUploadRepository,
	)
	logUploadService := services.NewLogUploadService(config.DB, logUploadRepository)

	// Controller
	branchLabaSebelumPajakPenghasilanTaxController := controllers.NewBranchLabaSebelumPajakPenghasilanTaxController(branchLabaSebelumPajakPenghasilanTaxService)
	logUploadController := controllers.NewLogUploadController(logUploadService)

	// Routes
	api := r.Group("/api")
	api.GET("/getData", branchLabaSebelumPajakPenghasilanTaxController.GetDistinctPeriodeData)
	api.POST("/uploadData", branchLabaSebelumPajakPenghasilanTaxController.UploadExcel)
	api.GET("/getDataAll", branchLabaSebelumPajakPenghasilanTaxController.GetAllDistinctData)

	api.GET("/logUpload", logUploadController.GetAllLog)
}
