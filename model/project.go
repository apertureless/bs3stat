package model

import (
	"github.com/jinzhu/gorm"
)

// Project represends a backup project
type Project struct {
	gorm.Model
	Title    string `gorm:"size:255"`
	Name     string `gorm:"size:255"`
	Status   string `gorm:"size:255"`
	Duration string `gorm:"size:255"`
}

// AddProject is the action to add a new project to the database
func AddProject(project Project) (err error) {

	// Start transaction
	tx := DB.Begin()

	// Rollback if error occurs
	if err := tx.Create(&project).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	tx.Commit()

	return nil

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

// DeleteProject deletes an project based on the project id
func DeleteProject(id string) {
	var project Project
	DB.First(&project, id)
	DB.Delete(&project)
}

// UpdateProject can change project data
func UpdateProject(project Project) {

}
