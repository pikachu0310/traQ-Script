package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"traQ-Script/api"

	"github.com/joho/godotenv"
	"golang.org/x/term"
)

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	fmt.Printf("traQに接続して検索するためにログインが必要です。\n")
	username, password, _ := credentials()
	//fmt.Printf("Username: %s, Password: %s\n", username, "********")
	_, err = api.Login(username, password)
	if err != nil {
		panic(err)
	}
	if api.CookieCache == "" {
		panic("CookieCache is empty(usernameかpasswordが間違えています)")
	}

	toUserID := ""
	fmt.Println("検索するユーザーのIDを入力してください(例:a4f4ca7e-054e-4d8b-842e-8b777c353a5d)")
	_, err = fmt.Scanln(&toUserID)
	if err != nil {
		panic(err)
	}

	//now := time.Now()
	//fiveMinutesAgo := now.Add(time.Duration(-10) * time.Minute)
	//requestData := map[string]any{"from": "f60166fb-c153-409a-811d-272426eda32b", "in": "5202813c-0063-46e0-9afa-401dc1bbb250", "after": fiveMinutesAgo.UTC().Format("2006-01-02T15:04:05Z")}
	//requestData := map[string]any{"from": "f60166fb-c153-409a-811d-272426eda32b"}

	offset := 0
	for {
		requestData := map[string]any{"from": toUserID, "limit": 100, "offset": offset}
		messages, err := api.GetMessages(requestData)
		if err != nil {
			panic(err)
		}
		fmt.Println(offset)
		if len(messages.Hits) == 0 {
			break
		}

		for _, message := range messages.Hits {
			count := 0
			for _, stamp := range message.Stamps {
				if stamp.StampId == api.Stamp_Kusa {
					count++
				}
			}
			if count >= 1 {
				fmt.Println(message.Id, count)
			}
		}

		time.Sleep(time.Second * 1)
		offset += 100
	}
	//
	//for {
	//	for i := 0; i <= 10; i++ {
	//		api.PostStamp("7fc2ff8b-f230-466a-91ec-24ae689d2382", "6308a443-69f0-45e5-866f-56cc2c93578f")
	//	}
	//	time.Sleep(time.Second * 1)
	//}
}
