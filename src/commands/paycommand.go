package commands

import (
	"fmt"
	"strconv"
)

// PayCommand allows you to transfer funds to another user.
func PayCommand(commandMessage CommandMessage) {
	if len(commandMessage.Args) != 2 {
		commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
			fmt.Sprintf("%s Usage : %spay [mention] [amount]", commandMessage.Message.Author.Mention(), commandMessage.CommandHandler.config.Prefix),
		)
	} else {
		database := commandMessage.CommandHandler.Database
		amount, err := strconv.Atoi(commandMessage.Args[1])
		if err != nil {
			return
		}

		// Get user data
		sender, err := database.GetUser(commandMessage.Message.Author.ID, commandMessage.Guild.ID)
		if err != nil {
			return
		}
		receiver, err := database.GetUser(commandMessage.Message.Mentions[0].ID, commandMessage.Guild.ID)
		if err != nil {
			return
		}

		// Perform checks
		if amount <= 0 || sender.Balance < amount || commandMessage.Message.Mentions[0].ID == commandMessage.Message.Author.ID {
			return
		}

		// Change balance and update in database
		sender.Balance -= amount
		receiver.Balance += amount
		err = database.UpdateUser(sender)
		if err != nil {
			return
		}
		err = database.UpdateUser(receiver)
		if err != nil {
			return
		}

		commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
			fmt.Sprintf("%s Successfully transferred tacos!", commandMessage.Message.Author.Mention()),
		)
	}
}
