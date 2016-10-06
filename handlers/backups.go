package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetBackups Endpoint
func GetBackups(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "backups")
	}
}

// PutBackup Endpoint
func PutBackup(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}
