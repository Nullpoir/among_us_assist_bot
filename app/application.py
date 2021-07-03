import discord
from service import *
from settings.settings import *

# 接続に必要なオブジェクトを生成
client = discord.Client()

# 起動時に動作する処理
@client.event
async def on_ready():
  # 起動したらターミナルにログイン通知が表示される
  print('Ready')

# メッセージ受信時に動作する処理
@client.event
async def on_message(message):
  # ゲームルームのミュート制御
  if message.channel.name == 'bot操作' or message.channel.name == 'チャット':
    if message.content == 'm':
      await switcher(message)
      return 0
    elif message.content == 'c':
      await clear(message)
      return 0
  else:
    return 0

# VC移動管理
@client.event
async def on_voice_state_update(member, before, after):
  if before.channel.name == 'game' and after.channel.name == '天国':
    await member.edit(mute=False)
  elif before.channel.name == '天国' and after.channel.name == 'game':
    await member.edit(mute=True)
  else:
    return 0

# Botの起動とDiscordサーバーへの接続
client.run(ACCESS_TOKEN)
