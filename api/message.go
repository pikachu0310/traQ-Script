package api

import (
	"fmt"
	"time"
)

const getMessagesURL = BaseURL + "/messages"

type GetMessagesResponse struct {
	TotalHits int `json:"totalHits"`
	Hits      []struct {
		Id        string    `json:"id"`
		UserId    string    `json:"userId"`
		ChannelId string    `json:"channelId"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Pinned    bool      `json:"pinned"`
		Stamps    []struct {
			UserId    string    `json:"userId"`
			StampId   string    `json:"stampId"`
			Count     int       `json:"count"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"stamps"`
		ThreadId string `json:"threadId"`
	} `json:"hits"`
}

func GetMessages() (*GetMessagesResponse, error) {
	requestData := map[string]string{"word": "pikachu"}
	var getMeResponse *GetMessagesResponse
	err := RequestAndGetResponse("GET", getMessagesURL, requestData, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return &GetMessagesResponse{}, err
	}
	return getMeResponse, nil
}
