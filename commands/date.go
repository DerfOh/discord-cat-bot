package command

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//Date returns the date of the system catbot is running on.
func Date(s *discordgo.Session, m *discordgo.MessageCreate) {
	t := time.Now()
	response := "The date is " + t.Format("01-02-2006") + " in catsville."
	s.ChannelMessageSend(m.ChannelID, response)
}
