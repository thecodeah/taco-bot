package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// LordCommand replies with the username of the player with the most
// tacos in the guild.
func LordCommand(commandInfo CommandInfo) {
	database := commandInfo.CommandHandler.Database

	user, err := database.GetTopUser(commandInfo.Guild.ID)
	if err != nil {
		return
	}

	member, err := commandInfo.Session.GuildMember(commandInfo.Guild.ID, user.UserID)

	displayName := member.Nick
	if displayName == "" {
		displayName = member.User.Username
	}

	commandInfo.Session.ChannelMessageSend(commandInfo.Message.ChannelID,
		fmt.Sprintf("%s The lord of the tacos is **%s**! (%s tacos)",
			commandInfo.Message.Author.Mention(),
			displayName,
			humanize.Comma(int64(user.Balance)),
		),
	)
}
