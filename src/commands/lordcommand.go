package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/dustin/go-humanize/english"
)

// LordCommand replies with the username of the player with the most
// tacos in the guild.
func LordCommand(commandMessage CommandMessage) {
	database := commandMessage.CommandHandler.Database

	user, found, err := database.GetTopUser(commandMessage.Guild.ID)
	if err != nil {
		return
	}

	if user.Balance > 0 && found {
		member, err := commandMessage.Session.GuildMember(commandMessage.Guild.ID, user.UserID)
		if err != nil {
			return
		}

		displayName := member.Nick
		if displayName == "" {
			displayName = member.User.Username
		}

		commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
			fmt.Sprintf("%s The lord of the tacos is **%s**! (%s %s)",
				commandMessage.Message.Author.Mention(),
				displayName,
				humanize.Comma(int64(user.Balance)),
				english.PluralWord(user.Balance, "taco", "tacos"),
			),
		)
	} else {
		commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
			fmt.Sprintf("%s Hmmm... It looks like nobody has any tacos yet. Keep chatting!",
				commandMessage.Message.Author.Mention(),
			),
		)
	}
}
