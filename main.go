package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"lazlink/command"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Variables used for command line parameters
var (
	Token string
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("読み込み出来ませんでした: %v", err)
	} 
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := command.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	// In this example, we care about receiving message events and their content.
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// Set the bot's status
	err = dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "彩音がGoを討伐中",
				Type: discordgo.ActivityTypeGame,
			},
		},
		Status: "online",
	})
	if err != nil {
		log.Println("status update error:", err)
	}


	log.Println("Adding commands...")
	existingCommands, err := dg.ApplicationCommands(dg.State.User.ID, "")
	if err != nil {
		log.Panicf("Cannot get existing commands: %v", err)
	}

	// 現在のコードで使用しているコマンド名をセットにする
	currentCommands := make(map[string]bool)

	for _, cmd := range command.Commands {
		currentCommands[cmd.Name] = true
	}

	// 不要になったコマンドを削除
	for _, cmd := range existingCommands {
		if !currentCommands[cmd.Name] {
			log.Printf("Deleting old command: %s", cmd.Name)

			err := dg.ApplicationCommandDelete(
				dg.State.User.ID,
				"",
				cmd.ID,
			)
			if err != nil {
				log.Printf("Cannot delete command '%s': %v", cmd.Name, err)
			}
		}
	}
	registeredCommands := make([]*discordgo.ApplicationCommand, len(command.Commands))
	for i, v := range command.Commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}