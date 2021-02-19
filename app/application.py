import discord
from settings.settings import *

# 接続に必要なオブジェクトを生成
client = discord.Client()

# 起動時に動作する処理
@client.event
async def on_ready():
    # 起動したらターミナルにログイン通知が表示される
    print('ログインしました')

# メッセージ受信時に動作する処理
@client.event
async def on_message(message):
    if message.channel.name == 'bot操作':
      if message.content == '/testtest':
        await message.channel.send('にゃーん')
    

# Botの起動とDiscordサーバーへの接続
client.run(ACCESS_TOKEN)
