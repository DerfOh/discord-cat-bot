package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Token retrieved from discord api docs and placed into the config.json
	Token string
	// BotPrefix is the prefix defined in the config.json
	BotPrefix string
	// BotOwner is the user ID of the user running the bot
	BotOwner string
	// DigitalOceanToken Digital Ocean personal access token (PAT) generated from: https://cloud.digitalocean.com/settings/applications
	DigitalOceanToken string
	// Openai token
	OpenaiToken string

	config *configStruct
)

type configStruct struct {
	Token             string `json:"Token"`
	BotPrefix         string `json:"BotPrefix"`
	BotOwner          string `json:"BotOwner"`
	DigitalOceanToken string `json:"DigitalOceanToken"`
	OpenaiToken       string `json:"OpenaiToken"`
}

// ReadConfig reads the contents of the config file
func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")
	checkExit(err)

	// if reading out a json the file needs to be cast as a string otherwise it
	//	it will be a byte array
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	checkExit(err)

	Token = config.Token
	BotPrefix = config.BotPrefix
	BotOwner = config.BotOwner
	DigitalOceanToken = config.DigitalOceanToken
	OpenaiToken = config.OpenaiToken

	return nil
}

// basic error checker for go, logs then keeps running
func checkLog(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// basic error checker for go, logs then exits
func checkExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
