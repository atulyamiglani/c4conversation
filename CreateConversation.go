package main

import "github.com/slack-go/slack"

func CreateConversation(client *slack.Client) {
	params := slack.OpenConversationParameters{
		Users: []string{"hello", "bye"}, 
		ChannelID: "",
		ReturnIM: false,
	}

	client.OpenConversation(&params)
	

}