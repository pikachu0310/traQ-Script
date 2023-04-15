package api

import (
	"fmt"
	"time"
)

const getMessagesURL = BaseURL + "/messages"
const MessagesURL = BaseURL + "/messages/"

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

func GetMessages(requestData map[string]any) (*GetMessagesResponse, error) {
	// requestData := map[string]string{"from": "a4f4ca7e-054e-4d8b-842e-8b777c353a5d"}
	var getMeResponse *GetMessagesResponse
	err := RequestAndGetResponse("GET", getMessagesURL, requestData, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return &GetMessagesResponse{}, err
	}
	return getMeResponse, nil
}

func EditMessages(messageID string, content string) (*GetMessagesResponse, error) {
	requestData := map[string]any{"content": content}
	var getMeResponse *GetMessagesResponse
	err := RequestAndGetResponse("PUT", MessagesURL+messageID, requestData, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return &GetMessagesResponse{}, err
	}
	return getMeResponse, nil
}

func PostStamp(messageID string, stampID string) error {
	requestData := map[string]any{}
	var getMeResponse *GetMessagesResponse
	err := RequestAndGetResponse("POST", MessagesURL+messageID+"/stamps/"+stampID, requestData, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
