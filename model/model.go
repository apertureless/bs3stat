package model

import "github.com/jinzhu/gorm"

// Model is the gorm Model structure
type Model struct{ DB *gorm.DB }
