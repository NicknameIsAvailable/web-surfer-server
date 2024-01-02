package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	searcher "github.com/serpapi/google-search-results-golang"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Print(find("Как сделать кофе"))
}

func find(query string) any {
	apiKey := os.Getenv("SEARCH_API")

	fmt.Printf("apikey: %v\n", apiKey)

	parameter := map[string]string{
		"q":             query,
		"location":      "Austin, Texas, United States",
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
		"api_key":       apiKey,
	}

	fmt.Print(parameter)

	search := searcher.NewGoogleSearch(parameter, apiKey)
	results, err := search.GetJSON()

	if err != nil {
		fmt.Print(err)
	}

	organic_results := results["organic_results"].([]interface{})
	return organic_results
}

func generate(content string, useWeb bool) string {
	apiKey := os.Getenv("OPENAI_API")
	if apiKey == "" {
		log.Fatal("OPENAI_API is empty. Make sure it is set in your .env file.")
	}
	fmt.Printf("apikey: %v\n", apiKey)
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
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
