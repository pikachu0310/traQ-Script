package main

import (
	"fmt"
	"os"

	"traQ-Script/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	name := "pikachu"
	password := os.Getenv("PASSWORD")
	_, err = api.Login(name, password)
	if err != nil {
		panic(err)
	}
	if api.CookieCache == "" {
		panic("CookieCache is empty")
	}

	body, err := api.GetMessages()
	if err != nil {
		panic(err)
	}
	fmt.Println(*body)
	// fmt.Printf("%#v", body)
}
