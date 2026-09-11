package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"

	"lazlink/command"
)

var token string

func main() {
	slog.Info("starting bot...")

	// .envを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("error loading .env file", slog.Any("err", err))
		return
	}

	token = os.Getenv("DISCORD_TOKEN")

	if token == "" {
		slog.Error("DISCORD_TOKEN is not set")
		return
	}

	// Botを作成
	client, err := disgo.New(
		token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
			),
		),
		bot.WithEventListenerFunc(commandListener),
	)

	if err != nil {
		slog.Error("error while building disgo", slog.Any("err", err))
		return
	}

	defer client.Close(context.TODO())

	// /pingなどのコマンドをDiscordに登録
	_, err = client.Rest.SetGlobalCommands(
		client.ApplicationID,
		command.Commands,
	)

	if err != nil {
		slog.Error(
			"error while registering commands",
			slog.Any("err", err),
		)
		return
	}

	slog.Info("commands registered")

	// Discord Gatewayに接続
	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error(
			"errors while connecting to gateway",
			slog.Any("err", err),
		)
		return
	}

	slog.Info("bot is now running. Press CTRL-C to exit.")

	// Ctrl+Cで終了
	s := make(chan os.Signal, 1)
	signal.Notify(
		s,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	<-s
}

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	switch data.CommandName() {
	case "ping":
		command.Ping(event)
	}
}