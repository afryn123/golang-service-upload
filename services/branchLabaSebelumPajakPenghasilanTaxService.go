package services

import (
	"afryn123/technical-test-go/models"
	"afryn123/technical-test-go/repositories"
	"afryn123/technical-test-go/utils"
	"mime/multipart"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
)

type BranchLabaSebelumPajakPenghasilanTaxService interface {
	GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error)
	ImportExcel(file multipart.File) ([]*models.BranchLabaSebelumPajakPenghasilanTax, map[int]string, error)
}

type BranchLabaSebelumPajakPenghasilanTaxServiceImpl struct {
	BranchLabaSebelumPajakPenghasilanTaxRepository repositories.BranchLabaSebelumPajakPenghasilanTaxRepository
	validate                                       *validator.Validate
}

func NewBranchLabaSebelumPajakPenghasilanTaxService(repo repositories.BranchLabaSebelumPajakPenghasilanTaxRepository) BranchLabaSebelumPajakPenghasilanTaxService {
	return &BranchLabaSebelumPajakPenghasilanTaxServiceImpl{
		BranchLabaSebelumPajakPenghasilanTaxRepository: repo,
		validate: validator.New(),
	}
}

func (s *BranchLabaSebelumPajakPenghasilanTaxServiceImpl) GetDistinctPeriodeData(limit int, lastPeriode string) ([]*models.BranchLabaSebelumPajakPenghasilanTax, error) {
	return s.BranchLabaSebelumPajakPenghasilanTaxRepository.GetDistinctPeriodeData(limit, lastPeriode)
}

func (s *BranchLabaSebelumPajakPenghasilanTaxServiceImpl) ImportExcel(file multipart.File) ([]*models.BranchLabaSebelumPajakPenghasilanTax, map[int]string, error) {
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return nil, nil, err
	}

	sheet := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheet)
	if err != nil {
		return nil, nil, err
	}

	var result []*models.BranchLabaSebelumPajakPenghasilanTax
	errorLogs := make(map[int]string)

	for i, row := range rows {
		rowNumber := i + 2
		if i == 0 || len(row) < 3 {
			continue
		}

		nilai, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			errorLogs[rowNumber] = "Nilai harus berupa angka"
			continue
		}

		if !utils.IsValidDateFormat(row[1]) {
			errorLogs[rowNumber] = "Periode harus berupa format YYYY-MM-DD"
			continue
		}

		periode, err := utils.ParseDate(row[1])
		if err != nil {
			errorLogs[rowNumber] = "Periode tidak valid"
			continue
		}

		item := &models.BranchLabaSebelumPajakPenghasilanTax{
			LabelRekonsiliasiFiskal: row[0],
			Periode:                 periode,
			Nilai:                   nilai,
		}

		err = s.BranchLabaSebelumPajakPenghasilanTaxRepository.SaveFromUpload(item)
		if err != nil {
			return nil, nil, err
		}

		result = append(result, item)
	}

	return result, errorLogs, nil
}
