package main

import (

"github.com/bluele/slack"

	"os"
)



const (

	channelName = "random"
)

func main() {

	token := os.Getenv("hoddspslacktoken")
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, "hi slackers - test message", nil)
	if err != nil {
		panic(err)
	}


}