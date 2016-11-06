package db

import (
	"bs3stat/migration"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// DBCon is the connection handle for the database
	DB *gorm.DB
)

func InitDB(filepath string, runMigrations bool, seedDatabase bool) *gorm.DB {
	var err error
	if DB, err = gorm.Open("sqlite3", filepath); err != nil {
		panic(fmt.Sprintf("failed to connect database. err=%+v", err))
	}

	defer DB.Close()

	if runMigrations == true {
		migration.Migrate(DB)
	}

	if seedDatabase == true {
		migration.Seed(DB)
	}

	return DB
}
