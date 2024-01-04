package chat

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var (
	APIKey string
	Client *openai.Client
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	APIKey = os.Getenv("OPENAI_API")
	Client = openai.NewClient(APIKey)
}
