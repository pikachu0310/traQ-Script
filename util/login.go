package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"traQ-Script/api"
)

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func Login() (username, password string, err error) {
	fmt.Printf("traQに接続して検索するためにログインが必要です。\n")
	username, password, _ = credentials()
	// fmt.Printf("Username: %s, Password: %s\n", username, "********")
	_, err = api.Login(username, password)
	if err != nil {
		return
	}
	if api.CookieCache == "" {
		err = fmt.Errorf("CookieCache is empty(usernameかpasswordが間違えています)")
		return
	}
	return
}
