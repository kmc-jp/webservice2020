package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"../idol"
	"github.com/nlopes/slack"
)

//EV put new slack events
var EV *slack.MessageEvent

//RTM use for sending events to slack
var RTM *slack.RTM

//BotToken Put your slackbot token here
const BotToken = "YOUR_TOKEN"

//DictPath Put dictionary path
var DictPath string

func init() {
	DictPath = filepath.Join("..", "data.json")
	var err error = idol.MakeDict(DictPath)
	if err != nil {
		panic(err.Error())
	}
}

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

//ListenTo Manage incomming messages
func ListenTo() {
	switch {
	case strings.Contains(EV.Text, "ちゃん"):
		NameCheck()
	case strings.Contains(EV.Text, "デレマスランダム"):
		Random()
	}
}

//Random chose one Idol
func Random() {
	var Info []idol.Idol = idol.Read(DictPath)
	rand.Seed(time.Now().UnixNano())

	var SendText string = MakeText(Info[int(rand.Float64()*float64(len(Info)))])
	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))
}

//NameCheck find the idol from dict
func NameCheck() {
	var Info []idol.Idol = idol.Read(DictPath)
	var NAME string = strings.Split(EV.Text, "ちゃん")[0]
	var FoundNUM int

	if NAME == "" {
		return
	}

	for i, info := range Info {
		if info.Name == NAME {
			FoundNUM = i
			goto Found
		}
	}
	return

Found:
	var SendText string = MakeText(Info[FoundNUM])

	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))
}

//MakeText Make idol introduction
func MakeText(Idol idol.Idol) string {
	var SendText string
	SendText = "名前: " + Idol.Name + "\n"
	SendText += "年齢: " + Idol.Age + "\n"
	SendText += "身長: " + Idol.Height + "\n"
	SendText += "利き手: " + Idol.Hand + "\n"
	SendText += "趣味: " + Idol.Hobby + "\n"
	SendText += "誕生日: " + Idol.Birth + "\n"
	SendText += "3Size: " + Idol.B + "/" + Idol.W + "/" + Idol.H + "\n"
	SendText += Idol.GetImgURL() + "\n"

	return SendText
}
