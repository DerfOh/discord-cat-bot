package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/derfoh/discord-cat-bot/bot"
	"github.com/derfoh/discord-cat-bot/config"
)

func main() {
	err := config.ReadConfig()
	checkExit(err)

	bot.Start()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	//bot.Stop()
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
