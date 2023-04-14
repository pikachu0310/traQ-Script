package api

import (
	"fmt"
	"time"
)

const meURL = BaseURL + "/users/me"

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

type Tag struct {
	TagId     string    `json:"tagId"`
	Tag       string    `json:"tag"`
	IsLocked  bool      `json:"isLocked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetMe() (*GetMeResponse, error) {
	var getMeResponse *GetMeResponse
	err := RequestAndGetResponse("GET", meURL, nil, &getMeResponse)
	if err != nil {
		fmt.Println(err)
		return getMeResponse, err
	}
	return getMeResponse, nil
}
