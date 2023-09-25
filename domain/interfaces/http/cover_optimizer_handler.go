package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"html"
	"net/http"
	"os"
	"text/template"
)

type job struct {
	Company        string `json:"company"`
	Location       string `json:"location"`
	Description    string `json:"description"`
	Qualifications string `json:"qualifications"`
	Name           string `json:"name"`
}

type resume struct {
	Text string `json:"text"`
}

type optimizeResumeRequest struct {
	Job    job    `json:"job"`
	Resume resume `json:"resume"`
}
type optimizationResult struct {
	Result string `json:"result"`
}

func GenerateCoverLetter(c echo.Context) error {
	u := new(optimizeResumeRequest)
	if err := c.Bind(u); err != nil {
		return err
	}
	// Sanitize text fields
	u.Job.Company = html.EscapeString(u.Job.Company)
	u.Job.Location = html.EscapeString(u.Job.Location)
	u.Job.Name = html.EscapeString(u.Job.Name)
	u.Job.Description = html.EscapeString(u.Job.Description)
	u.Job.Qualifications = html.EscapeString(u.Job.Qualifications)
	u.Resume.Text = html.EscapeString(u.Resume.Text)

	client := openai.NewClient(os.Getenv("CHAT_GPT_TOKEN"))
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Hello, I'd like to get a help, I need you to write a cover letter based on a job. The job is at %v, and here's the description of the job: %v. The qualifications for this position are: %v. Could you please optimize the following resume by paraphrase all sections to be optimized for this job? Here's the resume in text: %v", u.Job.Company, u.Job.Description, u.Job.Qualifications, u.Resume.Text),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, optimizationResult{Result: response.Choices[0].Message.Content})
}

func RenderResumeForm(c echo.Context) error {
	tmpl, err := template.ParseFiles("templates/resume_form.html")
	if err != nil {
		return err
	}

	// Render the HTML template
	return tmpl.Execute(c.Response().Writer, nil)
}
