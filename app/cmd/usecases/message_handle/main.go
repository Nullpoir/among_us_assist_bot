package message_handle

import (
	"log"
	"strings"
	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/configs"
	"among_us_assist_bot/cmd/utils"
)

// 操作chからのコマンドを受信
func MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	chName, err := utils.ChannelNameFromID(s, m.ChannelID)
	if err != nil {
		log.Println("failed to get channel name")
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
		return
	}

	userIDs, err := utils.GetUserIDsInVoiceChannel(s, guild, configs.MeetingVC)
	if err != nil {
		log.Println("failed to get users")
		return
	}

	if len(userIDs) == 0 {
		return
	}

	// mute 状態は最初の1人を見て反転
	member, err := s.GuildMember(m.GuildID, userIDs[0])
	if err != nil {
		return
	}
	newMute := !member.Mute

	utils.ExecMuteParallel(s, m.GuildID, userIDs, newMute)

	s.ChannelMessageSend(m.ChannelID, "議論してください！")
}
