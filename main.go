package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"lazlink/command"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("failed to load .env", slog.Any("err", err))
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		slog.Error("DISCORD_TOKEN is not set")
		return
	}

	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
			),
		),
		bot.WithEventListenerFunc(onApplicationCommand),
	)
	if err != nil {
		slog.Error("failed to create DisGo client", slog.Any("err", err))
		return
	}
	defer client.Close(context.Background())

	// SetGlobalCommands replaces the global application-command set, so commands
	// removed from command.Commands are removed automatically as well.
	if _, err = client.Rest.SetGlobalCommands(client.ApplicationID, command.Commands); err != nil {
		slog.Error("failed to sync application commands", slog.Any("err", err))
		return
	}

	if err = client.OpenGateway(context.Background()); err != nil {
		slog.Error("failed to open gateway", slog.Any("err", err))
		return
	}

	slog.Info("bot is now running", slog.String("disgo_version", disgo.Version))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func onApplicationCommand(event *events.ApplicationCommandInteractionCreate) {
	if handler, ok := command.Handlers[event.SlashCommandInteractionData().CommandName()]; ok {
		handler(event)
	}
}
