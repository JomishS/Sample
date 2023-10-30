package model

import (
	"time"
)


type Document struct {
	Id   int    `gorm:"primary_key;column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title" validate:"required"` 
	Format   string `gorm:"column:format" json:"format" validate:"required"` 
	Author   string `gorm:"column:author" json:"author" validate:"required"`
	Owner    string `gorm:"column:owner" json:"owner" validate:"required"`
	Validity string `gorm:"column:validity" json:"validity" validate:"required"`
	CreatedAt	time.Time	`gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   *time.Time  `gorm:"column:deleted_at" json:"deletedAt"`
}

