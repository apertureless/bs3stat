package db

import (
	"bs3stat/migration"
	"fmt"

	"github.com/jinzhu/gorm"
	// Needs to be blank import because of sqlite init function
	_ "github.com/mattn/go-sqlite3"
)

// InitDB creates an database and opens a connection
func InitDB(filepath string, runMigrations bool, seedDatabase bool) *gorm.DB {
	var err error
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database. err=%+v", err))
	}

	if runMigrations == true {
		migration.Migrate(db)
	}

	if seedDatabase == true {
		migration.Seed(db)
	}

	db.LogMode(true)
	return db
}
