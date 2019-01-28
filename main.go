package main

import (
	"log"
	"time"

	"github.com/sbstjn/hanu"
	"github.com/leop0ld/magic-shell-bot/cmd"
	"github.com/spf13/viper"
)

var (
	// Version is the bot version
	Version = "0.0.1"
	// SlackToken will be set by ENV or config file in init()
	SlackToken = ""
)

func init() {
	viper.SetConfigName(".magic-shell-bot")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	SlackToken = viper.GetString("SLACK_TOKEN")
}

func main() {
	bot, err := hanu.New(SlackToken)

	if err != nil {
		log.Fatal(err)
	}

	cmd.Version = Version
	cmd.Start = time.Now()
	cmdList := cmd.List()
	for i := 0; i < len(cmdList); i++ {
		bot.Register(cmdList[i])
	}

	bot.Listen()
}
