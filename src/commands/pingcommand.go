package commands

// PingCommand sends a "Pong!" message back.
func PingCommand(commandMessage CommandMessage) {
	commandMessage.Session.ChannelMessageSend(commandMessage.Message.ChannelID, "Pong!")
}
