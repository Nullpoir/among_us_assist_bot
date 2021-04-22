from helpers import *

async def switcher(message):
  # ミュート切替
    game = get_game_vc(message.channel.category.channels)
    state = update_state(message.channel.category_id)

    # gameからmuteへ
    if state == WILL_MUTE:
      await mute(game.members)
      text = 'ミュートにしました。'
    # muteからgameへ
    elif state == WILL_DISCUSS:
      await unmute(game.members)
      text = '議論してください！'
    
    return text
      