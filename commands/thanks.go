package command

import (
	"github.com/bwmarrin/discordgo"
)

//Thanks returns thanks response
func Thanks(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "You're welcome " + m.Author.Mention() + "!"
	s.ChannelMessageSend(m.ChannelID, response)
}
