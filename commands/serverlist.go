package command

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	config "github.com/derfoh/discord-cat-bot/config"
)

// DropletID is the integer value of the droplet identified by name
var DropletID int

// DropletList defines the structure for the json object returned by digital ocean api calls
type DropletList struct {
	Droplets []struct {
		ID               int           `json:"id"`
		Name             string        `json:"name"`
		Memory           int           `json:"memory"`
		Vcpus            int           `json:"vcpus"`
		Disk             int           `json:"disk"`
		Locked           bool          `json:"locked"`
		Status           string        `json:"status"`
		Kernel           interface{}   `json:"kernel"`
		CreatedAt        time.Time     `json:"created_at"`
		Features         []interface{} `json:"features"`
		BackupIds        []interface{} `json:"backup_ids"`
		NextBackupWindow interface{}   `json:"next_backup_window"`
		SnapshotIds      []int         `json:"snapshot_ids"`
		Image            struct {
			ID            int           `json:"id"`
			Name          string        `json:"name"`
			Distribution  string        `json:"distribution"`
			Slug          interface{}   `json:"slug"`
			Public        bool          `json:"public"`
			Regions       []string      `json:"regions"`
			CreatedAt     time.Time     `json:"created_at"`
			MinDiskSize   int           `json:"min_disk_size"`
			Type          string        `json:"type"`
			SizeGigabytes float64       `json:"size_gigabytes"`
			Description   string        `json:"description"`
			Tags          []interface{} `json:"tags"`
			Status        string        `json:"status"`
		} `json:"image"`
		VolumeIds []interface{} `json:"volume_ids"`
		Size      struct {
			Slug         string   `json:"slug"`
			Memory       int      `json:"memory"`
			Vcpus        int      `json:"vcpus"`
			Disk         int      `json:"disk"`
			Transfer     float64  `json:"transfer"`
			PriceMonthly float64  `json:"price_monthly"`
			PriceHourly  float64  `json:"price_hourly"`
			Regions      []string `json:"regions"`
			Available    bool     `json:"available"`
		} `json:"size"`
		SizeSlug string `json:"size_slug"`
		Networks struct {
			V4 []struct {
				IPAddress string `json:"ip_address"`
				Netmask   string `json:"netmask"`
				Gateway   string `json:"gateway"`
				Type      string `json:"type"`
			} `json:"v4"`
			V6 []interface{} `json:"v6"`
		} `json:"networks"`
		Region struct {
			Name      string   `json:"name"`
			Slug      string   `json:"slug"`
			Features  []string `json:"features"`
			Available bool     `json:"available"`
			Sizes     []string `json:"sizes"`
		} `json:"region"`
		Tags []interface{} `json:"tags"`
	} `json:"droplets"`
	Links struct {
	} `json:"links"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

// ServerList returns a list of servers associated with a personal access token
func ServerList(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	url := "https://api.digitalocean.com/v2/droplets"

	var bearer = "Bearer " + config.DigitalOceanToken
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jStr, _ := ioutil.ReadAll(resp.Body)
	res := DropletList{}

	if err := json.Unmarshal([]byte(jStr), &res); err != nil {
		log.Fatal(err)
	}

	// Gets all of the droplets in the list

	for i := range res.Droplets {
		// fmt.Println(res.Droplets[i].Name)
		// go s.ChannelMessageSend(m.ChannelID, res.Droplets[i].Name+" --- "+strconv.Itoa(res.Droplets[i].ID))
		go s.ChannelMessageSend(m.ChannelID, res.Droplets[i].Name)
	}
}

// GetServerID takes name and returns server ID
func GetServerID(name string) int {
	url := "https://api.digitalocean.com/v2/droplets"

	var bearer = "Bearer " + config.DigitalOceanToken
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jStr, _ := ioutil.ReadAll(resp.Body)
	res := DropletList{}

	if err := json.Unmarshal([]byte(jStr), &res); err != nil {
		log.Fatal(err)
	}

	for i := range res.Droplets {
		if res.Droplets[i].Name == name {
			DropletID = res.Droplets[i].ID
		}
	}
	return DropletID
}
