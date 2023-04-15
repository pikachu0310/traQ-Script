package main

import (
	"os"
	"time"

	"traQ-Script/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	name := "pikachu"
	password := os.Getenv("PASSWORD")
	_, err = api.Login(name, password)
	if err != nil {
		panic(err)
	}
	if api.CookieCache == "" {
		panic("CookieCache is empty")
	}

	now := time.Now()
	fiveMinutesAgo := now.Add(time.Duration(-30) * time.Minute)
	requestData := map[string]any{"from": "f60166fb-c153-409a-811d-272426eda32b", "in": "5202813c-0063-46e0-9afa-401dc1bbb250", "after": fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z")}

	messages, err := api.GetMessages(requestData)
	if err != nil {
		panic(err)
	}

	for {
		for _, message := range messages.Hits {
			api.PostStamp(message.Id, "4a7315c0-c8d8-4a65-b6f4-5866a29bd113")
		}
		time.Sleep(time.Second * 1)
	}
}
