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
	"example/Project3/internal/model"
	// "example/Project3/db"
	"example/Project3/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Documentcontroller interface{
	CreateDoc(c *gin.Context)
	GetByIdDoc(ctx *gin.Context)
	GetAllDoc(ctx *gin.Context)
	UpdateDoc(ctx *gin.Context)
	DeleteDoc(ctx *gin.Context)
}

type Doccon struct{
	Docserv Service.Documentservice
}

func DocumentConFunc(Docserv Service.Documentservice) Documentcontroller{
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
		context.JSON(http.StatusOK,"error in binding")
		fmt.Println("hello inside param context")
	}
	println("inside add function")
	crtdocu,err:=c.Docserv.CreateDoc(context,params)
	if err!=nil{
			context.JSON(http.StatusBadRequest, gin.H{"error": "error in creating"})
			fmt.Println("inside crreate doc return")
			return
		
	}
	 context.IndentedJSON(http.StatusOK,crtdocu)
	
}
	
	
func (c *Doccon) GetByIdDoc(ctx *gin.Context) {
	var statuscode int
	id := ctx.Param("id")
	params, err := c.Docserv.GetByIdDoc(ctx, id)
	fmt.Println("after  request")
	// err := ctx.ShouldBindJSON(&params)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, params)

}

	
func (c *Doccon) GetAllDoc(ctx *gin.Context) {
	var statuscode int
	todo, err := c.Docserv.GetAllDoc(ctx)
	if err != nil {
		statuscode = http.StatusBadRequest
		ctx.JSON(statuscode, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, todo)

}

func (c *Doccon) UpdateDoc(ctx *gin.Context) {

	// var	 statuscode int
	id := ctx.Param("id")
	var paramss model.Document
	err := ctx.ShouldBindJSON(&paramss)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in update"})
	}
	todo,err := c.Docserv.UpdateDoc(ctx, id, paramss)
	// err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in finding")
		return
	}
	ctx.JSON(http.StatusOK, todo)
}
func (c *Doccon) DeleteDoc(ctx *gin.Context) {

	id := ctx.Param("id")

	err := c.Docserv.DeleteDoc(ctx,id)
	// err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error in deleting")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " deleted"})
}







// func SortDoc(context *gin.Context) {
// 	println("getting document details")
// 	param := context.Param("param")
	
// 	db :=database.SetupDB()
// 	context.Header("Content-Type", "application/json")
	
	
// 	stmt, err := db.Prepare("SELECT id,title,format,author,owner,validity FROM document WHERE deleted=false ORDER BY " + param)
// 	//here we are referring to a column so we need to use prepare statement like this .In the other case we are just referering to a value.
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
	
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
	
	
// 	var doc []model.Document
// 	for rows.Next() {
// 		var a model.Document
// 		err := rows.Scan(&a.Id, &a.Title, &a.Format, &a.Author, &a.Owner, &a.Validity)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		doc = append(doc, a)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	context.IndentedJSON(http.StatusOK, doc)
// 	}
