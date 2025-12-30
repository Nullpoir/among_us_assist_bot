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
	TargetVoiceChannel  = os.Getenv("TARGET_VOICE_CH") // mute 対象 VC
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

	if m.ChannelID != ControlTextChannel {
		return
	}

	if strings.TrimSpace(m.Content) != "m" {
		return
	}

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		log.Println("guild state error:", err)
		return
	}

	userIDs := collectUsersInVC(guild, TargetVoiceChannel)
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

func collectUsersInVC(guild *discordgo.Guild, vcID string) []string {
	users := []string{}
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == vcID {
			users = append(users, vs.UserID)
		}
	}
	return users
}
