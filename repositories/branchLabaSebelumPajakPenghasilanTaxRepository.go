package repositories

import (
	"afryn123/technical-test-go/models"

	"gorm.io/gorm"
)

type BranchLabaSebelumPajakPenghasilanTaxRepository interface {
	GetDistinctPeriodeData(
		tx *gorm.DB, limit int, lastPeriode string,
	) ([]*models.BranchLabaSebelumPajakPenghasilanTax, string, bool, error)
	SaveFromUpload(tx *gorm.DB, data *models.BranchLabaSebelumPajakPenghasilanTax) error
}

type BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl struct{}

func NewBranchLabaSebelumPajakPenghasilanTaxRepository() BranchLabaSebelumPajakPenghasilanTaxRepository {
	return &BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl{}
}

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) GetDistinctPeriodeData(
	tx *gorm.DB, limit int, lastPeriode string,
) ([]*models.BranchLabaSebelumPajakPenghasilanTax, string, bool, error) {
	subQuery := tx.
		Model(&models.BranchLabaSebelumPajakPenghasilanTax{}).
		Select("MIN(id)").
		Group("periode")

	query := tx.
		Model(&models.BranchLabaSebelumPajakPenghasilanTax{}).
		Where("id IN (?)", subQuery).
		Order("periode ASC")

	if lastPeriode != "" {
		query = query.Where("DATE(periode) > ?", lastPeriode)
	}

	var tempResults []*models.BranchLabaSebelumPajakPenghasilanTax
	if err := query.Limit(limit + 1).Find(&tempResults).Error; err != nil {
		return nil, "", false, err
	}

	hasMore := false
	if len(tempResults) > limit {
		hasMore = true
		tempResults = tempResults[:limit] // potong ke jumlah asli
	}

	// Ambil lastPeriode dari hasil terakhir
	lastPeriodeResult := ""
	if len(tempResults) > 0 {
		lastPeriodeResult = tempResults[len(tempResults)-1].Periode
	}

	return tempResults, lastPeriodeResult, hasMore, nil
}

func (r *BranchLabaSebelumPajakPenghasilanTaxRepositoryImpl) SaveFromUpload(tx *gorm.DB, data *models.BranchLabaSebelumPajakPenghasilanTax) error {
	return tx.Create(data).Error
}
