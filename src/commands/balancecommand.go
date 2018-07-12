package commands

import (
	"fmt"
)

// BalanceCommand shows you your balance.
func BalanceCommand(commandInfo CommandInfo) {
	database := commandInfo.CommandHandler.Database
	channel, err := commandInfo.Session.Channel(commandInfo.Message.ChannelID)
	if err != nil {
		return
	}

	user, err := database.GetUser(commandInfo.Message.Author.ID, channel.GuildID)
	if err != nil {
		return
	}

	commandInfo.Session.ChannelMessageSend(commandInfo.Message.ChannelID,
		fmt.Sprintf("<@%s> You have %d shells :shell:", commandInfo.Message.Author.ID, user.Balance),
	)
}
