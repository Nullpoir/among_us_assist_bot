import redis
from settings.settings import *

def update_state(category_id):
  rc = redis.Redis(host=REDIS_HOST, port=REDIS_PORT)
  key = key = f'mHiyori:gaming state { category_id }'

  if rc.exists(key):
    state = rc.get(key)
    if state == b'1':
      rc.set(key, b'0')
      return WILL_MUTE
    else:
      rc.set(key, b'1')
      return WILL_DISCUSS
  else:
    rc.set(key, b'0')
    state = b'1'
    return WILL_MUTE
