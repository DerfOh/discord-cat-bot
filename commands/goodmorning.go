package command

import (
	"github.com/bwmarrin/discordgo"
)

//GoodMorning returns meow response
func GoodMorning(s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Good morning " + m.Author.Mention() + "!"
	s.ChannelMessageSend(m.ChannelID, response)
}
