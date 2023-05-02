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
	fiveMinutesAgo := now.Add(time.Duration(-10) * time.Minute)
	requestData := map[string]any{"from": "a4f4ca7e-054e-4d8b-842e-8b777c353a5d", "after": fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z")}
	//requestData := map[string]any{"from": "f60166fb-c153-409a-811d-272426eda32b", "in": "5202813c-0063-46e0-9afa-401dc1bbb250", "after": fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z")}
	//requestData := map[string]any{"from": "f60166fb-c153-409a-811d-272426eda32b"}

	messages, err := api.GetMessages(requestData)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(messages.Hits))

	for {
		for _, message := range messages.Hits {
			count := 0
			for _, stamp := range message.Stamps {
				if stamp.StampId == "6308a443-69f0-45e5-866f-56cc2c93578f" {
					count++
				}
			}
			if count >= 1 {
				fmt.Println(message.Id, count)
			}
		}
		time.Sleep(time.Second * 100)
	}
	//
	//for {
	//	for i := 0; i <= 10; i++ {
	//		api.PostStamp("7fc2ff8b-f230-466a-91ec-24ae689d2382", "6308a443-69f0-45e5-866f-56cc2c93578f")
	//	}
	//	time.Sleep(time.Second * 1)
	//}
}
