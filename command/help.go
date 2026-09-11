package command

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Handlers["help"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		latency := s.HeartbeatLatency()
		embed := &discordgo.MessageEmbed{
			Title:       "Help Command",
			Description: fmt.Sprintf("This is a help command response. Latency: %v", latency),
			Color:       0x5865F2,
			Timestamp:   time.Now().Format(time.RFC3339),
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})

		if err != nil {
			// 必要ならログ処理
		}
	}
}