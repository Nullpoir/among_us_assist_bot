package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "log"
  . "among_us_assist_bot/config"
  "github.com/bwmarrin/discordgo"
)

func main() {
  discord, err := discordgo.New()
  discord.Token = ACCESS_TOKEN
  err = discord.Open()

  // ログイン失敗時のエラーハンドリング
  if err != nil {
    fmt.Println("ログインに失敗しました")
    fmt.Println(err)
  }
  //イベントハンドラを追加
  discord.AddHandler(onMessageCreate)

  err = discord.Open()
  if err != nil {
    fmt.Println(err)
  }
  // 直近の関数（main）の最後に実行される
  defer discord.Close()

  fmt.Println("Listening...")
  stopBot := make(chan os.Signal, 1)
  signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-stopBot
  return
}
func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  Channel, err := s.Channel(m.ChannelID)
  if err != nil {
    log.Println("Error sending message: ", err)
  }
  var ChannelName string = Channel.Name

  if ChannelName == "チャット" || ChannelName == "bot操作" {
    if m.Content == "m" {
      ParentChannel, err := s.Channel(Channel.ParentID)
      if err != nil {
        log.Println("Error sending message: ", err)
      }
      fmt.Println(ParentChannel.Recipients)
    }
  }
}
  
func sendMessage(s *discordgo.Session, channelID string, msg string) {
  _, err := s.ChannelMessageSend(channelID, msg)

  if err != nil {
    log.Println("Error sending message: ", err)
  }
}
