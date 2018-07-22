package commands

import (
	"strings"
	"time"

	"github.com/thecodeah/taco-bot/src/storage"

	"github.com/bwmarrin/discordgo"
)

// CommandMessage stores information about the message sent by the player.
type CommandMessage struct {
	CommandHandler *CommandHandler
	Session        *discordgo.Session
	Guild          *discordgo.Guild
	Message        *discordgo.MessageCreate
	Args           []string
}

// CommandInfo stores information about a command.
type CommandInfo struct {
	Function    Command
	Description string
	Hidden      bool
	Cooldown    map[*discordgo.Guild]time.Time
}

// Command is a function that requires CommandMessage as its argument.
type Command func(CommandMessage)

// CommandMap stores all command functions by their name.
type CommandMap map[string]CommandInfo

// CommandHandler stores command information/state.
type CommandHandler struct {
	Database *storage.Database
	commands CommandMap
	config   Config
}

// Config stores command handler configurations.
type Config struct {
	Prefix   string
	Cooldown time.Duration
}

// New creates a new command handler.
func New(database *storage.Database, config Config) (ch *CommandHandler) {
	ch = &CommandHandler{
		commands: make(CommandMap),
		config:   config,
		Database: database,
	}
	return
}

// Register registers a command to be handled by the command handler.
func (ch CommandHandler) Register(name string, commandInfo CommandInfo) {
	commandInfo.Cooldown = make(map[*discordgo.Guild]time.Time)
	ch.commands[name] = commandInfo
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (ch CommandHandler) Get(name string) (*CommandInfo, bool) {
	commandInfo, found := ch.commands[name]
	return &commandInfo, found
}

// Process processes incoming messages and calls the command's function.
func (ch CommandHandler) Process(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, ch.config.Prefix) {
		arguments := strings.Fields(strings.TrimPrefix(message.Content, ch.config.Prefix))
		cmdName := arguments[0]

		// Removing the command from the arguments slice
		arguments = arguments[1:]

		// Get the guild in which the message was sent
		channel, err := session.Channel(message.ChannelID)
		if err != nil {
			return
		}
		guild, err := session.Guild(channel.GuildID)
		if err != nil {
			return
		}

		commandMessage := CommandMessage{
			CommandHandler: &ch,
			Session:        session,
			Guild:          guild,
			Message:        message,
			Args:           arguments,
		}

		commandInfo, found := ch.Get(cmdName)
		if !found {
			return
		}
		cmdFunction := commandInfo.Function

		// Check if the command is on cooldown
		if lastCooldown, ok := commandInfo.Cooldown[guild]; ok {
			if time.Since(lastCooldown) < ch.config.Cooldown {
				return
			}
		}

		if ch.config.Cooldown > 0 {
			commandInfo.Cooldown[guild] = time.Now()
		}

		// Call the command's function
		cmdFunction(commandMessage)
	}
}
