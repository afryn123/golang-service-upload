package repositories

import (
	"afryn123/technical-test-go/models"

	"gorm.io/gorm"
)

type LogUploadRepository interface {
	Save(tx *gorm.DB, log *models.LogUpload) error
}

type LogUploadRepositoryImpl struct {
}

func NewLogUploadRepository() LogUploadRepository {
	return &LogUploadRepositoryImpl{}
}

func (r *LogUploadRepositoryImpl) Save(tx *gorm.DB, log *models.LogUpload) error {
	return tx.Create(log).Error
}
