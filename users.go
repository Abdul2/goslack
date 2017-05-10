package  main

import (

	"fmt"
	"github.com/bluele/slack"
	"os"
)

//list slack users 
func main() {

	token := os.Getenv("hoddspslacktoken")
	api := slack.New(token)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}
}
