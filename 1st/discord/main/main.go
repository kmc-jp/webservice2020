package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"../wikipedia"
	"github.com/bwmarrin/discordgo"
)

//Setting put setting data
type Setting struct {
	//DiscordToken put discord token
	DiscordToken string `json:"discord"`
}

//Settings pit setting data
var Settings Setting

func init() {
	b, e := ioutil.ReadFile(filepath.Join("..", "credit.json"))
	if e != nil {
		panic(e.Error())
	}
	e = json.Unmarshal(b, &Settings)
	if e != nil {
		panic(e.Error())
	}
}

func main() {
	if Settings.DiscordToken == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(Settings.DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(watch)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func watch(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "/wikipedia_random") {
		return
	}
	var wiki wikipedia.Wikipedia
	wiki.Get()

	_, err := s.ChannelMessageSend(m.ChannelID, "**"+wiki.Title+"**\n"+wiki.Text)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	return
}
