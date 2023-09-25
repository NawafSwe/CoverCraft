package http

import (
	"coverCraft/domain/aggregates"
	"coverCraft/domain/services"
	valueObjects "coverCraft/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"html"
	"net/http"
	"text/template"
)

func GenerateCoverLetter(c echo.Context) error {
	optimizationReq := new(valueObjects.GenerateCoverLetterRequest)
	if err := c.Bind(optimizationReq); err != nil {
		return err
	}
	// Sanitize text fields
	optimizationReq.Job.Company = html.EscapeString(optimizationReq.Job.Company)
	optimizationReq.Job.Location = html.EscapeString(optimizationReq.Job.Location)
	optimizationReq.Job.Name = html.EscapeString(optimizationReq.Job.Name)
	optimizationReq.Job.Description = html.EscapeString(optimizationReq.Job.Description)
	optimizationReq.Job.Qualifications = html.EscapeString(optimizationReq.Job.Qualifications)
	optimizationReq.Resume.Text = html.EscapeString(optimizationReq.Resume.Text)

	job := aggregates.Job{
		Id:             uuid.New(),
		Name:           optimizationReq.Job.Name,
		Location:       optimizationReq.Job.Location,
		Company:        optimizationReq.Job.Company,
		Description:    optimizationReq.Job.Description,
		Qualifications: optimizationReq.Job.Qualifications,
	}
	if err := aggregates.CreateJob(&job); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to generate a cover letter")
	}

	response, err := services.ProcessResume(*optimizationReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, valueObjects.GenerateCoverLetterResult{Result: response})
}

func RenderResumeForm(c echo.Context) error {
	tmpl, err := template.ParseFiles("templates/resume_form.html")
	if err != nil {
		return err
	}

	// Render the HTML template
	return tmpl.Execute(c.Response().Writer, nil)
}
