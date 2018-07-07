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

// Register registers a command to be handled by the command handler.
func (cm CommandMap) Register(name string, command Command) {
	cm[name] = command
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (cm CommandMap) Get(name string) (*Command, bool) {
	command, found := cm[name]
	return &command, found
}
