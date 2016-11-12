package main

import (
	"bs3stat/db"
	"bs3stat/handlers"
	"bs3stat/model"
	"flag"
	"fmt"

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

	project := model.Project{
		Title: "Lalalalla dongo",
		Name:  "acme_dongo",
	}
	db.Create(&project)

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
