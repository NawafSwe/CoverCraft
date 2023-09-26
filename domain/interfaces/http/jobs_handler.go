package http

import (
	"coverCraft/config"
	"coverCraft/domain/aggregates"
	"coverCraft/domain/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ListAllJobs , this is an endpoint to view posted jobs by people
func ListAllJobs(c echo.Context) error {
	jobRepo := repositories.NewGormJobRepository(config.DB())
	jobs, err := aggregates.ListAllJobs(jobRepo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, jobs)
}
