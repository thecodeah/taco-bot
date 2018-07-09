package commands

import "github.com/bwmarrin/discordgo"

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
	Commands CommandMap
}

// New creates a new command handler.
func New() (ch *CommandHandler) {
	ch = &CommandHandler{
		Commands: make(CommandMap),
	}
	return
}

// Register registers a command to be handled by the command handler.
func (ch CommandHandler) Register(name string, command Command) {
	ch.Commands[name] = command
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (ch CommandHandler) Get(name string) (*Command, bool) {
	command, found := ch.Commands[name]
	return &command, found
}
