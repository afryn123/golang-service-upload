package models

import (
	"gorm.io/gorm"
)

type BranchLabaSebelumPajakPenghasilanTax struct {
	gorm.Model
	LabelRekonsiliasiFiskal string  `gorm:"column:label_rekonsiliasi_fiskal" json:"label_rekonsiliasi_fiskal" binding:"required"`
	Periode                 string  `json:"periode" binding:"required,datetime=2006-01-02"`
	Nilai                   float64 `json:"nilai" binding:"required"`
}

// TableName overrides the table name used by GORM
func (BranchLabaSebelumPajakPenghasilanTax) TableName() string {
	return "branch_laba_sebelum_pajak_penghasilan_tax"
}
