package migration

import (
	"fmt"

	"github.com/apertureless/bs3stat/model"

	"github.com/jinzhu/gorm"
)

// Seed creates dummy data
func Seed(db *gorm.DB) {
	project := model.Project{
		Title: "Hello World",
		Name:  "acme",
	}

	db.Create(&project)

	fmt.Println("\x1b[32;1mDatabase seeding completed.\x1b[0m")
}
