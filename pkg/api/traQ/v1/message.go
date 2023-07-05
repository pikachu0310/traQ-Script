package v1

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

type GetMessageStamps struct {
	Stamps []GetMessageStamp
}

type GetMessageStamp struct {
	UserId    string    `json:"userId"`
	StampId   string    `json:"stampId"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetMessages(requestData map[string]any) (*GetMessagesResponse, error) {
	// requestData := map[string]string{"from": "a4f4ca7e-054e-4d8b-842e-8b777c353a5d"}
	var getMessagesResponse *GetMessagesResponse
	err := RequestAndGetResponse("GET", getMessagesURL, requestData, &getMessagesResponse)
	if err != nil {
		fmt.Println(err)
		return &GetMessagesResponse{}, err
	}
	return getMessagesResponse, nil
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

func GetStamps(messageID string) (*GetMessageStamps, error) {
	var getMessageStamps *GetMessageStamps
	err := RequestAndGetResponse("GET", MessagesURL+messageID+"/stamps", nil, &getMessageStamps)
	if err != nil {
		fmt.Println(err)
		return &GetMessageStamps{}, err
	}
	return getMessageStamps, nil
}

func PostStamp(messageID string, stampID string) error {
	requestData := map[string]any{"count": 100}
	var getMeResponse *GetMessagesResponse
	err := RequestAndGetResponse("POST", MessagesURL+messageID+"/stamps/"+stampID, requestData, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
