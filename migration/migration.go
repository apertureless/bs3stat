package migration

import (
	"bs3stat/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	driverUrl     string = "sqlite3://storage.db"
	storageFile   string = "./storage.db"
	migrationPath string = "./db/migrations/"
)

// Auto migrate schema, declared in our models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Project{})

	fmt.Println("\x1b[32;1mDatabase migration completed.\x1b[0m")
}
