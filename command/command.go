package command

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Pingコマンド",
	},
	{
		Name:        "help",
		Description: "ヘルプコマンド",
	},
}

var Handlers = map[string]func(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
){}