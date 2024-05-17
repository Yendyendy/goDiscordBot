package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token   = flag.String("token", "", "Bot authentication token")
	App     = flag.String("app", "", "Application ID")
	Guild   = flag.String("guild", "", "Guild ID")
	Channel = flag.String("c", "", "channel")
)

func main() {
	flag.Parse()
	// Replace "YOURBOT_TOKEN" with your bot token
	dg, err := discordgo.New("Bot " + *Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Open a connection to the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection: ", err)
		return
	}

	// Send a message every 10 seconds
	for {

		_, err = dg.ChannelMessageSend("1240697599397335111", "This is an automated message.")
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}

		time.Sleep(10 * time.Second) // Adjust the time interval as needed
	}
}
