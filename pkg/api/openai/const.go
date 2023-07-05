package openai

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var Client = openai.NewClient(GetApiKey())

func GetApiKey() (apiKey string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
	}
	apiKey = os.Getenv("APIKEY")
	return apiKey
}
