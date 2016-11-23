package handlers

import (
	"bs3stat/model"
	"net/http"

	"github.com/labstack/echo"
)

// H maps strings
type H map[string]interface{}

// ProjectIndex returns all Projects in the database
func ProjectIndex(c echo.Context) error {
	p := model.AllProjects()
	return c.JSON(http.StatusOK, p)
}

// CreateProject handles the creation of new projects
func CreateProject(c echo.Context) error {
	var project model.Project
	c.Bind(&project)

	err := model.AddProject(project)

	if err == nil {
		return c.JSON(http.StatusCreated, H{
			"created": "😊",
		})

	}

	return c.JSON(http.StatusOK, err)

}

// GetProject gets a project by ID
func GetProject(c echo.Context) error {
	id := c.Param("id")
	p := model.GetProject(id)
	return c.JSON(http.StatusOK, p)
}

// UpdateProject updates an project
func UpdateProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// DeleteProject deletes an project by ID
func DeleteProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
