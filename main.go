package main

import (
	"fmt"
	"time"

	"traQ-Script/api"
	"traQ-Script/util"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	_, _, err = util.Login()
	if err != nil {
		panic(err)
	}

	toUserID := ""
	fmt.Println("検索するユーザーのIDを入力してください(例:a4f4ca7e-054e-4d8b-842e-8b777c353a5d)")
	_, err = fmt.Scanln(&toUserID)
	if err != nil {
		panic(err)
	}

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
	// for {
	//	for i := 0; i <= 10; i++ {
	//		api.PostStamp("7fc2ff8b-f230-466a-91ec-24ae689d2382", "6308a443-69f0-45e5-866f-56cc2c93578f")
	//	}
	//	time.Sleep(time.Second * 1)
	// }
}
