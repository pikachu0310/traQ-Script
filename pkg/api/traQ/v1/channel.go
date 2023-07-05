package v1

import (
	"fmt"
)

const channelsURL = BaseURL + "/channels"

type GetChannelsResponse struct {
	Public []GetChannelResponsePublic `json:"public"`
	Dm     []struct {
		Id     string `json:"id"`
		UserId string `json:"userId"`
	} `json:"dm"`
}

type GetChannelResponsePublic struct {
	Id       string   `json:"id"`
	ParentId string   `json:"parentId"`
	Archived bool     `json:"archived"`
	Force    bool     `json:"force"`
	Topic    string   `json:"topic"`
	Name     string   `json:"name"`
	Children []string `json:"children"`
}

func GetChannels() (*GetChannelsResponse, error) {
	var getChannelsResponse *GetChannelsResponse
	err := RequestAndGetResponse("GET", channelsURL, nil, &getChannelsResponse)
	if err != nil {
		fmt.Println(err)
		return getChannelsResponse, err
	}
	return getChannelsResponse, nil
}
