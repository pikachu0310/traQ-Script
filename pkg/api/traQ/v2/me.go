package v2

import (
	"net/http"

	"github.com/traPtitech/go-traq"
)

func GetMyUnreadChannels() ([]traq.UnreadChannel, *http.Response, error) {
	return Client.MeApi.GetMyUnreadChannels(Auth).Execute()
}

func ReadChannel(channelId string) (*http.Response, error) {
	return Client.MeApi.ReadChannel(Auth, channelId).Execute()
}
