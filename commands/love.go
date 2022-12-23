package command

import (
	"github.com/bwmarrin/discordgo"
)

//Love returns a loving response
func Love(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Love ya too"
	s.ChannelMessageSend(m.ChannelID, response)
}
