package config

import (
	"example/Project3/database"
	"fmt"
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

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DBName,
	)
	return dsn
}

func (c *PostgresConfig) ToPostgresConnectionConfig() *database.PostgresConnectionConfig {
	return &database.PostgresConnectionConfig{
		DSN: c.GetDSN(),
	}

}

