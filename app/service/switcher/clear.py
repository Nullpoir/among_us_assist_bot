from .reset_game_state import *

async def clear(message):
  reset_game_state(message.channel.category_id)
  await message.channel.send('リセットされました。')
  return 0
