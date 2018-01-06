package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
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
