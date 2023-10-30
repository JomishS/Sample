package repository

import (
	"example/Project3/internal/constant"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"context"
	"example/Project3/internal/util"
)

type Documentrepository interface {
	CreateDoc(c context.Context, params *model.Document) (*model.Document, error)
	GetByIdDoc(ctx context.Context, id string) (model.Document, error)
	GetAllDoc(ctx context.Context, query *dto.AssetQueryParams, filterMap dto.SearchFilters) (doc []model.Document, totalCount int, err error)
	UpdateDoc(ctx context.Context, id string, paramss model.Document) (model.Document, error)
	DeleteDoc(ctx context.Context, id string) error
}

type Docrepo struct {
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
	if err != nil {
		return nil, err
	}
	if err := c.Db.Create(params).Error; err != nil {
		fmt.Println("inside error")
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
	err = c.Db.Where("deleted_at IS NULL").First(&plant, id).Error
	if err != nil {
		return plant, err
	}

	return plant, err
}

func (c *Docrepo) GetAllDoc(ctx context.Context, query *dto.AssetQueryParams, filterMap dto.SearchFilters) (doc []model.Document, totalCount int, err error) {

	var document model.Document
	txn := c.Db.Debug().WithContext(ctx).Model(document)
	test := txn.Model(document)
	params:=test.Scopes(util.ColumnValCheck(constant.ColDeletedAt))

	wherePredicates, err := util.GetWherePredicates(params, filterMap, document)
	if err != nil {
		return nil, 0, err
	}

	paginationPredicates := util.GetPaginationPredicates(query.Page, query.Limit)
	sortPredicates := util.GetSortPredicates(query.SortBy)

	 params = util.AddScopes(txn, wherePredicates)
	totalCount, err = util.GetTotalCount(params)
	if err != nil {
		return nil, 0, err
	}

	params = util.AddScopes(params, sortPredicates, paginationPredicates)
	err = params.Find(&doc).Error

	return doc, totalCount, err


}

func (c *Docrepo) UpdateDoc(ctx context.Context, id string, paramss model.Document) (model.Document, error) {
	fmt.Println(id)
	var update model.Document
	if err := c.Db.First(&update, id).Error; err != nil {
		return model.Document{}, err
	}
	update.Format = paramss.Format
	update.Owner = paramss.Owner
	update.Validity = paramss.Validity
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
