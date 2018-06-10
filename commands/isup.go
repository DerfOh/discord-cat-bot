package command

import (
	"fmt"
	"net"
	"time"

	"github.com/bwmarrin/discordgo"
)

//IsUp returns the status of a host
func IsUp(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var response string
	host := content[1]
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", host, timeOut)

	if err != nil {
		fmt.Println(err)
		response = "Unable to connect to " + host + " " + err.Error()
	} else {
		response = "Connection established to " + host + " (" + conn.RemoteAddr().String() + ")"
	}
	s.ChannelMessageSend(m.ChannelID, response)
}
