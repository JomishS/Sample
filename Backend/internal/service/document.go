package Service

import (

	"context"
	"example/Project3/internal/dto"
	"example/Project3/internal/model"
	"example/Project3/internal/repository"
	"example/Project3/internal/util"
	"fmt"
	_ "github.com/lib/pq"
)

type Documentservice interface {
	CreateDoc(c context.Context, params *model.Document) (*model.Document, error)
	UpdateDoc(c context.Context, id string, params model.Document) (model.Document, error)
	DeleteDoc(c context.Context, id string) ( error)
	GetByIdDoc(c context.Context, id string) (model.Document, error)
	GetAllDoc(c context.Context,query *dto.AssetQueryParams) ([]model.Document, error)
}

type Docserv struct {
	Docrepo repository.Documentrepository
}

func DocumentServFunc(Docrepo repository.Documentrepository) Documentservice {
	return &Docserv{
		Docrepo: Docrepo,
	}
}
func (c *Docserv) CreateDoc(context context.Context, params *model.Document) (*model.Document, error) {
	crtdoc, err := c.Docrepo.CreateDoc(context, params)
	fmt.Println("inside service")
	return crtdoc, err
}

func (c *Docserv) GetByIdDoc(ctx context.Context, id string) (model.Document, error) {
	var err error

	plant, err := c.Docrepo.GetByIdDoc(ctx, id)
	return plant, err
}

func (c *Docserv) GetAllDoc(ctx context.Context,query *dto.AssetQueryParams) ([]model.Document, error) {
	var err error
	filterMap, err := util.ParseFilters(query.Filter)
	if err != nil {
		return nil, err
	}
	doc, _,err := c.Docrepo.GetAllDoc(ctx,query,filterMap)
	return doc, err
}

func (c *Docserv) UpdateDoc(ctx context.Context, id string, paramss model.Document) (model.Document, error) {
	todo,err := c.Docrepo.UpdateDoc(ctx, id, paramss)
	return  todo,err

}

func (c *Docserv) DeleteDoc(ctx context.Context, id string) (  error) {
	 err := c.Docrepo.DeleteDoc(ctx,id)
	return  err

}
