package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*
====================
 設定（環境変数）
====================
*/
var (
	BotToken            = os.Getenv("ACCESS_TOKEN")
	ControlTextChannel  = os.Getenv("CONTROL_TEXT_CH") // m を受け付ける text ch
	LobbyVC             = os.Getenv("LOBBY_VC")         // 移動元 VC
	MeetingVC           = os.Getenv("MEETING_VC")       // 移動先 VC
	MaxConcurrency      = 5                              // 並列数（10人なら5で十分）
	OperationTimeout    = 2 * time.Second               // 全体制限
)

/*
====================
 main
====================
*/

func main() {
	if BotToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN is required")
	}

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildVoiceStates |
			discordgo.IntentsMessageContent

	dg.AddHandler(onMessage)
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
 Message: m コマンド
====================
*/

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	chName, err := channelNameFromID(s, m.ChannelID)
	if err != nil {
		log.Println("failed to get channel name")
		return
	}

	if chName != ControlTextChannel {
		return
	}

	if strings.TrimSpace(m.Content) != "m" {
		return
	}

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		log.Println("failed to get guild")
		return
	}

	userIDs, _ := GetUserIDsInVoiceChannel(s, guild, MeetingVC)

	log.Println(userIDs)

	if len(userIDs) == 0 {
		return
	}

	// mute 状態は最初の1人を見て反転
	member, err := s.GuildMember(m.GuildID, userIDs[0])
	if err != nil {
		return
	}
	newMute := !member.Mute

	execMuteParallel(s, m.GuildID, userIDs, newMute)
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

	from := v.BeforeUpdate.ChannelID
	to := v.ChannelID

	if from == to {
		return
	}

	switch {
	case from == LobbyVC && to == MeetingVC:
		_ = s.GuildMemberMute(v.GuildID, v.UserID, true)

	case from == MeetingVC && to == LobbyVC:
		_ = s.GuildMemberMute(v.GuildID, v.UserID, false)
	}
}

/*
====================
 並行 mute 実行
====================
*/

func execMuteParallel(
	s *discordgo.Session,
	guildID string,
	userIDs []string,
	mute bool,
) {
	ctx, cancel := context.WithTimeout(context.Background(), OperationTimeout)
	defer cancel()

	sem := make(chan struct{}, MaxConcurrency)
	var wg sync.WaitGroup

	for _, uid := range userIDs {
		wg.Add(1)

		go func(userID string) {
			defer wg.Done()

			select {
			case sem <- struct{}{}:
				defer func() { <-sem }()
			case <-ctx.Done():
				return
			}

			if err := s.GuildMemberMute(guildID, userID, mute); err != nil {
				log.Printf("mute failed %s: %v", userID, err)
			}
		}(uid)
	}

	wg.Wait()
}

/*
====================
 utils
====================
*/

func GetUserIDsInVoiceChannel(
	s *discordgo.Session,
	guild *discordgo.Guild,
	voiceChannelName string,
) ([]string, error) {
	userIDs := make([]string, 0)

	log.Printf("%#v", guild)


	for _, vs := range guild.VoiceStates {
		targetChannelName, _ := channelNameFromID(s, vs.ChannelID)
		log.Println(voiceChannelName)
		if targetChannelName == voiceChannelName {
			userIDs = append(userIDs, vs.UserID)
		}
	}

	return userIDs, nil
}

func channelNameFromID(s *discordgo.Session, channelID string) (string, error) {
	ch, err := s.State.Channel(channelID)
	if err !=  nil {
		ch, err := s.Channel(channelID)

		if err != nil {
			return "",  err
		}

		return ch.Name, nil
	}
	return ch.Name, nil
}
