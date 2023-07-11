package v2

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

const (
	Bot = iota
	Me
)

func GetClient(clientType int) (client *traq.APIClient) {
	switch clientType {
	case Bot:
		return BotClient
	case
		Me:
		return MeClient
	default:
		fmt.Printf("error: GetClient関数の引数が不正です。")
		return nil
	}
}

func GetContext(clientType int) (c context.Context) {
	switch clientType {
	case Bot:
		return context.Background()
	case
		Me:
		return Auth
	default:
		fmt.Printf("error: GetContext関数の引数が不正です。")
		return nil
	}
}

var MeClient = traq.NewAPIClient(traq.NewConfiguration())
var BotClient = GetBot().API()
var Auth = context.WithValue(context.Background(), traq.ContextAccessToken, GetToken())

func GetToken() (token string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
	}
	token = os.Getenv("TOKEN")
	return token
}

func GetBot() (bot *traqwsbot.Bot) {
	token := GetBotToken()

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: token,
	})
	if err != nil {
		fmt.Printf("error: Bot変数が作れなかった!: %v", err)
	}
	return bot
}

func GetBotToken() (token string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
	}
	token = os.Getenv("BOTTOKEN")
	return token
}
