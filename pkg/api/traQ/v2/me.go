package v2

import (
	"net/http"

	"github.com/traPtitech/go-traq"
)

func GetMyUnreadChannels(client int) ([]traq.UnreadChannel, *http.Response, error) {
	return GetClient(client).MeApi.GetMyUnreadChannels(GetContext(client)).Execute()
}

func ReadChannel(client int, channelId string) (*http.Response, error) {
	return GetClient(client).MeApi.ReadChannel(GetContext(client), channelId).Execute()
}
