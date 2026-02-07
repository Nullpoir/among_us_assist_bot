package utils

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ChannelNameFromID(s *discordgo.Session, channelID string) (string, error) {
	ch, err := s.State.Channel(channelID)
	if err != nil {
		ch, err := s.Channel(channelID)
		if err != nil {
			return "", err
		}
		return ch.Name, nil
	}
	return ch.Name, nil
}

// ChannelIDFromName はチャンネル名からIDを解決する
func ChannelIDFromName(s *discordgo.Session, guildID string, channelName string) (string, error) {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return "", fmt.Errorf("guild取得失敗: %w", err)
	}

	for _, ch := range guild.Channels {
		if ch.Name == channelName {
			return ch.ID, nil
		}
	}

	return "", fmt.Errorf("チャンネル '%s' が見つかりません", channelName)
}

// RoleIDFromName はロール名からIDを解決する
func RoleIDFromName(s *discordgo.Session, guildID string, roleName string) (string, error) {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return "", fmt.Errorf("guild取得失敗: %w", err)
	}

	for _, role := range guild.Roles {
		if role.Name == roleName {
			return role.ID, nil
		}
	}

	return "", fmt.Errorf("ロール '%s' が見つかりません", roleName)
}

// SetChannelSpeakPermission は指定ロールのSPEAK権限をトグルする
func SetChannelSpeakPermission(s *discordgo.Session, channelID string, roleID string, mute bool) error {
	ch, err := s.Channel(channelID)
	if err != nil {
		return fmt.Errorf("チャンネル取得失敗: %w", err)
	}

	// 既存のpermission overwriteを探す
	var existingAllow, existingDeny int64
	found := false
	for _, ow := range ch.PermissionOverwrites {
		if ow.ID == roleID && ow.Type == discordgo.PermissionOverwriteTypeRole {
			existingAllow = ow.Allow
			existingDeny = ow.Deny
			found = true
			break
		}
	}

	if mute {
		existingDeny |= discordgo.PermissionVoiceSpeak
		existingAllow &^= discordgo.PermissionVoiceSpeak
	} else {
		existingDeny &^= discordgo.PermissionVoiceSpeak
		// allowには明示的に追加しない（サーバーロール設定に委ねる）
	}

	if !found && !mute {
		// overwriteが存在せず、ミュート解除の場合は何もしない
		log.Println("permission overwriteが存在しないため、ミュート解除不要")
		return nil
	}

	return s.ChannelPermissionSet(channelID, roleID, discordgo.PermissionOverwriteTypeRole, existingAllow, existingDeny)
}

// IsUserAdmin はユーザーがADMINISTRATOR権限を持つかを判定する
func IsUserAdmin(s *discordgo.Session, guildID string, userID string) bool {
	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		return false
	}

	guild, err := s.State.Guild(guildID)
	if err != nil {
		return false
	}

	for _, role := range guild.Roles {
		for _, memberRoleID := range member.Roles {
			if role.ID == memberRoleID && (role.Permissions&discordgo.PermissionAdministrator) != 0 {
				return true
			}
		}
	}

	// サーバーオーナーも管理者扱い
	return guild.OwnerID == userID
}

// GetAdminUserIDsInVoiceChannel は指定VCにいる管理者ユーザーのIDを返す
func GetAdminUserIDsInVoiceChannel(s *discordgo.Session, guildID string, voiceChannelName string) []string {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return nil
	}

	var adminIDs []string
	for _, vs := range guild.VoiceStates {
		chName, _ := ChannelNameFromID(s, vs.ChannelID)
		if chName == voiceChannelName && IsUserAdmin(s, guildID, vs.UserID) {
			adminIDs = append(adminIDs, vs.UserID)
		}
	}

	return adminIDs
}

// IsChannelSpeakDenied は現在のミュート状態を確認する
func IsChannelSpeakDenied(s *discordgo.Session, channelID string, roleID string) (bool, error) {
	ch, err := s.Channel(channelID)
	if err != nil {
		return false, fmt.Errorf("チャンネル取得失敗: %w", err)
	}

	for _, ow := range ch.PermissionOverwrites {
		if ow.ID == roleID && ow.Type == discordgo.PermissionOverwriteTypeRole {
			return (ow.Deny & discordgo.PermissionVoiceSpeak) != 0, nil
		}
	}

	return false, nil
}
