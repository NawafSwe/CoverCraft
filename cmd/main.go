package main

import (
	"coverCraft/config"
	coverOptimizerRoutes "coverCraft/domain/interfaces/http"
	"coverCraft/domain/repositories"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type App struct {
	// Add your dependencies here
	JobRepo repositories.JobRepository
}

func NewApp(jobRepo repositories.JobRepository) *App {
	return &App{
		JobRepo: jobRepo,
	}
}
func SetupRouter(app *App) *echo.Echo {
	e := echo.New()

	// Define your routes and handlers here, and use app.JobRepo as needed
	e.GET("/jobs", func(c echo.Context) error {
		return coverOptimizerRoutes.ListAllJobs(c, app.JobRepo)
	})
	// Routes
	e.GET("/health", health)

	e.POST("/generate-cover-letter", func(c echo.Context) error {
		return coverOptimizerRoutes.GenerateCoverLetter(c, app.JobRepo)
	})
	e.GET("/generate-cover-letter", coverOptimizerRoutes.RenderResumeForm)
	// default redirect
	e.GET("/", func(c echo.Context) error {

		return c.Redirect(http.StatusSeeOther, "/generate-cover-letter")
	})

	return e
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// initialize db connection
	config.DatabaseInit()
	gorm := config.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	pingErr := dbGorm.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	// create app with dependencies
	app := NewApp(repositories.NewGormJobRepository(gorm))
	// Set up the Echo router and pass the app instance
	e := SetupRouter(app)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
		Status  bool   `json:"status"`
	}{
		Message: "Server Alive!",
		Status:  true,
	})
}
