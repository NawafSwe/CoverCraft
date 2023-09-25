package http

import (
	"coverCraft/domain/aggregates"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListAllJobs(c echo.Context) error {
	jobs, err := aggregates.ListAllJobs()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, jobs)
}
