package command

import (
	"github.com/bwmarrin/discordgo"
)

//GoodNight returns meow response
func GoodNight(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	response := "Good night " + m.Author.Mention() + "."
	s.ChannelMessageSend(m.ChannelID, response)
}
