package model

import "time"


// type User struct {
// 	gorm.Model
// 	User_id     int    `json:"user_id" gorm:"primary_key"`
// 	First_name  string `json:"first_name" validate:"required"`
// 	Last_name   string `json:"last_name" validate:"required"`
// 	Age         int    `json:"age" validate:"required"`
// 	Email       string `json:"email" validate:"required"`
// 	City        string `json:"city" validate:"required"`
// 	Phone       string `json:"phone" validate:"required"`
// 	Birth_date  string `json:"birth_date" validate:"required"`
// 	Sex         string `json:"sex" validate:"required"`
// 	Country     string `json:"country" validate:"required"`
// 	Document_id int    `json:"document_id" validate:"required"`
// }



type User struct {
	User_id     int    `gorm:"primary_key;column:user_id" json:"user_id"`
	First_name  string `gorm:"column:first_name" json:"first_name" validate:"required"`
	Last_name   string `gorm:"column:last_name" json:"last_name" validate:"required"`
	Age         int    `gorm:"column:age" json:"age" validate:"required"`
	Email       string `gorm:"column:email" json:"email" validate:"required"`
	City        string `gorm:"column:city" json:"city" validate:"required"`
	Phone       string `gorm:"column:phone" json:"phone" validate:"required"`
	Birth_date  string `gorm:"column:birth_date" json:"birth_date" validate:"required"`
	Sex         string `gorm:"column:sex" json:"sex" validate:"required"`
	Country     string `gorm:"column:country" json:"country" validate:"required"`
	Document_id int    `gorm:"column:document_id" json:"document_id" validate:"required"`
	CreatedAt	time.Time	`gorm:"column:createdat" json:"createdat"`
	UpdatedAt   time.Time   `gorm:"column:updatedat" json:"updatedat"`
	DeletedAt   *time.Time  `gorm:"column:deletedat" json:"deletedat"`
}