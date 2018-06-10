package command

import "github.com/bwmarrin/discordgo"

//Vote returns the date of the system catbot is running on.
func Vote(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
	s.MessageReactionAdd(m.ChannelID, m.ID, "👎")
}
