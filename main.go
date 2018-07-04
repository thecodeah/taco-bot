package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Configuration is a pretty darn nice structure, that does stuff a structure
// is supposed to do. I know, it's truly amazing!
type Configuration struct {
	Token  string `required:"true"`
	Prefix string `default:"!"`
}

func main() {
	var conf Configuration
	err := envconfig.Process("config", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session," + err.Error())
	}

	err = session.Open()
	if err != nil {
		log.Fatal("Error creating websocket connection," + err.Error())
	}

	session.AddHandler(messageCreate)

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.EqualFold(message.Content, "ping") {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
