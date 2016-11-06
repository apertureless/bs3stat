package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ProjectIndex(c echo.Context) error {
	// u := new(User)
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	//
	// return c.JSON(http.StatusCreated, u)
	return c.NoContent(http.StatusNoContent)
}

func CreateProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func GetProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func UpdateProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func DeleteProject(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
