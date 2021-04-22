import discord
from services import *
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
    # ミュート切替
    if message.content == 'm':
      text = await switcher(message)
      return await message.channel.send(text)
    # 状態クリア
    elif message.content == 'c':
      reset_game_state(message.channel.category_id)
      return await message.channel.send('リセットされました。')

  # エンタメ 
  response = create_response(message.content)
  if response != None:
    return await message.channel.send(response)
  
  return None

# VC移動管理
@client.event
async def on_voice_state_update(member, before, after):
  try:
    if before.channel.name == 'game' and after.channel.name == '天国':
      return await member.edit(mute=False)
    elif before.channel.name == '天国' and after.channel.name == 'game':
      return await member.edit(mute=True)
    else:
      return None
  except AttributeError:
    return None

# Botの起動とDiscordサーバーへの接続
client.run(ACCESS_TOKEN)
