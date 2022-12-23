package command

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//Time returns the date of the system catbot is running on.
func Time(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	t := time.Now()
	response := "The time is " + t.Format("15:04") + " in cat town."
	s.ChannelMessageSend(m.ChannelID, response)
}
