package Service

import (
	"context"
	"example/Project3/internal/model"
	"example/Project3/internal/repository"
	// "errors"
	// "log"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

type Userservice interface {
	CreateUser(c context.Context, params model.User) (model.User, error)
	UpdateUser(c context.Context, id string, params model.User) error
	DeleteUser(c context.Context, id string) error
	GetByIdUser(c context.Context, id string) (model.User, error)
	GetAllUser(c context.Context) ([]model.User, error)
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


func (c *Userserv) GetAllUser(ctx context.Context) ([]model.User, error) {
	var err error

	plant, err := c.Userrepo.GetAllUser(ctx)
	return plant, err

}

func (c *Userserv) UpdateUser(ctx context.Context, id string, paramss model.User) error {
	err := c.Userrepo.UpdateUser(ctx, id, paramss)
	return err

}

func (c *Userserv) DeleteUser(ctx context.Context, id string) error {
	err := c.Userrepo.DeleteUser(ctx, id)
	return err

}
