package command

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

//Bark returns bark response
func Bark(s *discordgo.Session, m *discordgo.MessageCreate) {

	barks := make([]string, 0)
	barks = append(barks,
		"woof",
		"bark",
		"arfarf",
		"ruff")

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	response := fmt.Sprint(barks[rand.Intn(len(barks))])
	s.ChannelMessageSend(m.ChannelID, response)
}
