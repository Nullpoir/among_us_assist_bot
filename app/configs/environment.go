package configs

import (
	"os"
)

var (
	BotToken            = os.Getenv("ACCESS_TOKEN")
	DiscordToken        = "Bot " + BotToken
	ControlTextChannel  = os.Getenv("CONTROL_TEXT_CH") // m を受け付ける text ch
	LobbyVC             = os.Getenv("LOBBY_VC")         // 移動元 VC
	MeetingVC           = os.Getenv("MEETING_VC")       // 移動先 VC
	MaxConcurrency      = 5                              // 並列数（10人なら5で十分）
)
