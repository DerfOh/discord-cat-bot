package command

import (
	"github.com/bwmarrin/discordgo"
)

//Goodbye returns goodbye response
func Goodbye(s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Good night " + m.Author.Mention() + "."
	s.ChannelMessageSend(m.ChannelID, response)
}
