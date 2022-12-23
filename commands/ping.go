package command

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//Ping returns response time
func Ping(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	start := time.Now()
	s.ChannelMessageSend(m.ChannelID, "")
	elapsed := time.Since(start)
	response := "Pong! " + elapsed.String()

	s.ChannelMessageSend(m.ChannelID, response)
}
