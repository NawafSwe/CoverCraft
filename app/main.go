package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"os"
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

	// Routes
	e.GET("/health", health)
	e.POST("/optimize-resume", optimizeResumeFromJob)

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

func optimizeResumeFromJob(c echo.Context) error {
	u := new(optimizeResumeRequest)
	if err := c.Bind(u); err != nil {
		return err
	}
	client := openai.NewClient(os.Getenv("CHAT_GPT_TOKEN"))
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that optimizes resumes.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Hello, I'd like to share a job opportunity with you. The job is at %v, and here's the description: %v. The qualifications for this position are: %v. Could you please optimize the following resume for this job? Here's the resume in text: %v", u.Job.Company, u.Job.Description, u.Job.Qualifications, u.Resume.Text),
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
