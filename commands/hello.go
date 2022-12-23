package command

import (
	"github.com/bwmarrin/discordgo"
)

//Hello returns hello response
func Hello(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Welcome " + m.Author.Mention() + "."
	s.ChannelMessageSend(m.ChannelID, response)
}
