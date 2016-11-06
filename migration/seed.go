package migration

import (
	"bs3stat/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	project := model.Project{
		Title: "Hello World",
		Name:  "acme",
	}

	db.Create(&project)

	fmt.Println("\x1b[32;1mDatabase seeding completed.\x1b[0m")
}
