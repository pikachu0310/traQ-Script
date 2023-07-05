package v1

import (
	"fmt"
)

const loginURL = BaseURL + "/login"

func Login(name string, password string) (map[string]string, error) {
	requestData := map[string]string{"name": name, "password": password}
	res := map[string]string{}
	err := RequestAndGetResponseAndSetCookie("POST", loginURL, requestData, res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
