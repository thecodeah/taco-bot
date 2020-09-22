package commands

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

type Quote struct {
	Message string
	Author string // Discord userid
}

var quotes = [...]Quote {
	Quote{
		Message: "Grrrrr.",
		Author: "622791275560828948",
	},
	Quote{
		Message: "One does not simply eat more than one taco at a time.",
		Author: "615838975931711489",
	},
	Quote{
		Message: "<takes taco>",
		Author: "635254666027859968",
	},
	Quote{
		Message: "Watch over your tacos, as one just got swallowed!",
		Author: "451795408382066699",
	},
	Quote{
		Message: "Thank you for choosing TACO Airlines, enjoy your flight",
		Author: "665902262907961360",
	},
	Quote{
		Message: "Don't get taconstipated!",
		Author: "665902262907961360",
	},
	Quote{
		Message: "Now listen to the Taco Song!",
		Author: "665902262907961360",
	},
	Quote{
		Message: "If you eat too much, you'll become T H I C C",
		Author: "276472885063974912",
	},
	Quote{
		Message: "For Kingdom and Glory!",
		Author: "622791275560828948",
	},
	Quote{
		Message: "All hail the Taco bot!",
		Author: "622791275560828948",
	},
	Quote{
		Message: "If you don’t like tacos, I’m nacho type.",
		Author: "238335599767977985",
	},
	Quote{
		Message: "Taco big or taco home",
		Author: "238335599767977985",
	},
	Quote{
		Message: "I ate a taco and all I got was this lousy message",
		Author: "669977652064550912",
	},
	Quote{
		Message: "Taco to you later",
		Author: "557123149674708994",
	},
	Quote{
		Message: "One step farther from the taco throne.",
		Author: "615838975931711489",
	},
	Quote{
		Message: "OM NOM NOM",
		Author: "298967832299831296",
	},
	Quote{
		Message: "AAAH IT BURNS TOO SPICY!!",
		Author: "404",
	},
}

var gifs = [...]string {
	"https://media3.giphy.com/media/3oKIPay5G3yCBXo3le/giphy.gif",
	"https://media4.giphy.com/media/TJrhuGPiZlUeXhJ4jI/giphy.gif",
	"https://media1.giphy.com/media/82okobf2F12rNV3XqD/giphy.gif",
	"https://media2.giphy.com/media/26BRDpgF3iwXFkEQU/giphy.gif",
	"https://media0.giphy.com/media/3CgHKSDwAT92o/giphy.gif",
	"https://media2.giphy.com/media/13VJu6tRPDBF72/giphy.gif",
	"https://media3.giphy.com/media/fBDHF3FYQ4ygaXw1HW/giphy.gif",
}

// EatCommand replies with a GIF of a taco being eaten, and takes away one taco from the user.
func EatCommand(commandMessage CommandMessage) {
	database := commandMessage.CommandHandler.Database

	// Get user data
	sender, err := database.GetUser(commandMessage.Message.Author.ID, commandMessage.Guild.ID)
	if err != nil {
		return
	}

	// Perform checks
	if sender.Balance < 1 {
		commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID,
			fmt.Sprintf("%s You don't have any tacos!", commandMessage.Message.Author.Mention()),
		)
		return
	}

	// Change balance and update in database
	sender.Balance -= 1
	err = database.UpdateUser(sender)
	if err != nil {
		return
	}

	randomQuote := quotes[rand.Intn(len(quotes))]

	var quoteAuthorString string;
	quoteAuthorUser, err := commandMessage.Session.User(randomQuote.Author)
	if err == nil {
		quoteAuthorString = fmt.Sprintf(
			"%s#%s",
			quoteAuthorUser.Username,
			quoteAuthorUser.Discriminator,
		);
	} else {
		quoteAuthorString = "Unknown User";
	}

	commandMessage.Session.ChannelMessageSendComplex(commandMessage.Message.ChannelID,
		&discordgo.MessageSend{
			// Content: fmt.Sprintf("%s", commandMessage.Message.Author.Mention()),
			Embed: &discordgo.MessageEmbed{
				Title: fmt.Sprintf("%s just ate a delicious taco! :taco:", commandMessage.Message.Author.Username),
				Description: fmt.Sprintf("*%s*", randomQuote.Message),
				Type: "image",
				Image: &discordgo.MessageEmbedImage{
					URL: gifs[rand.Intn(len(gifs))],
				},
				Color: 0xFFAC33,
				Provider: &discordgo.MessageEmbedProvider{
					URL: "http://giphy.com",
					Name: "Giphy",
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("Image Source: Giphy.com • Message submitted by: %s", quoteAuthorString),
				},
			},
		},
	)
}
