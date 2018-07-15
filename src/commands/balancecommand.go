package commands

import (
	"fmt"
)

// BalanceCommand shows you your balance.
func BalanceCommand(commandInfo CommandInfo) {
	database := commandInfo.CommandHandler.Database

	user, err := database.GetUser(commandInfo.Message.Author.ID, commandInfo.Guild.ID)
	if err != nil {
		return
	}

	commandInfo.Session.ChannelMessageSend(commandInfo.Message.ChannelID,
		fmt.Sprintf("%s You have %d shells :shell:", commandInfo.Message.Author.Mention(), user.Balance),
	)
}
