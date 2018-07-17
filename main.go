package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/thecodeah/taco-bot/src/bot"
)

func main() {
	godotenv.Load("credentials.env", "config.env")

	var config bot.Configuration
	err := envconfig.Process("TACOBOT", &config)
	if err != nil {
		panic(err)
	}

	bot, err := bot.New(config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChannel

	bot.Close()
}
