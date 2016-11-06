package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Name  string `gorm:"size:255"`
}

func AddProject() {
	fmt.Println("yo")
}

func (this *Model) AllProjects() []Project {
	var projects []Project
	this.DB.First(&projects)
	// this.First(&projects)
	fmt.Println(projects)

	return projects
}
