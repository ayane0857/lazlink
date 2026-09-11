package command

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func PingCommand() discord.ApplicationCommandCreate {
	return discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Pingを返します",
	}
}

func Ping(event *events.ApplicationCommandInteractionCreate) {
	latency := event.Client().Gateway.Latency()
	embed := discord.Embed{
		Title:       "Pong!",
		Description: fmt.Sprintf("Ping: %v", latency),
		Color:       0x1E1F22,
	}
	event.CreateMessage(
		discord.NewMessageCreate().
			WithEmbeds(embed),
	)
}