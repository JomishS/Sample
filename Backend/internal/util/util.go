package util

import (
	"errors"
	"example/Project3/internal/model"
	"example/Project3/internal/dto"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SplitIntoArray(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "|")
}

func IsEmptyStringOrZero(i interface{}) bool {
	if v, ok := i.(string); ok {
		return v == ""
	}
	if v, ok := i.(int); ok {
		return v == 0
	}
	if v, ok := i.(int64); ok {
		return v == 0
	}
	if v, ok := i.(float64); ok {
		return v == 0
	}

	return false
}

func SplitStringFromBack(str, delimiter string) (string, string) {
	lastIndex := strings.LastIndex(str, delimiter)
	if lastIndex == -1 {
		return "", str
	}
	return str[:lastIndex], str[lastIndex+len(delimiter):]
}

func ReadTag(modelStruct model.Document, fieldName string, tag string) (string, error) {
	field, ok := reflect.TypeOf(modelStruct).Elem().FieldByName(fieldName)
	if !ok {
		return "", errors.New("field not found")
	}
	return field.Tag.Get(tag), nil
}

func ReadTagUser(modelStruct model.User, fieldName string, tag string) (string, error) {
	field, ok := reflect.TypeOf(modelStruct).Elem().FieldByName(fieldName)
	if !ok {
		return "", errors.New("field not found")
	}
	return field.Tag.Get(tag), nil
}

func GetTableName(db *gorm.DB, doc model.Document) (string, error) {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(doc)
	if err != nil {
		return "", err
	}
	return stmt.Schema.Table, nil
}

func GetTableNameUser(db *gorm.DB, user model.User) (string, error) {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(user)
	if err != nil {
		return "", err
	}
	return stmt.Schema.Table, nil
}

func GetTotalCount(txn *gorm.DB) (totalCount int, err error) {
	var count int64
	err = txn.Count(&count).Error
	return int(count), err
}

func SuccessResponseWithCount(ctx *gin.Context, statusCode int, message string, count int, data interface{}) {
	statusResponse := &dto.Status{Code: statusCode, Message: message, TotalCount: &count, Type: "success"}
	successResponse := &dto.BaseResponse{Status: statusResponse, Data: data}
	ctx.JSON(statusCode, successResponse)
}
