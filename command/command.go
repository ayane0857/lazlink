package command

import "github.com/disgoorg/disgo/discord"

var Commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Pingコマンド",
	},
	discord.SlashCommandCreate{
		Name:        "help",
		Description: "ヘルプコマンド",
	},
}
