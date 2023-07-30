package v2

import (
	"net/http"

	"github.com/traPtitech/go-traq"
)

var ChannelIdToChannelName, ChannelNameToChannelId, ChannelIdToChannelParentId = ChannelCache()

func ChannelCache() (ChannelIdToChannelName, ChannelNameToChannelId, ChannelIdToChannelParentId map[string]string) {
	Channel, _, err := GetChannels()
	if err != nil {
		panic(err)
	}

	ChannelIdToChannelName = map[string]string{}
	for _, channel := range Channel.Public {
		ChannelIdToChannelName[channel.Id] = channel.Name
	}
	ChannelNameToChannelId = map[string]string{}
	for _, channel := range Channel.Public {
		ChannelNameToChannelId[channel.Name] = channel.Id
	}
	ChannelIdToChannelParentId = map[string]string{}
	for _, channel := range Channel.Public {
		parentId := channel.GetParentId()
		if parentId == "" {
			continue
		}
		ChannelIdToChannelParentId[channel.Id] = parentId
	}
	return
}

func ChannelIdToChannelNameFunc(ChannelId string) string {
	return ChannelIdToChannelName[ChannelId]
}

func ChannelNameToChannelIdFunc(ChannelName string) string {
	return ChannelNameToChannelId[ChannelName]
}

func ChannelIdToChannelParentIdFunc(ChannelId string) (string, bool) {
	channelId, ok := ChannelIdToChannelParentId[ChannelId]
	return channelId, ok
}

func ChannelIdToAllParentChannelName(ChannelId string) string {
	var ChannelName string
	for {
		channelId, ok := ChannelIdToChannelParentIdFunc(ChannelId)
		if !ok {
			ChannelName = ChannelIdToChannelNameFunc(ChannelId) + ChannelName
			break
		}
		ChannelName = "/" + ChannelIdToChannelNameFunc(ChannelId) + ChannelName
		ChannelId = channelId
	}
	return ChannelName
}

func GetChannel(channelId string) (*traq.Channel, *http.Response, error) {
	return GetClient(Bot).ChannelApi.GetChannel(GetContext(Bot), channelId).Execute()
}

func GetChannels() (*traq.ChannelList, *http.Response, error) {
	return GetClient(Bot).ChannelApi.GetChannels(GetContext(Bot)).Execute()
}
