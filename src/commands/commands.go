package commands

import "github.com/bwmarrin/discordgo"

// CommandInfo stores all information about the message sent by the player.
type CommandInfo struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Args    []string
}

// Command is a function that requires CommandInfo as its argument.
type Command func(CommandInfo)

// CommandMap stores all commands, accessable by the command name.
var CommandMap map[string]Command

func init() {
	CommandMap = make(map[string]Command)
}

// Register registers a command to be handled by the command handler.
func Register(name string, command Command) {
	CommandMap[name] = command
}

// Get retrieves the Command (Data type) from the CommandMap map.
func Get(name string) (*Command, bool) {
	command, found := CommandMap[name]
	return &command, found
}
