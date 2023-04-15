package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://q.trap.jp/api/v3"

var CookieCache = ""

func RequestAndGetResponse(method string, URL string, requestBodyData map[string]any, responseData any) error {
	requestJson, _ := json.Marshal(requestBodyData)
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", CookieCache)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Println(response)
	// fmt.Printf("%#v", response)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal(responseBody, &responseData)
	return nil
}

func RequestAndGetResponseByJson(method string, URL string, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", CookieCache)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return responseBody, nil
}

func RequestAndGetResponseAndSetCookie(method string, URL string, requestBodyData map[string]string, responseData any) error {
	requestJson, _ := json.Marshal(requestBodyData)
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", CookieCache)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	CookieCache = response.Header.Get("Set-Cookie")
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal(responseBody, &responseData)
	return nil
}
