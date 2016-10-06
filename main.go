package main

import (
	"bs3stat/handlers"
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Setup Database
	db := initDB("storage.db")
	migrate(db)
	// Echo Framework instance
	e := echo.New()

	// Serve index file
	e.File("/", "public/index.html")
	// Routes
	e.GET("/backups", handlers.GetBackups(db))
	e.PUT("/backups", handlers.PutBackup(db))

	// Start Webserver
	e.Run(standard.New(":8080"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Check for DB errors
	if err != nil {
		panic(err)
	}

	// Check if no db erros, but still no connection
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS backups(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		project VARCHAR NOT NULL,
		started DATETIME NOT NULL,
		finished DATETIME NOT NULL,
		status VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
