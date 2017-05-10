package main

import (

"github.com/bluele/slack"

	"os"
)

const (

	channelName = "random"
)
//send messages to a channel
func main() {

	
	//to obtain a token go to this site :- https://api.slack.com/custom-integrations/legacy-tokens
	token := os.Getenv("hoddspslacktoken")
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, "hi slackers - test message", nil)
	if err != nil {
		panic(err)
	}


}
