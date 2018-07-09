package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/thecodeah/Gopher/src/commands"
)

// Configuration contains settings loaded in from environment
// variable (.env) files.
type Configuration struct {
	Token  string `required:"true"`
	Prefix string `default:"!"`
}

// Bot contains information that's necessary for the bot.
type Bot struct {
	config         Configuration
	session        *discordgo.Session
	commandHandler *commands.CommandHandler
}

// New initializes the bot as well as all commands.
func New(config Configuration) (bot *Bot, err error) {
	bot = &Bot{
		config: config,
	}

	bot.session, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		return
	}

	err = bot.session.Open()
	if err != nil {
		return
	}

	bot.commandHandler = commands.New(commands.Config{
		Prefix: config.Prefix,
	})
	bot.registerCommands()

	bot.session.AddHandler(bot.messageCreate)

	return
}

// Close closes the bot
func (bot Bot) Close() {
	bot.session.Close()
}

func (bot Bot) messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	bot.commandHandler.Process(session, message)
}

func (bot Bot) registerCommands() {
	bot.commandHandler.Register("ping", commands.PingCommand)
}
