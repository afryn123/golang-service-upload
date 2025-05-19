package repositories

import (
	"afryn123/technical-test-go/config"
	"afryn123/technical-test-go/models"
	"errors"

	"gorm.io/gorm"
)

type BranchLabaSebelumPajakPenghasilanTaxRepository interface {
	GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error)
	SaveFromUpload(data *models.BranchLabaSebelumPajakPenghasilanTax) error
}

type BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl struct {
	DB *gorm.DB
}

func NewBranchLabaSebelumPajakPenghasilanTaxRepository() BranchLabaSebelumPajakPenghasilanTaxRepository {
	if config.DB == nil {
		panic("database connection is nil")
	}
	return &BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl{
		DB: config.DB,
	}
}

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error) {
	if r.DB == nil {
		return nil, errors.New("database connection is nil")
	}

	var results []*models.BranchLabaSebelumPajakPenghasilanTax

	subQuery := r.DB.
		Model(&models.BranchLabaSebelumPajakPenghasilanTax{}).
		Select("MIN(id)").
		Group("periode")

	query := r.DB.
		Model(&models.BranchLabaSebelumPajakPenghasilanTax{}).
		Where("id IN (?)", subQuery).
		Order("periode ASC").
		Limit(limit)

	if lastPeriode != "" {
		query = query.Where("DATE(periode) > ?", lastPeriode)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil

}

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) SaveFromUpload(data *models.BranchLabaSebelumPajakPenghasilanTax) error {
	return r.DB.Create(data).Error
}
