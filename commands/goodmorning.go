package command

import (
	"github.com/bwmarrin/discordgo"
)

//GoodMorning returns meow response
func GoodMorning(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Good morning " + m.Author.Mention() + "!"
	s.ChannelMessageSend(m.ChannelID, response)
}
