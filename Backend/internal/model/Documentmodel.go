package model

import (
	// "gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
	"time"
)

// type Document struct {
// 	gorm.Model
// 	Doc_id   int    `json:"doc_id" gorm:"primary_key"`
// 	Title    string `json:"Title" validate:"required"`
// 	Format   string `json:"format" validate:"required"`
// 	Author   string `json:"author" validate:"required"`
// 	Owner    string `json:"owner" validate:"required"`
// 	Validity string `json:"validity" validate:"required"`
// }

type Document struct {
	Doc_id   int    `gorm:"primary_key;column:doc_id" json:"doc_id"`
	Title    string `gorm:"column:title" json:"title" validate:"required"` 
	Format   string `gorm:"column:format" json:"format" validate:"required"` 
	Author   string `gorm:"column:author" json:"author" validate:"required"`
	Owner    string `gorm:"column:owner" json:"owner" validate:"required"`
	Validity string `gorm:"column:validity" json:"validity" validate:"required"`
	CreatedAt	time.Time	`gorm:"column:createdat" json:"createdat"`
	UpdatedAt   time.Time   `gorm:"column:updatedat" json:"updatedat"`
	DeletedAt   *time.Time  `gorm:"column:deletedat" json:"deletedat"`
}

