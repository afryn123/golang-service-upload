package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type LogUpload struct {
	gorm.Model
	FileName     string         `gorm:"column:file_name" json:"file_name" binding:"required"`
	TotalRows    int            `gorm:"column:total_rows" json:"total_rows" binding:"required"`
	TotalSuccess int            `gorm:"column:total_success" json:"total_success" binding:"required"`
	TotalFailed  int            `gorm:"column:total_failed" json:"total_failed" binding:"required"`
	ErrorJson    datatypes.JSON `gorm:"column:error_json" json:"error_json" binding:"required"`
}
