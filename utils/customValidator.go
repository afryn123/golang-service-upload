package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

func IsValidExcelFile(fileHeader *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	return ext == ".xlsx"
}

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
