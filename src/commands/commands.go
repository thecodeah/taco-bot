package commands

import (
	"strings"
	"time"

	"github.com/thecodeah/Gopher/src/storage"

	"github.com/bwmarrin/discordgo"
)

// CommandInfo stores information about the message sent by the player.
type CommandInfo struct {
	CommandHandler *CommandHandler
	Session        *discordgo.Session
	Message        *discordgo.MessageCreate
	Args           []string
}

// Command is a function that requires CommandInfo as its argument.
type Command func(CommandInfo)

// CommandMap stores all command functions by their name.
type CommandMap map[string]Command

// Cooldowns stores when commands have been used last.
type Cooldowns map[string]time.Time

// CommandHandler stores command information/state.
type CommandHandler struct {
	Database  *storage.Database
	commands  CommandMap
	config    Config
	cooldowns Cooldowns
}

// Config stores command handler configurations.
type Config struct {
	Prefix   string
	Cooldown time.Duration
}

// New creates a new command handler.
func New(database *storage.Database, config Config) (ch *CommandHandler) {
	ch = &CommandHandler{
		commands:  make(CommandMap),
		cooldowns: make(Cooldowns),
		config:    config,
		Database:  database,
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

// Process processes incoming messages and calls the command's function.
func (ch CommandHandler) Process(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, ch.config.Prefix) {
		arguments := strings.Fields(message.Content)
		cmdName := strings.TrimPrefix(arguments[0], ch.config.Prefix)

		// Removing the command from the arguments slice
		arguments = arguments[1:]

		commandInfo := CommandInfo{
			CommandHandler: &ch,
			Session:        session,
			Message:        message,
			Args:           arguments,
		}

		cmdFunction, found := ch.Get(cmdName)
		if !found {
			return
		}

		// Check if the command is on cooldown
		if lastCooldown, ok := ch.cooldowns[cmdName]; ok {
			if time.Since(lastCooldown) < ch.config.Cooldown {
				return
			}
		}

		if ch.config.Cooldown > 0 {
			ch.cooldowns[cmdName] = time.Now()
		}

		// Call the command's function
		c := *cmdFunction
		c(commandInfo)
	}
}
