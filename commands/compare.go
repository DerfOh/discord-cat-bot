package command

import "github.com/bwmarrin/discordgo"

//Compare returns url of a for steam companion similar games
func Compare(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	url := "https://steamcompanion.com/games/index.php?"
	for i := range content {
		if i != 0 {
			url += "&steamID=" + content[i]
		}
	}
	s.ChannelMessageSend(m.ChannelID, url)
}
