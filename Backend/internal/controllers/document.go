package controllers

import (
	"fmt"
	"net/http"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	Service "example/Project3/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Documentcontroller interface {
	CreateDoc(c *gin.Context)
	GetByIdDoc(ctx *gin.Context)
	GetAllDoc(ctx *gin.Context)
	UpdateDoc(ctx *gin.Context)
	DeleteDoc(ctx *gin.Context)
}

type Doccon struct {
	Docserv Service.Documentservice
}

func DocumentConFunc(Docserv Service.Documentservice) Documentcontroller {
	return &Doccon{
		Docserv: Docserv,
	}
}

func (c *Doccon) CreateDoc(context *gin.Context) {

	context.Header("Content_Type", "application/json")

	var params *model.Document
	fmt.Println(params)
	err := context.ShouldBindJSON(&params)
	fmt.Println(params)
	if err != nil {
		context.JSON(http.StatusOK, "error in binding")
		fmt.Println("hello inside param context")
	}
	println("inside add function")
	crtdocu, err := c.Docserv.CreateDoc(context, params)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error in creating"})
		fmt.Println("inside crreate doc return")
		return

	}
	context.IndentedJSON(http.StatusOK, crtdocu)

}

func (c *Doccon) GetByIdDoc(ctx *gin.Context) {
	var statuscode int
	id := ctx.Param("id")
	params, err := c.Docserv.GetByIdDoc(ctx, id)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, params)

}

func (c *Doccon) GetAllDoc(ctx *gin.Context) {
	var statuscode int
	params := &dto.AssetQueryParams{}
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, "error in query parameter")
		return
	}

	todo, err := c.Docserv.GetAllDoc(ctx, params)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, todo)

}

func (c *Doccon) UpdateDoc(ctx *gin.Context) {

	id := ctx.Param("id")
	var paramss model.Document
	err := ctx.ShouldBindJSON(&paramss)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in update"})
	}
	todo, err := c.Docserv.UpdateDoc(ctx, id, paramss)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, todo)
}
func (c *Doccon) DeleteDoc(ctx *gin.Context) {

	id := ctx.Param("id")

	err := c.Docserv.DeleteDoc(ctx, id)
	// err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in deleting")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " deleted"})
}
