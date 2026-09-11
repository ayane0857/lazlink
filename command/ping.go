package command

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func init() {
	Handlers["ping"] = func(event *events.ApplicationCommandInteractionCreate) {
		latency := event.Client().Gateway.Latency()

		now := time.Now()
		embed := discord.Embed{
			Title:       "Ping Command",
			Description: fmt.Sprintf("This is a ping command response. Latency: %v", latency),
			Color:       0x5865F2,
			Timestamp:   &now,
		}

		if err := event.CreateMessage(
			discord.NewMessageCreate().
				WithEmbeds(embed).
				WithEphemeral(true),
		); err != nil {
			slog.Error("failed to send ping response", slog.Any("err", err))
		}
	}
}
