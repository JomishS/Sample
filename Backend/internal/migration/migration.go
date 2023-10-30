package migration

import (
	// "database/sql"
	"errors"
	"example/Project3/internal/config"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

)

type Migrate struct {
}

func (m *Migrate) Up(args []string) error {
	var err error
	var steps int
	
	if len(args) == 1 {
		steps, err = strconv.Atoi(args[0])
		if err != nil || steps < 1 {
			return errors.New("argument must be an positive integer")
		}
	}
	fmt.Println(steps)
	migrationPath := config.ViperConfig.GetString("POSTGRES_MIGRATION_PATH")
	conf := config.GetConfig()
	DSN := conf.Database.Postgres.GetDSN()
	newMigrate, err := migrate.New(migrationPath, DSN)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		err = newMigrate.Steps(steps)
	} else {
		err = newMigrate.Up()
	}

	if err == nil {
		return nil
	} else if strings.Contains(err.Error(), "no change") {
		fmt.Println("No pending changes")
		return nil
	}
	return err
}

func (m *Migrate) Down(args []string) error {
	var err error
	var steps int

	if len(args) == 1 {
		steps, err = strconv.Atoi(args[0])
		if err != nil || steps < 1 {
			return errors.New("argument must be an positive integer")
		}
	}

	migrationPath := config.ViperConfig.GetString("POSTGRES_MIGRATION_PATH")
	conf := config.GetConfig()
	DSN := conf.Database.Postgres.GetDSN()
	fmt.Println(migrationPath)
	fmt.Println(DSN)
	newMigrate, err := migrate.New(migrationPath, DSN)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		err = newMigrate.Steps(-1 * steps)
	} else {
		err = newMigrate.Down()
	}

	if err == nil {
		return nil
	} else if strings.Contains(err.Error(), "no change") {
		fmt.Println("No pending changes")
		return nil
	}
	return err
}
