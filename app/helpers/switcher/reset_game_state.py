import redis
from settings.settings import *

def reset_game_state(channel_id):
  rc = redis.Redis(host=REDIS_HOST, port=REDIS_PORT)
  key = f'mHiyori:gaming state {channel_id}'
  rc.delete(key)
