package main

import (
	"flag"
	"fmt"

	"github.com/apertureless/bs3stat/db"
	"github.com/apertureless/bs3stat/handlers"
	"github.com/apertureless/bs3stat/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	port := flag.String("port", "3000", "Your port number (default to 3000).")
	migrate := flag.Bool("migrate", false, "Run database migrations.")
	seed := flag.Bool("seed", false, "Seed database with some demo data.")
	flag.Parse()

	db := db.InitDB("storage.db", *migrate, *seed)
	model.SetDatabase(db)

	e := echo.New()
	e = routes(e)

	// Serve static files
	e.Static("/", "web/dist")
	e.File("/", "web/dist/index.html")

	// Start server
	fmt.Println("\n\x1b[32;1mRunning on: http://localhost:" + *port + ".\x1b[0m")
	defer db.Close()
	e.Run(standard.New(":" + *port))
}

func routes(e *echo.Echo) *echo.Echo {
	e.GET("/projects", handlers.ProjectIndex)
	e.POST("/projects", handlers.CreateProject)
	e.GET("/projects/:id", handlers.GetProject)
	e.PUT("/projects/:id", handlers.UpdateProject)
	e.DELETE("/projects/:id", handlers.DeleteProject)

	return e
}
