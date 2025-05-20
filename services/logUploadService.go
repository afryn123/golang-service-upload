package services

import (
	"afryn123/technical-test-go/models"
	"afryn123/technical-test-go/repositories"

	"gorm.io/gorm"
)

type LogUploadService interface {
	GetAllLog() ([]*models.LogUpload, error)
}

type LogUploadServiceImpl struct {
	DB                  *gorm.DB
	LogUploadRepository repositories.LogUploadRepository
}

func NewLogUploadService(db *gorm.DB, repo repositories.LogUploadRepository) LogUploadService {
	return &LogUploadServiceImpl{
		DB:                  db,
		LogUploadRepository: repo,
	}
}

func (s *LogUploadServiceImpl) GetAllLog() ([]*models.LogUpload, error) {
	return s.LogUploadRepository.GetAllLog(s.DB)
}
