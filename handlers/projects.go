package handlers

import (
	"bs3stat/model"
	"net/http"

	"github.com/labstack/echo"
)

// ProjectIndex returns all Projects in the database
func ProjectIndex(c echo.Context) error {
	p := model.AllProjects()
	return c.JSON(http.StatusOK, p)
}

// CreateProject handles the creation of new projects
func CreateProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// GetProject gets a project by ID
func GetProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// UpdateProject updates an project
func UpdateProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// DeleteProject deletes an project by ID
func DeleteProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
