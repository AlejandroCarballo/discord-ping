package main

import (
	"fmt"

	"github.com/AlejandroCarballo/discord-ping/bot"
	"github.com/AlejandroCarballo/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
