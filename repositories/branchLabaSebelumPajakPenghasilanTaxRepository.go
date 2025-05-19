package repositories

import (
	"afryn123/technical-test-go/models"

	"gorm.io/gorm"
)

type BranchLabaSebelumPajakPenghasilanTaxRepository interface {
	GetDistinctPeriodeData(tx *gorm.DB, limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error)
	SaveFromUpload(tx *gorm.DB, data *models.BranchLabaSebelumPajakPenghasilanTax) error
}

type BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl struct{}

func NewBranchLabaSebelumPajakPenghasilanTaxRepository() BranchLabaSebelumPajakPenghasilanTaxRepository {
	return &BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl{}
}

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) GetDistinctPeriodeData(tx *gorm.DB, limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error) {

	var results []*models.BranchLabaSebelumPajakPenghasilanTax

	subQuery := tx.
		Model(&models.BranchLabaSebelumPajakPenghasilanTax{}).
		Select("MIN(id)").
		Group("periode")

	query := tx.
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

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) SaveFromUpload(tx *gorm.DB, data *models.BranchLabaSebelumPajakPenghasilanTax) error {
	return tx.Create(data).Error
}
