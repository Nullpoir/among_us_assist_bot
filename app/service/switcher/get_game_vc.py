import discord

def get_game_vc(channels):
  voice_channels = [channel for channel in channels if type(channel) is discord.channel.VoiceChannel]
  game = None
  mute = None
  for voice_channel in voice_channels:
    if voice_channel.name == 'game':
      game = voice_channel
      break
  return game
