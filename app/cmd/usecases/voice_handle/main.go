package voice_handle

import (
	"log"

	"among_us_assist_bot/cmd/utils"
	"among_us_assist_bot/configs"

	"github.com/bwmarrin/discordgo"
)

// VoiceHandle はVC移動イベントを処理する
// 通常ユーザーはチャンネルパーミッションで自動適用されるが、
// 管理者はパーミッションオーバーライドを無視するため個別にGuildMemberMuteする
func VoiceHandle(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.BeforeUpdate == nil {
		return
	}

	if !utils.IsUserAdmin(s, v.GuildID, v.UserID) {
		return
	}

	from, err := utils.ChannelNameFromID(s, v.BeforeUpdate.ChannelID)
	if err != nil {
		return
	}
	to, err := utils.ChannelNameFromID(s, v.ChannelID)
	if err != nil {
		return
	}

	if from == to {
		return
	}

	// Lobby→Meeting: ミュート状態が有効なら管理者を個別ミュート
	if from == configs.LobbyVC && to == configs.MeetingVC {
		meetingChID, err := utils.ChannelIDFromName(s, v.GuildID, configs.MeetingVC)
		if err != nil {
			return
		}
		roleID, err := utils.RoleIDFromName(s, v.GuildID, configs.MuteRole)
		if err != nil {
			return
		}
		isMuted, err := utils.IsChannelSpeakDenied(s, meetingChID, roleID)
		if err != nil {
			return
		}
		if isMuted {
			if err := s.GuildMemberMute(v.GuildID, v.UserID, true); err != nil {
				log.Printf("管理者ミュート失敗 %s: %v", v.UserID, err)
			}
		}
	}

	// Meeting→Lobby: 管理者のサーバーミュートを解除
	if from == configs.MeetingVC && to == configs.LobbyVC {
		if err := s.GuildMemberMute(v.GuildID, v.UserID, false); err != nil {
			log.Printf("管理者ミュート解除失敗 %s: %v", v.UserID, err)
		}
	}
}
