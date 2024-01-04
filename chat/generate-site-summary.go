package chat

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"time"
)

func GenerateSummary(content string) string {
	instruction := "Your task: create a concise summary of the site's content. you need to highlight only the main gist of the content. if required, get the links from this tutorial then cut it down"
	time.Sleep(15 * time.Second)
	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: instruction,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		fmt.Printf("Full resp: %v\n", resp)
		return "\"ChatCompletion error: %v\\n\", err"
	}

	return resp.Choices[0].Message.Content
}
