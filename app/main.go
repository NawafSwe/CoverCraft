package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", health)

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
