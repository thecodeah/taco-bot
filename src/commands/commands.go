package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandInfo stores information about the message sent by the player.
type CommandInfo struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Args    []string
}

// Command is a function that requires CommandInfo as its argument.
type Command func(CommandInfo)

// CommandMap stores all command functions by their name.
type CommandMap map[string]Command

// CommandHandler stores command information/state.
type CommandHandler struct {
	commands CommandMap
	config   Config
}

// Config stores command handler configurations.
type Config struct {
	Prefix string
}

// New creates a new command handler.
func New(config Config) (ch *CommandHandler) {
	ch = &CommandHandler{
		commands: make(CommandMap),
		config:   config,
	}
	return
}

// Register registers a command to be handled by the command handler.
func (ch CommandHandler) Register(name string, command Command) {
	ch.commands[name] = command
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (ch CommandHandler) Get(name string) (*Command, bool) {
	command, found := ch.commands[name]
	return &command, found
}

// Process processes incoming messages and calls the command's
// function.
func (ch CommandHandler) Process(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, ch.config.Prefix) {
		arguments := strings.Fields(message.Content)
		cmdName := arguments[0]

		// Removing the command from the arguments slice
		arguments = arguments[1:]

		var commandInfo CommandInfo
		commandInfo.Session = session
		commandInfo.Message = message
		commandInfo.Args = arguments

		cmdFunction, found := ch.Get(strings.TrimPrefix(cmdName, ch.config.Prefix))
		if !found {
			return
		}

		c := *cmdFunction
		c(commandInfo)
	}
}
