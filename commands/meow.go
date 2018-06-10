package command

import (
	"github.com/bwmarrin/discordgo"
)

//Meow returns meow response
func Meow(s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Meow"
	s.ChannelMessageSend(m.ChannelID, response)
}
