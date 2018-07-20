package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// LordCommand replies with the username of the player with the most
// tacos in the guild.
func LordCommand(commandMessage CommandMessage) {
	database := commandMessage.CommandHandler.Database

	user, err := database.GetTopUser(commandMessage.Guild.ID)
	if err != nil {
		return
	}

	member, err := commandMessage.Session.GuildMember(commandMessage.Guild.ID, user.UserID)

	displayName := member.Nick
	if displayName == "" {
		displayName = member.User.Username
	}

	commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
		fmt.Sprintf("%s The lord of the tacos is **%s**! (%s tacos)",
			commandMessage.Message.Author.Mention(),
			displayName,
			humanize.Comma(int64(user.Balance)),
		),
	)
}
