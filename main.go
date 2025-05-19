package main

import (
	"afryn123/technical-test-go/config"
	"afryn123/technical-test-go/middlewares"
	"afryn123/technical-test-go/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middlewares.CustomRecoverPanic())
	r.Use(cors.Default())

	routes.BranchLabaSebelumPajakPenghasilanTaxRoutes(r)
	r.Run()
}
