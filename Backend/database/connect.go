package database

import (
	// "example/Project3/internal/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Connect interface{
// 	GetDB() *gorm.DB
// }

// type Connectdb struct{

// }

// func ConnectFunc() Connect{
// 	return &Connectdb{}
// }

type PostgresConnectionConfig struct {
	DSN string
	// MaxIdleConnections int
	// MaxOpenConnections int
}

var db *gorm.DB

func InitPostgres(config *PostgresConnectionConfig) (err error) {
	db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("unable to connect to postgres:%w", err))

	}
	fmt.Println("inside Initpostgrse")

	fmt.Println(db)
	// db.AutoMigrate(&model.Document{})
	// db.AutoMigrate(&model.User{})
	fmt.Println("inside database")
	// sqlDB,err:=db.DB()
	// if err!=nil{
	// 	panic(fmt.Errorf("unable to create an underlying database instance:%w",err))
	// }
	// if config.MaxIdleConnections>0{
	// 	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	// }
	// if config.MaxOpenConnections>0{
	// 	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	// }
	return err
}

//	func (c *Connectdb) GetDB() *gorm.DB{
//		return db
//	}
func GetDB() *gorm.DB {
	fmt.Println("inside GetDB")
	fmt.Println(db)
	return db
}
