package repository

import (
	"example/Project3/internal/constant"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	"example/Project3/internal/util"
	"fmt"
	"time"
	"context"
	"example/Project3/database"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Userrepository interface {
	CreateUser(c context.Context, params model.User) (model.User, error)
	GetByIdUser(ctx context.Context, id string) (model.User, error)
	GetAllUser(ctx context.Context, query *dto.AssetQueryParams, filterMap dto.SearchFilters) (user []model.User, totalCount int, err error)
	UpdateUser(ctx context.Context, id string, paramss model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type Userrepo struct {
	Db *gorm.DB
}

func UserRepFunc(Db *gorm.DB) Userrepository {
	return &Userrepo{
		Db: Db,
	}
}

func (c *Userrepo) CreateUser(ctx context.Context, params model.User) (model.User, error) {
	// var user model.User

	if err := validateUser(&params); err != nil {
		return model.User{}, err
	}
	if err := database.GetDB().Create(&params).Error; err != nil {
		return model.User{}, err
	}

	return params, nil
}

func validateUser(user *model.User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}

func (c *Userrepo) GetByIdUser(ctx context.Context, id string) (user model.User, err error) {
	err = c.Db.Where("deleted_at IS NULL").First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (c *Userrepo) GetAllUser(ctx context.Context, query *dto.AssetQueryParams, filterMap dto.SearchFilters) (user []model.User, totalCount int, err error) {

	var userd model.User
	txn := c.Db.Debug().WithContext(ctx).Model(userd)
	test := txn.Model(userd)
	params := test.Scopes(util.ColumnValCheck(constant.ColDeletedAt))

	wherePredicates, err := util.GetWherePredicatesUser(params, filterMap, userd)
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

	err = params.Find(&user).Error

	return user, totalCount, err
}

func (c *Userrepo) UpdateUser(ctx context.Context, id string, paramss model.User) error {
	fmt.Println(id)
	var update model.User
	if err := c.Db.First(&update, id).Error; err != nil {
		return err
	}
	update.Age = paramss.Age
	update.Email = paramss.Email
	update.City = paramss.City
	update.Country = paramss.Country
	update.Phone = paramss.Phone
	if err := c.Db.Save(&update).Error; err != nil {
		return err
	}
	return nil
}
func (c *Userrepo) DeleteUser(ctx context.Context, id string) error {

	var user *model.User
	if err := c.Db.First(&user, id).Error; err != nil {
		return err
	}
	current_time := time.Now()
	user.DeletedAt = &current_time

	err := c.Db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
