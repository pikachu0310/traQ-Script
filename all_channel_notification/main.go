package main

import (
	"fmt"

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

	// toUserID := ""
	// fmt.Println("検索するユーザーのIDを入力してください(例:a4f4ca7e-054e-4d8b-842e-8b777c353a5d)")
	// _, err = fmt.Scanln(&toUserID)
	// if err != nil {
	// 	panic(err)
	// }

	channels, err := api.GetChannels()
	if err != nil {
		panic(err)
	}

	for _, channel := range channels.Public {
		fmt.Println(channel.Id, channel.Name)
	}
}
