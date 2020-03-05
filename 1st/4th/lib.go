package main

import "regexp"

//MessageSend SlackにMessageを送る
func MessageSend(channel string, message string) {
	RTM.SendMessage(RTM.NewOutgoingMessage(message, channel))
}

//ResTo 正規表現による一致確認(Slack用)
func ResTo(str string) bool {
	rsp := regexp.MustCompile(str)
	m := rsp.MatchString(EV.Text)
	return m
}
