package command

import (
	"bytes"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

//StopServer brings down a digital ocean server
func StopServer(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var droplet string
	host := content[1]

	switch host {
	case "minecraft":
		droplet = "145416705"
	default:
		response := "Unable to determine server ID"
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	url := "https://api.digitalocean.com/v2/droplets/" + droplet + "/actions"
	//fmt.Println("URL:>", url)
	var bearer = "Bearer " + "93fec65a8a9669e4c2b6748bde1ba1124be2950e3f80954230add997c699766d"
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
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	response := resp.Status
	s.ChannelMessageSend(m.ChannelID, response)
}

//curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer b7d03a6947b217efb6f3ec3bd3504582" -d '{"type":"power_off"}' "https://api.digitalocean.com/v2/droplets/3164450/actions"
