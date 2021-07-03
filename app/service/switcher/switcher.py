from .get_game_vc import *
from .update_state import *
from .mute import *
import time
import math

async def switcher(message):
  game = get_game_vc(message.channel.category.channels)
  state = update_state(message.channel.category_id)
  start = time.time()
  # gameからmuteへ
  if state == True:
    await set_mute(game.members, True)
    text = 'ミュートにしました。'
  # muteからgameへ
  elif state == False:
    await set_mute(game.members, False)
    text = '議論してください！'
  elapsed_time = math.floor((time.time() - start) * 1000)
  await message.channel.send(f'{text} in {elapsed_time}[ms]')
  return 0
    
