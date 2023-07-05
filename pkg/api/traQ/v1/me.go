package v1

import (
	"fmt"
	"time"
)

const meURL = BaseURL + "/users/me"
const subscriptionsURL = BaseURL + "/users/me/subscriptions"

type GetMeResponse struct {
	Id          string    `json:"id"`
	Bio         string    `json:"bio"`
	Groups      []string  `json:"groups"`
	Tags        []Tag     `json:"tags"`
	UpdatedAt   time.Time `json:"updatedAt"`
	LastOnline  time.Time `json:"lastOnline"`
	TwitterId   string    `json:"twitterId"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	IconFileId  string    `json:"iconFileId"`
	Bot         bool      `json:"bot"`
	State       int       `json:"state"`
	Permissions []string  `json:"permissions"`
	HomeChannel string    `json:"homeChannel"`
}

type GetMeSubscriptions struct {
	ChannelId string `json:"channelId"`
	Level     int    `json:"level"`
}

type Tag struct {
	TagId     string    `json:"tagId"`
	Tag       string    `json:"tag"`
	IsLocked  bool      `json:"isLocked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetMe() (getMeResponse *GetMeResponse, err error) {
	err = RequestAndGetResponse("GET", meURL, nil, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetSubscriptions() (getMeSubscriptions *[]GetMeSubscriptions, err error) {
	err = RequestAndGetResponse("GET", subscriptionsURL, nil, &getMeSubscriptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PutSubscriptions(channelID string, level int) error {
	// チャンネル購読レベル
	// 0：無し
	// 1：未読管理
	// 2：未読管理+通知
	requestData := map[string]any{"level": level}
	err := RequestAndGetResponse("PUT", meURL+"/subscriptions/"+channelID, requestData, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
