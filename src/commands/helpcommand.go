package commands

import (
	"github.com/bwmarrin/discordgo"
)

// HelpCommand displays information about all commands.
func HelpCommand(commandMessage CommandMessage) {
	embed := &discordgo.MessageEmbed{
		Color: 0xFFAC33,
	}

	var fieldCount int
	prefix := commandMessage.CommandHandler.config.Prefix
	for k, v := range commandMessage.CommandHandler.commands {
		if !v.Hidden {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:  prefix + k,
				Value: v.Description,
			})

			fieldCount++
		}
	}

	if fieldCount > 5 {
		channel, err := commandMessage.Session.UserChannelCreate(commandMessage.Message.Author.ID)
		if err != nil {
			return
		}

		commandMessage.Session.ChannelMessageSendEmbed(channel.ID, embed)
	} else {
		commandMessage.Session.ChannelMessageSendEmbed(commandMessage.Message.ChannelID, embed)
	}
}
