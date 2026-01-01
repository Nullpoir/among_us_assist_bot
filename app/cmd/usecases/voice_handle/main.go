package voice_handle

import (
	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/configs"
	"among_us_assist_bot/cmd/utils"
)

func VoiceHandle(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
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
