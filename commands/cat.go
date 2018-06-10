package command

/*
<response>
    <data>
        <images>
            <image>
                <url>http://25.media.tumblr.com/tumblr_lxp5hbDPXR1qzzfdxo1_400.gif</url>
                <id>914</id>
                <source_url>http://thecatapi.com/?id=914</source_url>
            </image>
        </images>
    </data>
</response>
*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type catpicStruct struct {
	XMLName   xml.Name `xml:"response"`
	URL       string   `xml:"data>images>image>url"`
	ID        string   `xml:"data>images>image>id"`
	SourceURL string   `xml:"data>images>image>source_url"`
}

var catpic = catpicStruct{}

//Cat returns url of a random cat image
func Cat(s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("http://thecatapi.com/api/images/get?format=xml&results_per_page=1&type=gif")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//xml.Unmarshal(body, &catpic)
	err = xml.Unmarshal([]byte(body), &catpic)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	response := catpic.URL + " ```Source Info: " + catpic.SourceURL + "\n```"
	s.ChannelMessageSend(m.ChannelID, response)
}
