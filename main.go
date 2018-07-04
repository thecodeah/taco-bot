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

	"github.com/thecodeah/Gopher/src/commands"
)

// Configuration is a pretty darn nice structure, that does stuff a structure
// is supposed to do. I know, it's truly amazing!
type Configuration struct {
	Token  string `required:"true"`
	Prefix string `default:"!"`
}

// Config is the name of a variable that i'll document later.
var Config Configuration

func main() {
	err := envconfig.Process("config", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	session, err := discordgo.New("Bot " + Config.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session," + err.Error())
	}

	err = session.Open()
	if err != nil {
		log.Fatal("Error creating websocket connection," + err.Error())
	}

	session.AddHandler(messageCreate)

	registerCommands()

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

	if strings.HasPrefix(message.Content, Config.Prefix) {
		arguments := strings.Fields(message.Content)
		cmdName := arguments[0]

		// Removing the command from the arguments slice
		arguments = arguments[1:]

		var commandInfo commands.CommandInfo
		commandInfo.Session = session
		commandInfo.Message = message
		commandInfo.Args = arguments

		cmdFunction, found := commands.Get(strings.TrimPrefix(cmdName, Config.Prefix))
		if !found {
			return
		}
		c := *cmdFunction
		c(commandInfo)

	}
}

func registerCommands() {
	commands.Register("ping", commands.PingCommand)
}
