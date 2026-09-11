package command

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Commands はDiscordに登録するスラッシュコマンド一覧
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Pingコマンド",
	},
}

// Handlers はコマンド名と処理を対応させる
var Handlers = map[string]func(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
){
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		latency := s.HeartbeatLatency()
		embed := &discordgo.MessageEmbed{
			Title:       "Ping Command",
			Description: fmt.Sprintf("This is a ping command response. Latency: %v", latency),
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
	},
}