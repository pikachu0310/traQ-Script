package main

import (
	"context"
	"fmt"

	"github.com/traPtitech/go-traq"
)

const TOKEN = ""

func main() {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

	v, _, _ := client.ChannelApi.
		GetChannels(auth).
		IncludeDm(true).
		Execute()
	// v, _, _ := client.MeApi.GetMyUnreadChannels(auth).Execute()
	fmt.Printf("%#v", v)
}
