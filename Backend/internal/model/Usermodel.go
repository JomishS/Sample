package model

import "time"



type User struct {
	Id     int    `gorm:"primary_key;column:id" json:"id"`
	First_name  string `gorm:"column:first_name" json:"first_name" validate:"required"`
	Last_name   string `gorm:"column:last_name" json:"last_name" validate:"required"`
	Age         int    `gorm:"column:age" json:"age" validate:"required"`
	Email       string `gorm:"column:email" json:"email" validate:"required"`
	City        string `gorm:"column:city" json:"city" validate:"required"`
	Phone       string `gorm:"column:phone" json:"phone" validate:"required"`
	Birth_date  string `gorm:"column:birth_date" json:"birth_date" validate:"required"`
	Sex         string `gorm:"column:sex" json:"sex" validate:"required"`
	Country     string `gorm:"column:country" json:"country" validate:"required"`
	Doc_id int    `gorm:"column:doc_id" json:"doc_id" validate:"required"`
	CreatedAt	time.Time	`gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   *time.Time  `gorm:"column:deleted_at" json:"deletedAt"`
}