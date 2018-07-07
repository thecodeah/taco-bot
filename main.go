package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/thecodeah/Gopher/src/bot"
)

func main() {
	godotenv.Load("token.env", "config.env")

	var config bot.Configuration
	err := envconfig.Process("GOPHER", &config)
	if err != nil {
		panic(err)
	}

	bot, err := bot.New(config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}
