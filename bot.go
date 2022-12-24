package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	command "github.com/derfoh/discord-cat-bot/commands"
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
	command.SetStart(startTime)
	command.SetGit(GitBranch, GitSummary, BuildDate)

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
		// servers := goBot.State.Guilds
		// err = goBot.UpdateStatus(0, "with yarn on "+strconv.Itoa(len(servers))+" servers.")
		// if err != nil {
		// 	fmt.Println("Error attempting to set my status")
		// }
	})
	goBot.AddHandler(voiceUpdateHandler)

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

		// only certain users can issue these commands
		// TODO: Make a command that lists the dropplets given a particular personal access token
		if m.Author.ID == config.BotOwner {
			if strings.Contains(m.Content, "serverstart") {
				go command.ServerStart(content, s, m)
			}

			if strings.Contains(m.Content, "serverstop") {
				go command.ServerStop(content, s, m)
			}

			if strings.Contains(m.Content, "serverlist") {
				go command.ServerList(content, s, m)
			}
		}

		// if the message contains the string then call a function that responds with a string
		if strings.Contains(m.Content, "help") {
			go command.Help(content, s, m)
		} else if strings.Contains(m.Content, "soundboard") || strings.Contains(m.Content, "sb") {
			command.SoundBoard(content, s, m)
		} else if strings.Contains(m.Content, "cat") {
			go command.Cat(content, s, m)
		} else if strings.Contains(m.Content, "dog") {
			go command.Dog(content, s, m)
		} else if strings.Contains(m.Content, "fact") {
			go command.CatFact(content, s, m)
		} else if strings.Contains(m.Content, "8ball") {
			go command.EightBall(content, s, m)
		} else if strings.Contains(m.Content, "complete") || strings.Contains(m.Content, "finish") || strings.Contains(m.Content, "predict") {
			go command.CompleteThis(content, s, m)
		} else if strings.Contains(m.Content, "compare") {
			go command.Compare(content, s, m)
		} else if strings.Contains(m.Content, "isup") {
			go command.IsUp(content, s, m)
		} else if strings.Contains(m.Content, "date") {
			go command.Date(content, s, m)
		} else if strings.Contains(m.Content, "time") {
			go command.Time(content, s, m)
		} else if strings.Contains(m.Content, "vote") {
			go command.Vote(content, s, m)
		} else if strings.Contains(m.Content, "about") {
			go command.About(content, s, m)
		} else if strings.Contains(m.Content, "meow") {
			newContent := "meow meow"
			content := strings.Split(newContent, " ")
			go command.SoundBoard(content, s, m)
			go command.Meow(content, s, m)
		} else if strings.Contains(m.Content, "bark") {
			go command.Bark(content, s, m)
		} else if strings.Contains(m.Content, "love") {
			go command.Love(content, s, m)
		} else if strings.Contains(m.Content, "thanks") || strings.Contains(m.Content, "Thanks") || strings.Contains(m.Content, "Thank you") || strings.Contains(m.Content, "thank you") {
			go command.Thanks(content, s, m)
		} else if strings.Contains(m.Content, "good night") || strings.Contains(m.Content, "Good night") || strings.Contains(m.Content, "goodnight") {
			go command.GoodNight(content, s, m)
		} else if strings.Contains(m.Content, "good morning") || strings.Contains(m.Content, "Good morning") {
			go command.GoodMorning(content, s, m)
		} else if strings.Contains(m.Content, "hello") || strings.Contains(m.Content, "Hello") {
			go command.Hello(content, s, m)
		} else if strings.Contains(m.Content, "good bye") || strings.Contains(m.Content, "Good bye") || strings.Contains(m.Content, "Goodbye") || strings.Contains(m.Content, "Good bye") {
			go command.Goodbye(content, s, m)
		} else if strings.Contains(m.Content, "ping") {
			go command.Ping(content, s, m)
		} else {
			go command.CompleteThis(content, s, m)
		}

	}

}

func voiceUpdateHandler(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	fmt.Println("ChannelID: " + vs.ChannelID)
	// fmt.Println("Deaf: " + vs.Deaf)
	fmt.Println("GuildID: " + vs.GuildID)
	// fmt.Println("Mute: " + vs.Mute)
	// fmt.Println("SelfDeaf: " + vs.SelfDeaf)
	// fmt.Println("SelfMute: " + vs.SelfMute)
	fmt.Println("SessionID: " + vs.SessionID)
	// fmt.Println("Suppress: " + vs.Suppress)
	fmt.Println("UserID: " + vs.UserID)

	if vs.UserID == BotID {
		return
	}

	// Find the channel that the message came from.
	c, err := s.State.Channel(vs.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}

	// Look for the notification target in that guild's current voice states.
	for _, states := range g.VoiceStates {
		if states.UserID == config.BotOwner && vs.ChannelID == states.ChannelID {
			fmt.Println("found bot owner")
			//notification.Notify("userjoined", s, vs.ChannelID)
			return
		}
	}
}

// Stop ends closes the connection
func Stop() {
	// Cleanly close down the Discord session.
	goBot.Close()
}
