package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

//EightBall returns an answer when the command !8ball is used
func EightBall(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	url := "https://8ball.delegator.com/magic/JSON/"
	for i := range content {
		if i != 0 {
			url += "+" + content[i]
		} else {
			url += content[i]
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	defer resp.Body.Close()
	jStr, _ := ioutil.ReadAll(resp.Body)

	type Inner struct {
		QuestionKey string `json:"question"`
		AnswerKey   string `json:"answer"`
		TypeKey     string `json:"type"`
	} // Define struct to match structure
	type Container struct {
		MagicKey Inner `json:"magic"`
	}
	var cont Container
	if err := json.Unmarshal([]byte(jStr), &cont); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(url)
	s.ChannelMessageSend(m.ChannelID, cont.MagicKey.AnswerKey)
}
