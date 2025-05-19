package config

import (
	"afryn123/technical-test-go/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	log.Println("👉 DSN:", dsn)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get generic DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf(" Failed to ping database: %v", err)
	}

	// Migrate
	if err := DB.AutoMigrate(&models.BranchLabaSebelumPajakPenghasilanTax{}, &models.LogUpload{}); err != nil {
		log.Fatalf(" Error migrating database: %v", err)
	}

	log.Println("✅ Database connected and migration successful")
}
