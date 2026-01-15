package utils

import (
	"log"
	"sync"
	"github.com/bwmarrin/discordgo"
)

func ExecMuteParallel(
	s *discordgo.Session,
	guildID string,
	userIDs []string,
	mute bool,
) {
	var wg sync.WaitGroup

	for _, uid := range userIDs {
		wg.Add(1)

		go func(userID string) {
			defer wg.Done()

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
