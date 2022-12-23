package command

/*
{
  "fact": "The ancestor of all domestic cats is the African Wild Cat which still exists today.",
  "length": 83
}
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type catFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

var catfact = catFact{}

//CatFact returns url of a random cat image
func CatFact(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &catfact)
	s.ChannelMessageSend(m.ChannelID, catfact.Fact)
}
