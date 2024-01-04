package chat

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Generate(chat []openai.ChatCompletionMessage, content string, role openai.ThreadMessageRole) string {
	message := openai.ChatCompletionMessage{
		Role:    string(role),
		Content: content,
	}
	var messages []openai.ChatCompletionMessage
	messages = append(chat, message)

	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		fmt.Printf("Full resp: %v\n", resp)
		return "\"ChatCompletion error: %v\\n\", err"
	}

	return resp.Choices[0].Message.Content
}
