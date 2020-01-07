package command

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

//Meow returns meow response
func Meow(s *discordgo.Session, m *discordgo.MessageCreate) {

	meows := make([]string, 0)
	meows = append(meows,
		"meow",
		"mow",
		"purrpurr",
		"мяу",
		"喵",
		"야옹",
		"	മ്യാവു",
		" 	मियांउ",
		" 	เหมียว")

	rand.Seed(time.Now().Unixnano()) // initialize global pseudo random generator
	response := fmt.Sprint(meows[rand.Intn(len(meows))])
	s.ChannelMessageSend(m.ChannelID, response)
}
