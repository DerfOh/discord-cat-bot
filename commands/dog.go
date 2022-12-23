package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type dogpicStruct struct {
	Link string `json:"message"`
}

var dogpic = dogpicStruct{}

//Dog returns url of a random cat image
func Dog(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &dogpic)
	fmt.Println("dogpic:" + dogpic.Link)

	response := dogpic.Link
	s.ChannelMessageSend(m.ChannelID, response)
}
