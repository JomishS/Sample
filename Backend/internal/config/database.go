package config

import (
	"example/Project3/database"
	"fmt"
	// "net/url"
	// "gorm.io/gorm"
)

type DatabaseConfig struct {
	Postgres  PostgresConfig  `mapstructure:"postgres"`
	Migration MigrationConfig `mapstructure:"migration"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type MigrationConfig struct {
	Path string `mapstructure:"path"`
}

func (c *PostgresConfig) GetDSN() string {

	// encodedPassword := url.QueryEscape(c.Password)
	fmt.Println("inside getdsn")
	fmt.Println(c.Host)
	fmt.Println(c.Port)
	fmt.Println(c.Password)
	fmt.Println(c.DBName)

	dsn := fmt.Sprintf(
		// "postgres://%s:%s:%s/%s",
		// c.User,c.Host,c.Port,c.DBName,
		"postgres://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.DBName,
	)
	return dsn
}

func (c *PostgresConfig) ToPostgresConnectionConfig() *database.PostgresConnectionConfig {
	return &database.PostgresConnectionConfig{
		DSN: c.GetDSN(),
	}

}

// var db *gorm.DB

// func InitPostgres() {
// 	config := GetPostgresConnectionConfig()
// 	db = database.InitPostgres(config)
// 	fmt.Println("connection established")
// }

// func GetPostgresConnectionConfig() *database.PostgresConnectionConfig {
// 	return &database.PostgresConnectionConfig{
// 		DSN: ViperConfig.GetString("POSTGRES_DSN"),
// 	}
// }

// func GetDB() *gorm.DB{
// 	return db
// }
