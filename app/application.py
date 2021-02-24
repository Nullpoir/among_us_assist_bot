import discord
from helpers import *
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
      game = get_game_vc(message.channel.category.channels)
      state = update_state(message.channel.category_id)
      
      # muteからgameへ
      if state == b'1':
        for member in game.members:
          await member.edit(mute=True)
      # muteからgameへ
      else:
        for member in game.members:
          await member.edit(mute=False)

      await message.channel.send('き、切り替えました')
      return 0
    # 状態クリア
    elif message.content == 'c':
      reset_game_state(message.channel.category_id)
      await message.channel.send('リセットされました。')
      return 0
    else:
      pass

  # エンタメ 
  response = create_response(message.content)
  if response == None:
    return 0
  else:
    await message.channel.send(response)

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
