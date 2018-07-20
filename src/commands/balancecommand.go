package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// BalanceCommand shows you your balance.
func BalanceCommand(commandMessage CommandMessage) {
	database := commandMessage.CommandHandler.Database

	user, err := database.GetUser(commandMessage.Message.Author.ID, commandMessage.Guild.ID)
	if err != nil {
		return
	}

	commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
		fmt.Sprintf("%s You have %s tacos :taco:", commandMessage.Message.Author.Mention(), humanize.Comma(int64(user.Balance))),
	)
}
