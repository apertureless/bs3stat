package handlers

import (
	"bs3stat/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetBackups Endpoint
func GetBackups(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetBackups(db))
	}
}

// PutBackup Endpoint
func PutBackup(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create instance of new backup
		var backup models.Backup
		// Map incoming json body to backup
		c.Bind(&backup)

		id, err := models.PutBackup(db, backup.Name, backup.Started, backup.Finished, backup.Duration, backup.Status)
		// Return JSON response if successfull

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}
