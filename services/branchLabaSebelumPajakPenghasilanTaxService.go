package services

import (
	"afryn123/technical-test-go/models"
	"afryn123/technical-test-go/repositories"
	"afryn123/technical-test-go/utils"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UploadErrorLog struct {
	Row     int    `json:"row"`
	Message string `json:"message"`
}

type BranchLabaSebelumPajakPenghasilanTaxService interface {
	GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, string, bool, error)
	ImportExcel(file multipart.File, filename string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error)
}

type BranchLabaSebelumPajakPenghasilanTaxServiceImpl struct {
	DB                                             *gorm.DB
	BranchLabaSebelumPajakPenghasilanTaxRepository repositories.BranchLabaSebelumPajakPenghasilanTaxRepository
	LogUploadRepository                            repositories.LogUploadRepository
	validate                                       *validator.Validate
}

func NewBranchLabaSebelumPajakPenghasilanTaxService(db *gorm.DB, repo repositories.BranchLabaSebelumPajakPenghasilanTaxRepository, logUpload repositories.LogUploadRepository) BranchLabaSebelumPajakPenghasilanTaxService {
	return &BranchLabaSebelumPajakPenghasilanTaxServiceImpl{
		DB: db,
		BranchLabaSebelumPajakPenghasilanTaxRepository: repo,
		LogUploadRepository:                            logUpload,
		validate:                                       validator.New(),
	}
}

func (s *BranchLabaSebelumPajakPenghasilanTaxServiceImpl) GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, string, bool, error) {
	return s.BranchLabaSebelumPajakPenghasilanTaxRepository.GetDistinctPeriodeData(s.DB, limit, lastPeriode)
}

func (s *BranchLabaSebelumPajakPenghasilanTaxServiceImpl) ImportExcel(file multipart.File, filename string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error) {
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	sheet := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var (
		result    []*models.BranchLabaSebelumPajakPenghasilanTax
		errorLogs []UploadErrorLog
		totalRows = len(rows) - 1
		success   = 0
	)

	tx := s.DB.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %w", tx.Error)
	}

	for i, row := range rows {
		rowNumber := i + 1
		if i == 0 || len(row) < 3 {
			continue
		}

		var errorMessages []string

		nilai, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			errorMessages = append(errorMessages, "Nilai harus berupa angka")
		}

		if !utils.IsValidDateFormat(row[1]) {
			errorMessages = append(errorMessages, "Periode harus berupa format YYYY-MM-DD")
		}

		periode, err := utils.ParseDate(row[1])
		if err != nil {
			errorMessages = append(errorMessages, "Periode tidak valid")
		}

		if len(errorMessages) > 0 {
			errorLogs = append(errorLogs, UploadErrorLog{
				Row:     rowNumber,
				Message: fmt.Sprintf("%s", utils.JoinMessages(errorMessages)),
			})
			continue
		}

		item := &models.BranchLabaSebelumPajakPenghasilanTax{
			LabelRekonsiliasiFiskal: row[0],
			Periode:                 periode,
			Nilai:                   nilai,
		}

		if err := s.BranchLabaSebelumPajakPenghasilanTaxRepository.SaveFromUpload(tx, item); err != nil {
			errorLogs = append(errorLogs, UploadErrorLog{
				Row:     rowNumber,
				Message: "Gagal simpan DB: " + err.Error(),
			})
			continue
		}

		success++
		result = append(result, item)
	}

	// Simpan log upload
	errorJSON, err := json.Marshal(errorLogs)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal marshal error json: %w", err)
	}

	logUpload := &models.LogUpload{
		FileName:     filename,
		TotalRows:    totalRows,
		TotalSuccess: success,
		TotalFailed:  totalRows - success,
		ErrorJson:    datatypes.JSON(errorJSON),
	}

	if err := s.LogUploadRepository.Save(tx, logUpload); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan log upload: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("gagal commit transaksi: %w", err)
	}

	return result, nil
}
