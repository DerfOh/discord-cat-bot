package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/derfoh/discord-cat-bot/commands"
	"github.com/derfoh/discord-cat-bot/config"
)

// BotID is the id set by discord in the main function
var BotID string
var goBot *discordgo.Session

// Start starts opens connections and starts the bot
func Start(GitCommit string, GitBranch string, GitState string, GitSummary string, BuildDate string, Version string) {

	// connect
	goBot, err := discordgo.New("Bot " + config.Token)
	checkExit(err)

	// get bot info and set bot id with user info
	u, err := goBot.User("@me")
	checkLog(err)

	BotID = u.ID

	// Add handlers
	goBot.AddHandler(messageHandler)

	// Open connection
	err = goBot.Open()
	checkExit(err)

	// Log bot is running
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.\nPress CTRL-C to exit.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// split the contents of the string into an array
	content := strings.Split(m.Content, " ")

	// check for prefix
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		// ignore the bots messages
		if m.Author.ID == BotID {
			return
		}

		// if the message contains the string then call a function that responds with a string
		if strings.Contains(m.Content, "cat") {
			response := command.Cat()
			s.ChannelMessageSend(m.ChannelID, response)
		}

		if strings.Contains(m.Content, "8ball") {
			response := command.EightBall()
			s.ChannelMessageSend(m.ChannelID, response)
		}

		if strings.Contains(m.Content, "compare") {
			response := command.Compare(content)
			s.ChannelMessageSend(m.ChannelID, response)
		}

		if strings.Contains(m.Content, "isup") {
			response := command.IsUp(content)
			s.ChannelMessageSend(m.ChannelID, response)
		}

		// return a simple message
		if m.Content == "!time" {
			t := time.Now()
			s.ChannelMessageSend(m.ChannelID, "The time is "+t.Format("15:04")+" in catsville.")
		}

		if m.Content == "!date" {
			t := time.Now()
			s.ChannelMessageSend(m.ChannelID, "The date is "+t.Format("01-02-2006")+" in catsville.")
		}

		if m.Content == "!iloveyou" {
			s.ChannelMessageSend(m.ChannelID, "I know."+m.Author.Username)
		}

		if strings.Contains(m.Content, "meow") {
			start := time.Now()
			s.ChannelMessageSend(m.ChannelID, "")
			elapsed := time.Since(start)
			s.ChannelMessageSend(m.ChannelID, "(Meow took "+elapsed.String()+")")
		}

		if strings.Contains(m.Content, "ping") {
			start := time.Now()
			s.ChannelMessageSend(m.ChannelID, "")
			elapsed := time.Since(start)
			s.ChannelMessageSend(m.ChannelID, "Pong! "+elapsed.String())
		}

	}

}

// Stop ends closes the connection
func Stop() {
	// Cleanly close down the Discord session.
	goBot.Close()
}

// basic error checker for go, logs then keeps running
func checkLog(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// basic error checker for go, logs then exits
func checkExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
