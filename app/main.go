package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/cmd/usecases/message_handle"
	"among_us_assist_bot/cmd/usecases/voice_handle"
	"among_us_assist_bot/configs"
)

func main() {
	if configs.BotToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN is required")
	}

	dg, err := discordgo.New(configs.DiscordToken)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuilds |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildVoiceStates |
			discordgo.IntentsMessageContent

	dg.AddHandler(message_handle.MessageHandle)
	dg.AddHandler(voice_handle.VoiceHandle)

	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	defer dg.Close()

	log.Println("Bot is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}

