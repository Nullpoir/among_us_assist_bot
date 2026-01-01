package message_handle

import (
	"log"
	"context"
	"strings"
	"sync"
	"github.com/bwmarrin/discordgo"
	"among_us_assist_bot/configs"
)

func MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	chName, err := ChannelNameFromID(s, m.ChannelID)
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

	userIDs, _ := GetUserIDsInVoiceChannel(s, guild, configs.MeetingVC)

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

func execMuteParallel(
	s *discordgo.Session,
	guildID string,
	userIDs []string,
	mute bool,
) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.OperationTimeout)
	defer cancel()

	sem := make(chan struct{}, configs.MaxConcurrency)
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

func GetUserIDsInVoiceChannel(
	s *discordgo.Session,
	guild *discordgo.Guild,
	voiceChannelName string,
) ([]string, error) {
	userIDs := make([]string, 0)

	log.Printf("%#v", guild)


	for _, vs := range guild.VoiceStates {
		targetChannelName, _ := ChannelNameFromID(s, vs.ChannelID)
		log.Println(voiceChannelName)
		if targetChannelName == voiceChannelName {
			userIDs = append(userIDs, vs.UserID)
		}
	}

	return userIDs, nil
}

func ChannelNameFromID(s *discordgo.Session, channelID string) (string, error) {
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
