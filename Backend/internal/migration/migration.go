package migration

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
)

type DBMigrate struct {
	Path string
	DSN  string
}

func NewMigrate(dsn string, migrationpath string) *DBMigrate {
	return &DBMigrate{
		Path: migrationpath,
		DSN:  dsn,
	}
}

func (m *DBMigrate) GetPath() (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	migrationFilePath := "file://" + mydir + m.Path
	return migrationFilePath, nil
}

func (m *DBMigrate) Up(args []string) error {
	var err error
	var steps int
	steps, err = strconv.Atoi(args[0])
	if err != nil || steps < 1 {
		return errors.New("argument must be an positive integer")
	}

	mydir, err := m.GetPath()
	if err != nil {
		return err
	}

	migrate, err := migrate.New(mydir, m.DSN)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		err = migrate.Steps(steps)
	} else {
		err = migrate.Up()
	}
	
	if err == nil {
		fmt.Println("Database migrations are successfull")
	} else if strings.Contains(err.Error(), "no change") {
		fmt.Println("No pending changes to apply on database migrations")
		return nil
	}
	return err
}

func (m *DBMigrate) Down(args []string) error {
	var err error
	var steps int
	steps, err = strconv.Atoi(args[0])
	if err!=nil || steps < 1{
		return errors.New("the argument must be a positive integer greater than 1")
	}
	mydir, err := m.GetPath()
	if err != nil {
		return err
	}
	migrate, err := migrate.New(mydir, m.DSN)
	if err != nil {
		return err
	}
	if len(args)==1{
		err=migrate.Steps(-1 * steps)
	} else{
		err=migrate.Down()
	}
	if err==nil{
		return nil
	} else if strings.Contains(err.Error(),"no change"){
		fmt.Println("no pending changes")
	}
	return err
	// migrate.Steps(1)
	// return migrate.Down()
}
