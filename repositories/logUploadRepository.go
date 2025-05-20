package repositories

import (
	"afryn123/technical-test-go/models"

	"gorm.io/gorm"
)

type LogUploadRepository interface {
	Save(tx *gorm.DB, log *models.LogUpload) error
	GetLastLog(tx *gorm.DB) (*models.LogUpload, error)
	GetAllLog(tx *gorm.DB) ([]*models.LogUpload, error)
	GetLogByDate(tx *gorm.DB, date string) (*models.LogUpload, error)
}

type LogUploadRepositoryImpl struct {
}

func NewLogUploadRepository() LogUploadRepository {
	return &LogUploadRepositoryImpl{}
}

func (r *LogUploadRepositoryImpl) Save(tx *gorm.DB, log *models.LogUpload) error {
	return tx.Create(log).Error
}

func (r *LogUploadRepositoryImpl) GetLastLog(tx *gorm.DB) (*models.LogUpload, error) {
	var log models.LogUpload
	if err := tx.Last(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *LogUploadRepositoryImpl) GetLogByDate(tx *gorm.DB, date string) (*models.LogUpload, error) {
	var log models.LogUpload
	if err := tx.Where("date = ?", date).First(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *LogUploadRepositoryImpl) GetAllLog(tx *gorm.DB) ([]*models.LogUpload, error) {
	var logs []*models.LogUpload
	if err := tx.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
