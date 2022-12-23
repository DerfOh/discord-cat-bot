package command

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// StartTime is when the bot started
var StartTime time.Time

// GitBranch is the current branch the bot is on
var GitBranch string

// GitSummary of the most recent commit
var GitSummary string

// BuildDate is the time of the last commit
var BuildDate string

//About provides info about the bot
func About(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	uptime := GetUptime()

	embed := NewEmbed().
		SetTitle("cat-bot statistics").
		SetDescription("Uptime: "+uptime.String()).
		AddField("Branch: ", GitBranch).
		AddField("Commit: ", GitSummary).
		SetImage("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetThumbnail("https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048").
		SetColor(0x00ff00).MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

// SetStart to be set once the bot is started
func SetStart(t time.Time) {
	StartTime = t
}

// GetUptime Provides the time the bot has been running as type time
func GetUptime() time.Duration {
	return time.Since(StartTime)
}

// SetGit Sets git branch info
func SetGit(gb string, gs string, bd string) {
	GitBranch = gb
	GitSummary = gs
	BuildDate = bd
}
