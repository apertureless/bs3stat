package migration

import (
	"fmt"

	"github.com/apertureless/bs3stat/model"

	"github.com/jinzhu/gorm"
)

const (
	driverURL     string = "sqlite3://storage.db"
	storageFile   string = "./storage.db"
	migrationPath string = "./db/migrations/"
)

// Migrate Auto migrate schema, declared in our models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Project{})

	fmt.Println("\x1b[32;1mDatabase migration completed.\x1b[0m")
}
