package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/cmd/usecases/message_handle"
	"among_us_assist_bot/configs"
	"among_us_assist_bot/cmd/utils"
)

/*
====================
 main
====================
*/

func main() {
	if configs.BotToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN is required")
	}

	dg, err := discordgo.New("Bot " + configs.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildVoiceStates |
			discordgo.IntentsMessageContent

	dg.AddHandler(message_handle.MessageHandle)
	dg.AddHandler(onVoiceStateUpdate)

	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	defer dg.Close()

	log.Println("Bot is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}

/*
====================
 VoiceState 移動補正
====================
*/

func onVoiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.BeforeUpdate == nil {
		return
	}

	from, err := utils.ChannelNameFromID(s, v.BeforeUpdate.ChannelID)
	if(err != nil) {
		return
	}
	to, err := utils.ChannelNameFromID(s, v.ChannelID)
	if(err != nil) {
		return
	}

	if from == to {
		return
	}

	switch {
	case from == configs.LobbyVC && to == configs.MeetingVC:
		_ = s.GuildMemberMute(v.GuildID, v.UserID, true)

	case from == configs.MeetingVC && to == configs.LobbyVC:
		_ = s.GuildMemberMute(v.GuildID, v.UserID, false)
	}
}


