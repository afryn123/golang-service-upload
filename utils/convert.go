package utils

import (
	"strconv"
	"time"
)

// StringToUint mengubah string ke uint
func StringToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

// StringToUint mengubah string ke uint
func StringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func ParseDate(value string) (string, error) {
	t, err := time.Parse("2006-01-02", value)
	// Set lokasi ke Asia/Jakarta (UTC+7)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}
	t = t.In(loc)

	// Format waktu seperti yang kamu mau: "2006-01-02 15:04:05+07"
	formatted := t.Format("2006-01-02 15:04:05") + "+07"

	return formatted, nil
}
