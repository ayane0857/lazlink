package command

import (
	"log/slog"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func init() {
	Handlers["help"] = func(event *events.ApplicationCommandInteractionCreate) {
		now := time.Now()
		embed := discord.Embed{
			Title:       "Help Command",
			Description: "This is a help command response.",
			Color:       0x5865F2,
			Timestamp:   &now,
		}

		components := discord.ActionRowComponent{
			Components: []discord.InteractiveComponent{
				discord.NewLinkButton("Documentation", "https://discord.com/developers/docs/interactions/message-components#buttons").
					WithEmoji(discord.ComponentEmoji{Name: "📜"}),
				discord.NewLinkButton("Discord developers", "https://discord.gg/discord-developers").
					WithEmoji(discord.ComponentEmoji{Name: "🔧"}),
				discord.NewLinkButton("Discord Gophers", "https://discord.gg/7RuRrVHyXF").
					WithEmoji(discord.ComponentEmoji{Name: "🦫"}),
			},
		}

		if err := event.CreateMessage(
			discord.NewMessageCreate().
				WithEmbeds(embed).
				WithComponents(components).
				WithEphemeral(true),
		); err != nil {
			slog.Error("failed to send help response", slog.Any("err", err))
		}
	}
}
