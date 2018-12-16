package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/koreset/timeslice/models"
	"github.com/pkg/errors"
)

var (
	db                        *gorm.DB
	err                       error
	ErrDatabaseInitialization = errors.New("There was an unspecified error initializing the persistence store")
)

// initializeDB sets up a database instance for the application
// Currently we are setting up for SQLite only but we will need to accommodate other
// dbs based on config paramaters
// returns a configured gorm.DB object
func InitializeDB() error {
	db, err = gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/timeslice?parseTime=True&charset=utf8&loc=Local")

	if err != nil {
		return err
	}
	db.LogMode(true)

	//At some point we have to automigrate database models. That will happen here
	db.AutoMigrate(&models.Employee{}, &models.Client{}, &models.Project{}, &models.TimesheetEntry{})
	return nil
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	_ = InitializeDB()

	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
