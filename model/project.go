package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Project represends a backup project
type Project struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Name  string `gorm:"size:255"`
}

// AddProject is the action to add a new project to the database
func AddProject() {
	fmt.Println("yo")
}

// AllProjects return a collection of Project
func AllProjects() []Project {
	var projects []Project
	DB.Find(&projects)

	return projects
}

// GetProject by ID
func GetProject(id string) Project {
	var project Project
	DB.First(&project, id)
	return project
}
