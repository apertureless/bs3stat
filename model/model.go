package model

import "github.com/jinzhu/gorm"

// Model is the gorm Model structure
type Model struct{ DB *gorm.DB }

// DB Global Database connection
var DB *gorm.DB

// SetDatabase assignts global db with local
func SetDatabase(db *gorm.DB) {
	DB = db
}
