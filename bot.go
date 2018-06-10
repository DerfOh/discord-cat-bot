package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/derfoh/discord-cat-bot/commands"
	"github.com/derfoh/discord-cat-bot/config"
)

// BotID is the id set by discord in the main function
var BotID string
var goBot *discordgo.Session
var startTime time.Time

// GetUptime Provides the time the bot has been running as type time
func GetUptime() time.Duration {
	return time.Since(startTime)
}

// Start starts opens connections and starts the bot
func Start(GitBranch string, GitSummary string, BuildDate string) {
	startTime = time.Now()
	about.SetStart(startTime)
	about.SetGit(GitBranch, GitSummary, BuildDate)

	// connect
	goBot, err := discordgo.New("Bot " + config.Token)
	checkExit(err)

	// get bot info and set bot id with user info
	u, err := goBot.User("@me")
	checkLog(err)

	BotID = u.ID

	// Add handlers
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		servers := goBot.State.Guilds
		err = goBot.UpdateStatus(0, "with yarn on "+strconv.Itoa(len(servers))+" servers.")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
	})

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

	// check for prefix or @ mention
	if strings.HasPrefix(m.Content, config.BotPrefix) || strings.Contains(m.Content, BotID) {
		// ignore the bots messages
		if m.Author.ID == BotID {
			return
		}

		// if the message contains the string then call a function that responds with a string
		if strings.Contains(m.Content, "cat") {
			command.Cat(s, m)
		}

		if strings.Contains(m.Content, "fact") {
			command.CatFact(s, m)
		}

		if strings.Contains(m.Content, "8ball") || strings.Contains(m.Content, "should") {
			command.EightBall(content, s, m)
		}

		if strings.Contains(m.Content, "compare") {
			command.Compare(content, s, m)
		}

		if strings.Contains(m.Content, "isup") {
			command.IsUp(content, s, m)
		}

		if strings.Contains(m.Content, "date") {
			command.Date(s, m)
		}

		if strings.Contains(m.Content, "time") {
			command.Time(s, m)
		}

		if strings.Contains(m.Content, "vote") {
			command.Vote(s, m)
		}
		if strings.Contains(m.Content, "about") {
			command.About(s, m)
		}

		if strings.Contains(m.Content, "meow") {
			command.Meow(s, m)
		}

		// simple commands that aren't in the commands folder
		if m.Content == "!iloveyou" {
			s.ChannelMessageSend(m.ChannelID, "I know "+m.Author.Username)
		}

		if strings.Contains(m.Content, "ping") {
			start := time.Now()
			s.ChannelMessageSend(m.ChannelID, "")
			elapsed := time.Since(start)
			s.ChannelMessageSend(m.ChannelID, "Pong! "+elapsed.String())
		}

		if strings.Contains(m.Content, "thanks") || strings.Contains(m.Content, "Thanks") || strings.Contains(m.Content, "Thank you") || strings.Contains(m.Content, "thank you") {
			s.ChannelMessageSend(m.ChannelID, "You're welcome "+m.Author.Mention()+"!")
		}

		if strings.Contains(m.Content, "good night") || strings.Contains(m.Content, "Good night") || strings.Contains(m.Content, "goodnight") {
			s.ChannelMessageSend(m.ChannelID, "Good night "+m.Author.Mention()+".")
		}

		if strings.Contains(m.Content, "good morning") || strings.Contains(m.Content, "Good morning") {
			s.ChannelMessageSend(m.ChannelID, "Good morning "+m.Author.Mention()+"!")
		}

		if strings.Contains(m.Content, "hello") || strings.Contains(m.Content, "Hello") {
			s.ChannelMessageSend(m.ChannelID, "Welcome "+m.Author.Mention()+".")
		}

		if strings.Contains(m.Content, "good bye") || strings.Contains(m.Content, "Good bye") || strings.Contains(m.Content, "Goodbye") || strings.Contains(m.Content, "Good bye") {
			s.ChannelMessageSend(m.ChannelID, "See ya "+m.Author.Mention()+"!")
		}

	}

}

// Stop ends closes the connection
func Stop() {
	// Cleanly close down the Discord session.
	goBot.Close()
}
