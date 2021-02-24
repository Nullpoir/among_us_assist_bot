import redis

def update_state(category_id):
  rc = redis.Redis(host='localhost', port=6379)
  key = key = f'mHiyori:gaming state { category_id }'

  if rc.exists(key):
    state = rc.get(key)
    if state == b'1':
      rc.set(key, b'0')
    else:
      rc.set(key, b'1')
  else:
    rc.set(key, b'0')
    state = b'1'
  return state
