package Service

import (
	"context"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	"example/Project3/internal/repository"
	"example/Project3/internal/util"
)

type Userservice interface {
	CreateUser(c context.Context, params model.User) (model.User, error)
	UpdateUser(c context.Context, id string, params model.User) error
	DeleteUser(c context.Context, id string) error
	GetByIdUser(c context.Context, id string) (model.User, error)
	GetAllUser(c context.Context, query *dto.AssetQueryParams) ([]model.User, int, error)
}

type Userserv struct {
	Userrepo repository.Userrepository
}

func UserServFunc(Userrepo repository.Userrepository) Userservice {
	return &Userserv{
		Userrepo: Userrepo,
	}
}

func (c *Userserv) CreateUser(context context.Context, params model.User) (model.User, error) {
	params, err := c.Userrepo.CreateUser(context, params)
	return params, err
}

func (c *Userserv) GetByIdUser(ctx context.Context, id string) (model.User, error) {
	var err error

	plant, err := c.Userrepo.GetByIdUser(ctx, id)
	return plant, err
}

func (c *Userserv) GetAllUser(ctx context.Context, query *dto.AssetQueryParams) ([]model.User, int, error) {
	var err error
	filterMap, err := util.ParseFilters(query.Filter)
	if err != nil {
		return nil, 0, err
	}
	
	user, totalCount, err := c.Userrepo.GetAllUser(ctx, query, filterMap)
	return user, totalCount, err
}

func (c *Userserv) UpdateUser(ctx context.Context, id string, paramss model.User) error {
	err := c.Userrepo.UpdateUser(ctx, id, paramss)
	return err

}

func (c *Userserv) DeleteUser(ctx context.Context, id string) error {
	err := c.Userrepo.DeleteUser(ctx, id)
	return err

}
