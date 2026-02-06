package message_handle

import (
	"log"
	"strings"
	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/configs"
	"among_us_assist_bot/cmd/utils"
	"time"
	"strconv"
)

// 操作chからのコマンドを受信
func MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	chName, err := utils.ChannelNameFromID(s, m.ChannelID)
	if err != nil {
		log.Println("failed to get channel name")
		s.ChannelMessageSend(m.ChannelID, "システムエラーです。")
		return
	}

	if chName != configs.ControlTextChannel {
		return
	}

	if strings.TrimSpace(m.Content) != "m" {
		return
	}

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		log.Println("failed to get guild")
		s.ChannelMessageSend(m.ChannelID, "システムエラーです。")
		return
	}

	userIDs, err := utils.GetUserIDsInVoiceChannel(s, guild, configs.MeetingVC)
	if err != nil {
		log.Println("failed to get users")
		s.ChannelMessageSend(m.ChannelID, "システムエラーです。")
		return
	}

	if len(userIDs) == 0 {
		s.ChannelMessageSend(m.ChannelID, "ミュート対象が検知できませんでした...再入室をお願いします。")
		return
	}

	// mute 状態は最初の1人を見て反転
	member, err := s.GuildMember(m.GuildID, userIDs[0])
	if err != nil {
		return
	}
	newMute := !member.Mute

	start := time.Now()
	utils.ExecMuteParallel(s, m.GuildID, userIDs, newMute)
	elapsed := time.Since(start)

	elapsed_str := strconv.FormatFloat(elapsed.Seconds(), 'f', 3, 64) + "s"

	if (newMute) {
		s.ChannelMessageSend(m.ChannelID, "ミュートしました！" + elapsed_str)
	} else {
		s.ChannelMessageSend(m.ChannelID, "議論してください！" + elapsed_str)
	}
}
