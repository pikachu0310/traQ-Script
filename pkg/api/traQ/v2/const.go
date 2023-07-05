package v2

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/traPtitech/go-traq"
)

var Client = traq.NewAPIClient(traq.NewConfiguration())
var Auth = context.WithValue(context.Background(), traq.ContextAccessToken, GetToken())

func GetToken() (token string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
	}
	token = os.Getenv("TOKEN")
	return token
}
