package commands

// PingCommand sends a "Pong!" message back.
func PingCommand(commandInfo CommandInfo) {
	commandInfo.Session.ChannelMessageSend(commandInfo.Message.ChannelID, "Pong!")
}
