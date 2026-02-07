package message_handle

import (
	"log"
	"strconv"
	"strings"
	"time"

	"among_us_assist_bot/cmd/utils"
	"among_us_assist_bot/configs"

	"github.com/bwmarrin/discordgo"
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

	// 会議VCのチャンネルIDを取得
	meetingChID, err := utils.ChannelIDFromName(s, m.GuildID, configs.MeetingVC)
	if err != nil {
		log.Printf("会議VCの取得失敗: %v", err)
		s.ChannelMessageSend(m.ChannelID, "会議VCが見つかりません。")
		return
	}

	// ミュート対象ロールのIDを取得
	roleID, err := utils.RoleIDFromName(s, m.GuildID, configs.MuteRole)
	if err != nil {
		log.Printf("ロールの取得失敗: %v", err)
		s.ChannelMessageSend(m.ChannelID, "ミュート対象ロールが見つかりません。")
		return
	}

	// 現在のミュート状態を確認
	isMuted, err := utils.IsChannelSpeakDenied(s, meetingChID, roleID)
	if err != nil {
		log.Printf("ミュート状態の確認失敗: %v", err)
		s.ChannelMessageSend(m.ChannelID, "システムエラーです。")
		return
	}

	newMute := !isMuted

	start := time.Now()
	err = utils.SetChannelSpeakPermission(s, meetingChID, roleID, newMute)
	if err != nil {
		log.Printf("パーミッション変更失敗: %v", err)
		s.ChannelMessageSend(m.ChannelID, "ミュート操作に失敗しました。")
		return
	}

	// 管理者はパーミッションオーバーライドを無視するため個別にミュート
	adminIDs := utils.GetAdminUserIDsInVoiceChannel(s, m.GuildID, configs.MeetingVC)
	for _, uid := range adminIDs {
		if err := s.GuildMemberMute(m.GuildID, uid, newMute); err != nil {
			log.Printf("管理者ミュート失敗 %s: %v", uid, err)
		}
	}

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 3, 64) + "s"

	if newMute {
		s.ChannelMessageSend(m.ChannelID, "ミュートしました！"+elapsedStr)
	} else {
		s.ChannelMessageSend(m.ChannelID, "議論してください！"+elapsedStr)
	}
}
