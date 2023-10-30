package controllers

import (
	// "database/sql"
	// "errors"
	// "log"

	// "encoding/json"
	// "fmt"
	"fmt"
	"net/http"

	// "strconv"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	"example/Project3/internal/util"

	// "example/Project3/db"
	Service "example/Project3/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Usercontroller interface {
	CreateUser(c *gin.Context)
	GetByIdUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type Usercon struct {
	Userserv Service.Userservice
}

func UserConFunc(Userserv Service.Userservice) Usercontroller {
	return &Usercon{
		Userserv: Userserv,
	}
}

func (c *Usercon) CreateUser(context *gin.Context) {
	context.Header("Content_Type", "application/json")
	var params model.User
	err := context.ShouldBindJSON(&params)
	if err != nil {
		context.JSON(http.StatusBadRequest, "error in binding")
	}
	todo, err := c.Userserv.CreateUser(context, params)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error in creating"})
		return
	}
	context.JSON(http.StatusCreated, todo)
}

func (c *Usercon) GetByIdUser(ctx *gin.Context) {
	fmt.Println("get request")

	var statuscode int
	id := ctx.Param("id")
	params, err := c.Userserv.GetByIdUser(ctx, id)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, params)

}

func (c *Usercon) GetAllUser(ctx *gin.Context) {
	var statuscode int
	params := &dto.AssetQueryParams{}
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, "error in query parameter")
		return
	}
	todo, count, err := c.Userserv.GetAllUser(ctx, params)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	util.SuccessResponseWithCount(ctx, http.StatusOK, "Successfully searched", count, todo)

}

func (c *Usercon) UpdateUser(ctx *gin.Context) {

	id := ctx.Param("id")
	var paramss model.User
	err := ctx.ShouldBindJSON(&paramss)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in update"})
	}
	err = c.Userserv.UpdateUser(ctx, id, paramss)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, paramss)
}
func (c *Usercon) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	err := c.Userserv.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in deleting")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " deleted"})
}
