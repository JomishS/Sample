package repository

import (
	"example/Project3/internal/model"
	"fmt"
	"time"

	// "log"
	// "strconv"
	// "errors"
	"context"

	// "github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	"example/Project3/database"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Userrepository interface {
	CreateUser(c context.Context, params model.User) (model.User, error)
	GetByIdUser(ctx context.Context, id string) (model.User, error)
	GetAllUser(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id string, paramss model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type Userrepo struct {
	// Datab database.Connect
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
	err = c.Db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (c *Userrepo) GetAllUser(ctx context.Context) ([]model.User, error) {

	var err error
	users := []model.User{}

	err = c.Db.Where("deletedat IS NULL").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, err
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
	// err := databseconn.Getdb().Where("id=?", id).Update("plant",paramss.Plant).Error
	// fmt.Println("repository")
	if err := c.Db.Save(&update).Error; err != nil {
		return err
	}
	return nil
}
func (c *Userrepo) DeleteUser(ctx context.Context, id string) error {
	// var user model.User
	// err := c.Db.Delete(&user, id).Error
	// return err
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
