package main

import (
	"fmt"
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
	fiveMinutesAgo := now.Add(time.Duration(-5) * time.Minute)
	fmt.Println(fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z"))

	requestData := map[string]string{"from": "a4f4ca7e-054e-4d8b-842e-8b777c353a5d", "after": fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z")}

	messages, err := api.GetMessages(requestData)
	if err != nil {
		panic(err)
	}
	fmt.Println(*messages)

	for i, message := range messages.Hits {
		api.EditMessages(message.Id, fmt.Sprintf("%s (Auto Edited %d)", message.Content, i))
	}

	// fmt.Printf("%#v", body)
}
