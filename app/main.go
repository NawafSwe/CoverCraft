package main

import (
	"coverCraft/config"
	coverOptimizerRoutes "coverCraft/cover_optimizer/http"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	// Routes
	e.GET("/health", health)
	e.POST("/generate-cover-letter", coverOptimizerRoutes.GenerateCoverLetter)
	e.GET("/generate-cover-letter", coverOptimizerRoutes.RenderResumeForm)
	// default redirect
	e.GET("/", coverOptimizerRoutes.RenderResumeForm)
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
