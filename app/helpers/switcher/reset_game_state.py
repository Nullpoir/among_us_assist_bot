import redis

def reset_game_state(channel_id):
  rc = redis.Redis(host='localhost', port=6379)
  key = f'mHiyori:gaming state {channel_id}'
  rc.delete(key)
