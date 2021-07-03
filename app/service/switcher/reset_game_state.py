from interface import rc

def reset_game_state(channel_id):
  key = f'mHiyori:gaming state {channel_id}'
  rc.delete(key)
