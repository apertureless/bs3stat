package db

import (
	"fmt"

	"github.com/apertureless/bs3stat/migration"

	"github.com/jinzhu/gorm"
	// Needs to be blank import because of sqlite init function
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	return db
}
