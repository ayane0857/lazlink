package command

import "github.com/disgoorg/disgo/discord"

var Commands = []discord.ApplicationCommandCreate{
	PingCommand(),
}