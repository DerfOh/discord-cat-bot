package command

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
	config "github.com/derfoh/discord-cat-bot/config"
)

//ServerStop brings down a digital ocean server
func ServerStop(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	host := content[1]
	droplet := strconv.Itoa(GetServerID(host))

	url := "https://api.digitalocean.com/v2/droplets/" + droplet + "/actions"
	var bearer = "Bearer " + config.DigitalOceanToken
	var jsonStr = []byte(`{"type":"power_off"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	// TODO: parse response body for readable status
	response := string(body)
	s.ChannelMessageSend(m.ChannelID, response)
}

//curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer b7d03a6947b217efb6f3ec3bd3504582" -d '{"type":"power_off"}' "https://api.digitalocean.com/v2/droplets/3164450/actions"
