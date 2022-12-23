package command

import (
	"github.com/bwmarrin/discordgo"
)

//Help lists all functions in the command package
func Help(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	commandList := "about\n catfact \ndate \ngoodmorning \nhello \nisup \nmeow \nthanks \nvote \ncat \ncompare \neightball \ngoodbye \ngoodnight \nhelp  \nlove \nping \ntime"
	embed := NewEmbed().
		AddField("Commands: ", commandList).
		SetImage("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetThumbnail("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetColor(0x00ff00).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
