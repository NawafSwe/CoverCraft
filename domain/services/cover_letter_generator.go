package services

import (
	"context"
	"coverCraft/domain/valueobjects"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func ProcessResume(op valueobjects.GenerateCoverLetterRequest) (string, error) {

	client := openai.NewClient(os.Getenv("CHAT_GPT_TOKEN"))
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Hello, I'd like to get a help, I need you to write a cover letter based on a job. The job is at %v, and here's the description of the job: %v. The qualifications for this position are: %v. Could you please optimize the following resume by paraphrase all sections to be optimized for this job? Here's the resume in text: %v", op.Job.Company, op.Job.Description, op.Job.Qualifications, op.Resume.Text),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}
