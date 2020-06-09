package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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
const BotToken string = "YOUR_TOKEN"

func init() {
	var x []idol.Idol = idol.Get()
	b, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		panic("Marshal Error!")
	}
	err = ioutil.WriteFile(filepath.Join("..", "data.json"), b, os.ModePerm)
	if err != nil {
		panic("Write Error!")
	}
}

func main() {
	// fmt.Printf("%#v\n", idol.Read())

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

func ListenTo() {
	switch {
	case strings.Contains(EV.Text, "ちゃん"):
		NameCheck()
	case strings.Contains(EV.Text, "デレマスランダム"):
		Random()
	}
}

func Random() {
	var Info []idol.Idol = idol.Read()
	rand.Seed(time.Now().UnixNano())

	var FoundNum int = int(rand.Float64() * float64(len(Info)))
	var SendText string

	SendText = "名前: " + Info[FoundNum].Name + "\n"
	SendText += "年齢: " + Info[FoundNum].Age + "\n"
	SendText += "身長: " + Info[FoundNum].Height + "\n"
	SendText += "利き手: " + Info[FoundNum].Hand + "\n"
	SendText += "趣味: " + Info[FoundNum].Hobby + "\n"
	SendText += "誕生日: " + Info[FoundNum].Birth + "\n"
	SendText += "3Size: " + Info[FoundNum].B + "/" + Info[FoundNum].W + "/" + Info[FoundNum].H + "\n"
	SendText += idol.GetImgURL(Info[FoundNum].URL) + "\n"
	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))
}

func NameCheck() {
	var Info []idol.Idol = idol.Read()
	var NAME string = strings.Split(EV.Text, "ちゃん")[0]
	var FoundNUM int

	for i, info := range Info {
		if info.Name == NAME {
			FoundNUM = i
			goto Found
		}
	}
	return
Found:
	var SendText string
	SendText = "名前: " + Info[FoundNUM].Name + "\n"
	SendText += "年齢: " + Info[FoundNUM].Age + "\n"
	SendText += "身長: " + Info[FoundNUM].Height + "\n"
	SendText += "3Size: " + Info[FoundNUM].B + "/" + Info[FoundNUM].W + "/" + Info[FoundNUM].H + "\n"
	SendText += "身長: " + idol.GetImgURL(Info[FoundNUM].URL) + "\n"
	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))
}
