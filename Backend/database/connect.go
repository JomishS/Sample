package database

import (

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type PostgresConnectionConfig struct {
	DSN string
}

var db *gorm.DB

func InitPostgres(config *PostgresConnectionConfig) (err error) {
	db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("unable to connect to postgres:%w", err))

	}
	fmt.Println("inside Initpostgrse")

	fmt.Println(db)
	fmt.Println("inside database")
	return err
}

func GetDB() *gorm.DB {
	fmt.Println("inside GetDB")
	fmt.Println(db)
	return db
}
