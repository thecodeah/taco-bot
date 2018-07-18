package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// BalanceCommand shows you your balance.
func BalanceCommand(commandInfo CommandInfo) {
	database := commandInfo.CommandHandler.Database

	user, err := database.GetUser(commandInfo.Message.Author.ID, commandInfo.Guild.ID)
	if err != nil {
		return
	}

	commandInfo.Session.ChannelMessageSend(commandInfo.Message.ChannelID,
		fmt.Sprintf("%s You have %s tacos :taco:", commandInfo.Message.Author.Mention(), humanize.Comma(int64(user.Balance))),
	)
}
