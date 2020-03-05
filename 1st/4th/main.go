package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

//EV put new slack event
var EV *slack.MessageEvent

//RTM use for sending event to slack
var RTM *slack.RTM

//BotToken Put your slackbot token here
const BotToken string = "YOUR_TOKEN"

//DefaultChannel Put your default channel
const DefaultChannel string = "#YOUR_Channel"

func main() {
	var api *slack.Client = slack.New(BotToken)

	RTM = api.NewRTM()

	go RTM.ManageConnection()

	for msg := range RTM.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Printf("Start connection with Slack\n")
		case *slack.MessageEvent:
			EV = ev
			ListenTo()
		}
	}
}

//ListenTo excute functions under suitable condition
func ListenTo() {
	switch {
	case ResTo(`Hello`):
		MessageSend(EV.Channel, "こんにちは。")
		return
	}
}
