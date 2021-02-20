import discord
import redis
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
  if message.channel.name == 'bot操作' or message.channel.name == 'チャット':
    if message.content == 'm':
      rc = redis.Redis(host='localhost', port=6379)
      game,mute = get_game_vc(message.channel.category.channels)
      key = f'mHiyori:gaming state {message.channel.category_id}'

      if rc.exists(key):
        state = rc.get(key)
        if state == b'1':
          rc.set(key, b'0')
        else:
          rc.set(key, b'1')
      else:
        rc.set(key, b'1')
        state = b'1'
      
      if state == b'1':
        for member in mute.members:
          if member.guild_permissions.value == 2147483647:
            await member.edit(mute=False)
          await member.move_to(game)
      else:
        for member in game.members:
          if member.guild_permissions.value == 2147483647:
            await member.edit(mute=True)
          await member.move_to(mute)

      await message.channel.send('き、切り替えました')
    if message.content == 'c':
      rc = redis.Redis(host='localhost', port=6379)
      key = f'mHiyori:gaming state {message.channel.category_id}'
      rc.delete(key)
      await message.channel.send('リセットされました。')
  if message.channel.name == '/hiyochi':
    await message.channel.send('に、にゃーん・・・///')

# Botの起動とDiscordサーバーへの接続
client.run(ACCESS_TOKEN)
