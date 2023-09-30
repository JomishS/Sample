package repository

import (
	// "example/Project3/database"
	"example/Project3/internal/model"
	"fmt"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"context"
	// "fmt"
	// "strconv"
	// "errors"
	// "github.com/go-playground/validator/v10"
)

type Documentrepository interface {
	CreateDoc(c context.Context, params *model.Document) (*model.Document, error)
	GetByIdDoc(ctx context.Context, id string) (model.Document, error)
	GetAllDoc(ctx context.Context) ([]model.Document, error)
	UpdateDoc(ctx context.Context, id string, paramss model.Document) (model.Document, error)
	DeleteDoc(ctx context.Context, id string) error
}

type Docrepo struct {
	// Datab database
	Db *gorm.DB
}

func DocumentRepFunc(Db *gorm.DB) Documentrepository {
	return &Docrepo{
		Db: Db,
	}
}

func (c *Docrepo) CreateDoc(context context.Context, params *model.Document) (*model.Document, error) {

	fmt.Println("inside repository")
	fmt.Println(c.Db)
	err := validateDocument(params)
	if err!=nil{
		return nil,err
	}
	if err := c.Db.Create(params).Error; err != nil {
		fmt.Println("inside error")
		// return model.Document{}, err
		return nil, err

	}
	fmt.Println("outside error")
	return params, nil
}

func validateDocument(document *model.Document) error {
	validate := validator.New()
	if err := validate.Struct(document); err != nil {
		return err
	}
	return nil
}

func (c *Docrepo) GetByIdDoc(ctx context.Context, id string) (plant model.Document, err error) {
	err = c.Db.First(&plant, id).Error
	if err != nil {
		return plant, err
	}

	return plant, err
}

func (c *Docrepo) GetAllDoc(ctx context.Context) ([]model.Document, error) {

	var err error
	documents := []model.Document{}

	err = c.Db.Where("deletedat IS NULL").Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return documents, err
}

func (c *Docrepo) UpdateDoc(ctx context.Context, id string, paramss model.Document) (model.Document, error) {
	fmt.Println(id)
	var update model.Document
	if err := c.Db.First(&update, id).Error; err != nil {
		return model.Document{}, err
	}
	update.Doc_id = paramss.Doc_id
	update.Format = paramss.Format
	update.Owner = paramss.Owner
	update.Validity = paramss.Validity
	// err := databseconn.Getdb().Where("id=?", id).Update("plant",paramss.Plant).Error
	// fmt.Println("repository")
	if err := c.Db.Save(&update).Error; err != nil {
		return model.Document{}, err
	}
	return update, nil
}
func (c *Docrepo) DeleteDoc(ctx context.Context, id string) error {
	var document *model.Document
	if err := c.Db.First(&document, id).Error; err != nil {
		return err
	}
	current_time := time.Now()
	document.DeletedAt = &current_time

	err := c.Db.Save(&document).Error
	if err != nil {
		return err
	}
	return nil

}
